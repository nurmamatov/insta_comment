package repo

import (
	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"
)

type Comment interface {
	CreateComment(*pc.CreateCommentReq) (*pc.GetCommentRes, error)
	GetComment(*pc.GetCommentReq) (*pc.GetCommentRes, error)
	UpdateComment(*pc.UpdateCommentReq) (*pc.GetCommentRes, error)
	DeleteComment(*pc.DeleteCommentReq) (*pc.Message, error)
}
