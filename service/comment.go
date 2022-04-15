package service

import (
	"context"
	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"
	l "tasks/Instagram_clone/insta_comment/pkg/logger"

	"tasks/Instagram_clone/insta_comment/storage"

	"github.com/jmoiron/sqlx"
)

type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewCommentService(db *sqlx.DB, log l.Logger) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (r *CommentService) CreateComment(ctx context.Context, req *pc.CreateCommentReq) (*pc.Res, error) {
	res, err := r.storage.Comment().CreateComment(req)
	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (r *CommentService) GetComment(ctx context.Context, req *pc.GetCommentReq) (*pc.ResG, error) {
	res, err := r.storage.Comment().GetComment(req)

	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
func (r *CommentService) UpdateComment(ctx context.Context, req *pc.UpdateCommentReq) (*pc.ResG, error) {
	res, err := r.storage.Comment().UpdateComment(req)

	if err != nil {
		r.logger.Error("Error: ", l.Error(err))
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
		r.logger.Error("Error: ", l.Error(err))
		return nil, err
	}
	return res, nil
}
