package postgres

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"
)

type CommentRepo struct {
	db *sqlx.DB
}

func NewCommentRepo(db *sqlx.DB) *CommentRepo {
	return &CommentRepo{db: db}
}
func (r *CommentRepo) CreateComment(req *pc.CreateCommentReq) (*pc.GetCommentRes, error) {

	res := pc.GetCommentRes{}
	query := `INSERT INTO comments (comment_id, user_id, post_id, text, created_at) 
				VALUES($1,$2,$3,$4,$5) 
				RETURNING comment_id, user_id, post_id, text, created_at`

	now := time.Now().Format(time.RFC3339)
	err := r.db.QueryRow(query, uuid.New(), req.UserId, req.PostId, req.Text, now).Scan(
		&res.CommentId,
		&res.UserId,
		&res.PostId,
		&res.Text,
		&res.CreatedAt,
	)
	if err != nil {
		log.Println("Error while insert comment: ", err)
		return nil, err
	}

	return &res, nil
}
func (r *CommentRepo) GetComment(req *pc.GetCommentReq) (*pc.GetCommentRes, error) {

	res := pc.GetCommentRes{}
	query := `SELECT comment_id, post_id, user_id, text, created_at FROM comments WHERE post_id=$1 AND deleted_at IS NULL`
	err := r.db.QueryRow(query, req.PostId).Scan(
		&res.CommentId,
		&res.PostId,
		&res.UserId,
		&res.Text,
		&res.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return &pc.GetCommentRes{CommentId: ""}, nil
	}
	if err != nil {
		log.Println("Error while Get comment", err)
		return nil, err
	}
	return &res, nil
}
func (r *CommentRepo) UpdateComment(req *pc.UpdateCommentReq) (*pc.GetCommentRes, error) {
	var (
		PostId string
	)

	query := `UPDATE comments SET text=$2 WHERE comment_id=$3 AND user_id=$1 AND deleted_at IS NULL RETURNING post_id`
	err := r.db.QueryRow(query, req.UserId, req.Text, req.CommentId).Scan(
		&PostId,
	)
	if err != nil {
		log.Println("Error while Update comment", err)
		return nil, err
	}
	return r.GetComment(&pc.GetCommentReq{PostId: PostId})
}
func (r *CommentRepo) DeleteComment(req *pc.DeleteCommentReq) (*pc.Message, error) {
	now := time.Now().Format(time.RFC3339)
	query := `UPDATE comments SET deleted_at=$1 WHERE comment_id=$2 AND user_id=$3 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, now, req.CommentId, req.UserId)
	if err == sql.ErrNoRows {
		return &pc.Message{}, err
	}
	if err != nil {
		return nil, err
	}
	return &pc.Message{Message: "Deleted!"}, nil
}
