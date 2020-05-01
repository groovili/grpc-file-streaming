package main

import (
	"context"
	"fmt"
	"grpc-file-streaming/proto"
	"io"
	"os"
	"path"
	"time"

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

	host, port := "0.0.0.0", "50051"

	con, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Failed to dial to %s:%s :", host, port, err)
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
				resp, err = stream.CloseAndRecv()

				break
			}

			log.Errorf("Failed to read file part: %v", err)
			_ = stream.CloseSend()
			break
		}

		req.Data.Content = buf

		if err := stream.Send(req); err != nil {
			if err != io.EOF {
				log.Errorf("Failed to send stream: %v", err)
			}

			_ = stream.CloseSend()

			break
		}

		offset += int64(n)
	}

	if resp.Success {
		log.Infof("Successfully send file with size %d mB", resp.Size>>10)
	}

	log.Infof("Task took %d seconds", int(time.Now().Sub(startTime).Seconds()))
}
