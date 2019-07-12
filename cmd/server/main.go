package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"github.com/zcong1993/istio-grpc-demo-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type UuidResp struct {
	Uuid string `json:"uuid"`
}

var client = &http.Client{Timeout: time.Second * 5}

func envOrDefault(key, defaultVal string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val
}

// HelloService is our grpc service
type UuidService struct {
	client  pb.UuidServiceClient
	appName string
	upstream string
}

// Echo impl Echo method
func (service *UuidService) Uuid(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	httpResp, err := client.Get(service.upstream)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()

	var uuidResp UuidResp
	err = json.NewDecoder(httpResp.Body).Decode(&uuidResp)
	if err != nil {
		return nil, err
	}

	fmt.Printf("recieve message: %+v\n", in)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Printf("incoming tracing: %+v\n", md)
	}

	resp := &pb.Response{
		Uuid: uuidResp.Uuid,
		Tracing: service.appName,
	}
	return resp, nil
}

func runRpcServer(port, upstream, appName string) {
	c, err := grpc.Dial(upstream, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUuidServiceClient(c)

	ss := grpc.NewServer()
	pb.RegisterUuidServiceServer(ss, &UuidService{client: client, appName: appName, upstream: upstream})

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
		UPSTREAM = envOrDefault("UPSTREAM", "https://httpbin.org/uuid")
		APP_NAME = envOrDefault("APP_NAME", "middleware-go")
	)
	runRpcServer(PORT, UPSTREAM, APP_NAME)
}
