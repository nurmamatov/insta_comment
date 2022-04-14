package main

import (
	"net"

	"tasks/Instagram_clone/insta_comment/config"
	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"
	"tasks/Instagram_clone/insta_comment/pkg/db"
	"tasks/Instagram_clone/insta_comment/pkg/logger"
	"tasks/Instagram_clone/insta_comment/service"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config := config.Load()

	log := logger.New(config.LogLevel, "comment_service")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", config.PostgresHost),
		logger.Int("port", config.PostgresPort),
		logger.String("database", config.PostgresDatabase),
		logger.String("password", config.PostgresPassword))
	connDB, err := db.ConnectToDB(config)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	commentService := service.NewCommentService(connDB, log)

	lis, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pc.RegisterCommentServiceServer(s, commentService)
	reflection.Register(s)
	log.Info("main: server running",
		logger.String("port", config.Port))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
