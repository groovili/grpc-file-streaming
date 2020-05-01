package main

import (
	"fmt"
	"grpc-file-streaming/proto"
	"io"
	"io/ioutil"
	"net"
	"os"
	"os/signal"

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
		return err
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
		return err
	}

	size += int64(n)

	// receive stream parts until EOF, this means client sent all parts
	for {
		// check deadline
		if done := stream.Context().Err(); done != nil {
			// TODO: remove file if deadline exceeded and sent gRPC status
			return err
		}

		st, err = stream.Recv()
		if err != nil {
			// all chunks are received
			if err == io.EOF {
				break
			}

			log.Errorf("Error while receiving stream: %v", err)
			return err
		}

		n, err := f.Write(st.Data.Content)
		if err != nil {
			log.Errorf("Error while writing file: %v", err)
			return err
		}

		size += int64(n)
	}

	resp := &proto.FileResponse{
		Success: false,
		Size:    size,
	}

	if size == expectedSize {
		resp.Success = true
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

	go func() {
		log.Info("Starting gRPC server..")

		if err := s.Serve(l); err != nil {
			log.Error(err)
			return
		}
	}()

	<-stop
	log.Info("Stopping gRPC server...")

	// close server for new cons, finish ongoing
	s.GracefulStop()

	log.Info("Server stopped")
}
