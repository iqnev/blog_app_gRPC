package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/iqnev/blog_app_gRPC/internal/client_blog"
	"google.golang.org/grpc"

	protos "github.com/iqnev/blog_app_gRPC/proto/blog"
)

func main() {

	log := hclog.Default()

	conn, err := grpc.Dial("localhost:8989", grpc.WithInsecure())
	if err != nil {
		log.Error("Did not connect: %v", err)
		panic(err)
	}

	defer conn.Close()

	blogService := protos.NewBlogServiceClient(conn)

	app := &client_blog.App{}
	app.Initialize(log, blogService)

	app.Run()

}
