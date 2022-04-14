package storage

import (
	"tasks/Instagram_clone/insta_comment/storage/postgres"
	"tasks/Instagram_clone/insta_comment/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Comment() repo.Comment
}
type storagePg struct {
	db          *sqlx.DB
	commentRepo repo.Comment
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		commentRepo: postgres.NewCommentRepo(db),
	}
}

func (s storagePg) Comment() repo.Comment {
	return s.commentRepo
}
