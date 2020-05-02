package main

import (
	"context"
	"fmt"
	"grpc-file-streaming/proto"
	"io"
	"os"
	"path"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"

	log "github.com/sirupsen/logrus"
	grpc "google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide filename as second arg")
	}

	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Errorf("Failed to open file: %v", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Errorf("Failed to close file: %v", err)
		}
	}()

	host, port := "localhost", "50051"

	cred, err := credentials.NewClientTLSFromFile("./ssl/ca.crt", "")
	if err != nil {
		log.Errorf("Error while creating tls client: %v", err)
		return
	}
	opts := grpc.WithTransportCredentials(cred)

	con, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), opts)
	if err != nil {
		log.Errorf("Failed to dial to %s:%s :%v", host, port, err)
		return
	}
	defer func() {
		log.Info("Closing gRPC connection..")
		if err := con.Close(); err != nil {
			log.Error(err)
		}
	}()

	client := proto.NewFileStreamingClient(con)

	_, name := path.Split(f.Name())
	stat, _ := f.Stat()

	req := &proto.FileRequest{
		Filename: name,
		Size:     stat.Size(),
		Data: &proto.File{
			Content: nil,
		},
	}

	var resp *proto.FileResponse

	bufSize := 1 << 5 // 0.5 MB
	buf := make([]byte, bufSize)

	ctx, canc := context.WithTimeout(context.Background(), time.Minute)
	defer canc()

	stream, err := client.SendStream(ctx)
	if err != nil {
		log.Errorf("Failed to call SendStream: %v", err)
		return
	}

	startTime := time.Now()
	var offset int64
	for {
		n, err := f.ReadAt(buf, offset)
		if err != nil {
			if err == io.EOF {
				// receive response and sending break loop
				if resp, err = stream.CloseAndRecv(); err != nil {
					log.Errorf("Failed closing stream: %v", err)
				}

				// successfully sent file
				break
			}

			log.Errorf("Failed to read file part: %v", err)

			if err := stream.CloseSend(); err != nil {
				log.Errorf("Failed closing send: %v", err)
			}

			break
		}

		req.Data.Content = buf

		if err := stream.SendMsg(req); err != nil {
			// receive status error message from server
			err = stream.RecvMsg(nil)

			sErr, ok := status.FromError(err)
			if ok {
				switch sErr.Code() {
				case codes.DeadlineExceeded:
					log.Infof("Deadline exceeded: %s", sErr.Message())
				case codes.Internal:
					log.Errorf("Server error: %s", sErr.Message())
				}
			} else {
				log.Errorf("Transport layer error: %v", err)
			}

			// in any case close send and break loop
			if err := stream.CloseSend(); err != nil {
				log.Errorf("Failed closing send: %v", err)
			}

			break
		}

		offset += int64(n)
	}

	if resp == nil {
		log.Error("Job failed")
		return
	}

	if resp.Success {
		log.Infof("Successfully send file with size %d kB", resp.Size>>10)
	}

	log.Infof("Task took %d seconds", int(time.Since(startTime).Seconds()))
}
