package main

import (
	"fmt"
	"grpc-file-streaming/proto"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"path"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

type server struct {
	tmpDir string
}

func newServer() (*server, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	tmpDir, err := ioutil.TempDir(fmt.Sprintf("%s%s", dir, string(os.PathSeparator)), "server_tmp_*")
	if err != nil {
		return nil, err
	}

	return &server{tmpDir: tmpDir}, nil
}

func (s *server) shutdown() error {
	return os.RemoveAll(s.tmpDir)
}

func (s *server) SendStream(stream proto.FileStreaming_SendStreamServer) error {
	// receive first part of stream with required metadata
	st, err := stream.Recv()
	if err != nil {
		return err
	}

	// create file in temporary dir
	f, err := os.Create(fmt.Sprintf("%s%s%s", s.tmpDir, string(os.PathSeparator), st.GetFilename()))
	if err != nil {
		log.Errorf("Error while creating file: %v", err)

		return status.Error(codes.Internal, "failed to create file")
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Error(err)
		}
	}()

	expectedSize := st.GetSize()
	var size int64

	// write first chunk of data
	n, err := f.Write(st.Data.Content)
	if err != nil {
		log.Errorf("Error while writing file: %v", err)

		return status.Error(codes.Internal, "failed to write file")
	}

	size += int64(n)
	var withErrors bool
	serverDeadline := time.After(time.Second * 30)

	// receive stream parts until EOF, this means client sent all parts
	for {

		// check deadlines
		if done := stream.Context().Err(); done != nil {
			log.Info("Client deadline exceeded")

			return status.Error(codes.DeadlineExceeded, "client deadline for stream exceeded")
		}

		select {
		case <-serverDeadline:
			log.Info("Server deadline exceeded")
			return status.Error(codes.DeadlineExceeded, "server deadline for stream exceeded")
		default:
		}

		st, err = stream.Recv()
		if err != nil {
			if err != io.EOF {
				withErrors = true

				log.Errorf("Error while receiving stream: %v", err)
			}

			// file successfully received
			break
		}

		// write stream content to file
		n, err := f.Write(st.Data.Content)
		if err != nil {
			log.Errorf("Error while writing file: %v", err)

			return status.Error(codes.Internal, "failed to write file")
		}

		size += int64(n)
	}

	resp := &proto.FileResponse{
		Success: false,
		Size:    size,
	}

	log.Infof("Received file %s with size %d kB", path.Base(f.Name()), size>>10)

	if size == expectedSize && !withErrors {
		resp.Success = true
	} else {
		go func() {
			log.Infof("File %s wasn't fully received. Removing..", path.Base(f.Name()))
			_ = os.Remove(f.Name())
		}()
	}

	return stream.SendAndClose(resp)
}

func main() {
	host, port := "0.0.0.0", "50051"

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	l, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("Failed to listen on %s:%s : %v", host, port, err)
	}
	defer func() {
		log.Info("Closing tcp connection..")

		if err := l.Close(); err != nil {
			log.Errorf("Failed to close tcp connection: %v", err)
		}
	}()

	service, err := newServer()
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	defer func() {
		log.Info("Shutting down service..")

		if err := service.shutdown(); err != nil {
			log.Error(err)
		}
	}()

	s := grpc.NewServer([]grpc.ServerOption{}...)

	proto.RegisterFileStreamingServer(s, service)

	serverErr := make(chan error, 1)
	go func() {
		log.Info("Starting gRPC server..")

		serverErr <- s.Serve(l)
	}()

	select {
	case err := <-serverErr:
		log.Errorf("Server error: %v", err)
		break
	case <-stop:
		log.Info("Stopped via signal")
	}

	log.Info("Stopping gRPC server...")

	// close server for new cons, finish ongoing
	s.GracefulStop()

	log.Info("Server stopped")
}
