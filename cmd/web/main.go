package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zcong1993/istio-grpc-demo-go/pb"
	"github.com/zcong1993/istio-helpers/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
)

func main() {
	var (
		UPSTREAM = os.Getenv("UPSTREAM")
		APP_NEME = os.Getenv("APP_NAME")
	)

	c, err := grpc.Dial(UPSTREAM, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUuidServiceClient(c)

	r := gin.Default()
	r.GET("/*all", func(c *gin.Context) {
		fmt.Printf("incoming tracing: %+v\n", c.Request.Header)
		ctx := tracing.Http2grpc(context.Background(), tracing.DefaultTracingKeysWeb, c.Request.Header)
		omd, ok := metadata.FromOutgoingContext(ctx)
		if ok {
			fmt.Printf("outgoing tracing: %+v\n", omd)
		}
		resp, err := client.Uuid(ctx, &pb.Request{})
		if err != nil {
			fmt.Printf("err: %+v\n", err)
			c.AbortWithError(500, err)
			return
		}
		resp.Tracing = fmt.Sprintf("%s -> %s", resp.Tracing, APP_NEME)
		c.JSON(200, resp)
	})

	r.Run(":8080")
}
