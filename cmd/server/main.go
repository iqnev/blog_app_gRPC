package main

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/iqnev/blog_app_gRPC/pkg/models"
	protos "github.com/iqnev/blog_app_gRPC/proto/blog"
	"google.golang.org/grpc"
)

type BlogServiceServer struct {
	log hclog.Logger
}

var BlogList = []*models.Blog{
	&models.Blog{
		ID:       "123",
		AuthorID: "Dimitar",
		Title:    "Zaglavie 1",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
	},
	&models.Blog{
		ID:       "456",
		AuthorID: "Aishe",
		Title:    "Zaglavie 2",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
	},
	&models.Blog{
		ID:       "789",
		AuthorID: "Niiii",
		Title:    "Zaglavie 3",
		Content:  "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s",
	},
}

func (b *BlogServiceServer) ListBlogs(req *protos.ListBlogsReq, stream protos.BlogService_ListBlogsServer) error {
	b.log.Info("Calling the listBlogs method")

	for _, v := range BlogList {
		stream.Send(&protos.ListBlogsRes{
			Blog: &protos.Blog{
				Id:       v.ID,
				AuthorId: v.AuthorID,
				Title:    v.Title,
				Content:  v.Content,
			},
		})
	}

	return nil
}

func (b *BlogServiceServer) CreateBlog(ctx context.Context, req *protos.CreateBlogReq) (*protos.CreateBlogRes, error) {

	return nil, nil
}

func (b *BlogServiceServer) ReadBlog(ctx context.Context, rB *protos.ReadBlogReq) (*protos.ReadBlogRes, error) {
	return nil, nil
}
func (b *BlogServiceServer) UpdateBlog(ctx context.Context, upB *protos.UpdateBlogReq) (*protos.UpdateBlogRes, error) {
	return nil, nil
}

func (b *BlogServiceServer) DeleteBlog(ctx context.Context, dB *protos.DeleteBlogReq) (*protos.DeleteBlogRes, error) {
	return nil, nil
}

func main() {

	log := hclog.Default()

	gc := grpc.NewServer()
	srv := &BlogServiceServer{log: log}

	protos.RegisterBlogServiceServer(gc, srv)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 8989))
	if err != nil {
		log.Error("Unable to create listener", "error", err)
		os.Exit(1)
	}

	gc.Serve(l)

}
