package rest

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/iqnev/blog_app_gRPC/internal/client_blog/utiles"
	protos "github.com/iqnev/blog_app_gRPC/proto/blog"
)

func GetAllBlogs(w http.ResponseWriter, r *http.Request, blog protos.BlogServiceClient) {

	req := &protos.ListBlogsReq{}
	stream, err := blog.ListBlogs(context.Background(), req)

	if err != nil {
		//return err
		fmt.Println("Here1")
	}

	resp := make([]*protos.Blog, 0)

	for {
		res, err := stream.Recv()
		// If end of stream, break the loop
		if err == io.EOF {
			break
		}
		if err != nil {
			//return err
			fmt.Println("Here2")
		}
		resp = append(resp, res.GetBlog())
		//fmt.Println(res.GetBlog())
	}

	utiles.RespondJSON(w, http.StatusOK, resp)
	//w.Write([]byte("Ivelin"))
}
