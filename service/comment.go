package service

import (
	"context"
	"log"
	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"

	grpcClient "tasks/Instagram_clone/insta_comment/service/grpc_client"
	"tasks/Instagram_clone/insta_comment/storage"

	"github.com/jmoiron/sqlx"
)

type CommentService struct {
	storage storage.IStorage
	client  grpcClient.GrpcClientI
}

func NewCommentService(db *sqlx.DB, client grpcClient.GrpcClientI) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		client:  client,
	}
}

func (r *CommentService) CreateComment(ctx context.Context, req *pc.CreateCommentReq) (*pc.GetCommentRes, error) {
	res, err := r.storage.Comment().CreateComment(req)
	if err != nil {
		log.Println("Error while Comment.go CreateComment")
		return nil, err
	}
	return res, nil
}
func (r *CommentService) GetComment(ctx context.Context, req *pc.GetCommentReq) (*pc.GetCommentRes, error) {
	res, err := r.storage.Comment().GetComment(req)
	if err != nil && res.CommentId == "" {
		return res, err
	}
	if err != nil {
		log.Println("Erro while get comment in comment.go")
		return nil, err
	}
	return res, nil
}
func (r *CommentService) UpdateComment(ctx context.Context, req *pc.UpdateCommentReq) (*pc.GetCommentRes, error) {
	res, err := r.storage.Comment().UpdateComment(req)
	if err != nil && res.CommentId == "" {
		return res, err
	}
	if err != nil {
		log.Println("Erro while update comment in comment.go")
		return nil, err
	}
	return res, nil
}
func (r *CommentService) DeleteComment(ctx context.Context, req *pc.DeleteCommentReq) (*pc.Message, error) {
	res, err := r.storage.Comment().DeleteComment(req)
	if res.Message == "" {
		return &pc.Message{Message: "Something wrong"}, nil
	}
	if err != nil {
		log.Println("Erro while delete comment in comment.go")
		return nil, err
	}
	return res, nil
}
