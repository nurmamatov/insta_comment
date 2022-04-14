package grpcclient

import (
	"fmt"
	"tasks/Instagram_clone/insta_comment/config"
	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"
	// pp "tasks/Instagram_clone/insta_comment/genproto/post_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	CommentService() pc.CommentServiceClient
}

// Client
type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (*GrpcClient, error) {

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"comment_service":    pc.NewCommentServiceClient(connComment),
		},
	}, nil
}


func (g *GrpcClient) CommentService() pc.CommentServiceClient {
	return g.connections["user_service"].(pc.CommentServiceClient)
}
