package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"dineflow-order-service/pb"

	"google.golang.org/grpc"
)

type CreatePostClient struct {
	service pb.PostServiceClient
}

func NewCreatePostClient(conn *grpc.ClientConn) *CreatePostClient {
	service := pb.NewPostServiceClient(conn)

	return &CreatePostClient{service}
}

func (createPostClient *CreatePostClient) CreatePost(args *pb.CreatePostRequest) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*1000))
	defer cancel()

	res, err := createPostClient.service.CreatePost(ctx, args)

	if err != nil {
		log.Fatalf("CreatePost: %v", err)
	}

	fmt.Println(res)
}
