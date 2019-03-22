package main

import (
	"context"
	"fmt"
	"github.com/zcong1993/istio-grpc-demo-go/pb"
	"github.com/zcong1993/istio-helpers/tracing"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func envOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// HelloService is our grpc service
type UuidService struct {
	client pb.UuidServiceClient
}

// Echo impl Echo method
func (service *UuidService) Uuid(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	fmt.Printf("recieve message: %+v\n", in)
	return service.client.Uuid(tracing.Grpc2Grpc(ctx), in)
}

func runRpcServer(port, upstream string) {
	c, err := grpc.Dial(upstream, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUuidServiceClient(c)

	ss := grpc.NewServer()
	pb.RegisterUuidServiceServer(ss, &UuidService{client: client})

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	if err = ss.Serve(listener); err != nil {
		log.Fatal("ListenTCP error:", err)
	}
}

func main() {
	var (
		PORT     = envOrDefault("PORT", ":1234")
		UPSTREAM = envOrDefault("UPSTREAM", "")
	)
	runRpcServer(PORT, UPSTREAM)
}
