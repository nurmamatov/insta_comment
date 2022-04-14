package main

import (
	"fmt"
	"log"
	"net"

	"tasks/Instagram_clone/insta_comment/config"
	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"
	"tasks/Instagram_clone/insta_comment/service"

	grpcClient "tasks/Instagram_clone/insta_comment/service/grpc_client"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.Load()

	psqlText := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.PostgresHost,
		config.PostgresPort,
		config.PostgresUser,
		config.PostgresPassword,
		config.PostgresDatabase,
	)

	grpcClient, err := grpcClient.New(config)
	if err != nil {
		log.Fatal("grpc dial error", err)
	}

	connDB, err := sqlx.Connect("postgres", psqlText)
	if err != nil {
		log.Fatal(err)
	}

	CommentService := service.NewCommentService(connDB, grpcClient)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatal("Error while listening:", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pc.RegisterCommentServiceServer(s, CommentService)
	log.Println("Main server runnning", config.Port)

	if err = s.Serve(lis); err != nil {
		log.Fatal("Error while listening:", err)
	}
}
