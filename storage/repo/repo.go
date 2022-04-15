package repo

import (
	pc "tasks/Instagram_clone/insta_comment/genproto/comment_proto"
)

type Comment interface {
	CreateComment(*pc.CreateCommentReq) (*pc.Res, error)
	GetComment(*pc.GetCommentReq) (*pc.ResG, error)
	UpdateComment(*pc.UpdateCommentReq) (*pc.ResG, error)
	DeleteComment(*pc.DeleteCommentReq) (*pc.Message, error)
}
