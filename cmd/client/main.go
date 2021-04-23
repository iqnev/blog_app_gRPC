package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/iqnev/blog_app_gRPC/internal/client_blog"
)

func main() {

	log := hclog.Default()

	app := &client_blog.App{}
	app.Initialize(log)

	app.Run()

}
