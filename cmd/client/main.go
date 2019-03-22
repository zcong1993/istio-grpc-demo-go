package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zcong1993/istio-grpc-demo-go/pb"
	"google.golang.org/grpc"
)

func main() {
	c, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	client := pb.NewUuidServiceClient(c)

	resp, err := client.Uuid(context.Background(), &pb.Request{Type: "hello grpc"})

	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}

	fmt.Println(resp.Uuid)
}
