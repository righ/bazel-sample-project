package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pb "github.com/righ/go-sample-bazel-project/protobuf"

	"google.golang.org/grpc"
)

func main() {
	echoHost := os.Getenv("ECHO_HOST")
	if echoHost == "" {
		echoHost = "localhost"
	}
	address := echoHost + ":8001"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	cli := pb.NewEchoClient(conn)
	serve := func(w http.ResponseWriter, req *http.Request) {
		msg, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Fatalf("could not read: %v", err)
			return
		}
		res, err := cli.Echo(context.Background(), &pb.Message{Message: string(msg)})
		if err != nil {
			log.Fatalf("could not receive: %v", err)
			return
		}
		_, err = w.Write([]byte(res.GetMessage()))
		if err != nil {
			log.Fatalf("could not write: %v", err)
		}
	}
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello world"))
	})
	http.HandleFunc("/echo", serve)
	http.ListenAndServe(":8000", nil)
}
