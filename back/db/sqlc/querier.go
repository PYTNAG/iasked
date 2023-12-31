// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"context"
)

type Querier interface {
	CountRFCComments(ctx context.Context, rfcID int32) (int64, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (int32, error)
	CreateRFC(ctx context.Context, arg CreateRFCParams) (int32, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	DeleteComment(ctx context.Context, commentID int32) error
	DeleteRFC(ctx context.Context, rfcID int32) error
	DeleteUser(ctx context.Context, userID int32) error
	GetLastRFCs(ctx context.Context, arg GetLastRFCsParams) ([]Rfc, error)
	GetRFCComments(ctx context.Context, arg GetRFCCommentsParams) ([]Comment, error)
	GetUserHash(ctx context.Context, userID int32) ([]byte, error)
	GetUserInfo(ctx context.Context, userID int32) (GetUserInfoRow, error)
	UpdateEmail(ctx context.Context, arg UpdateEmailParams) error
	UpdateHash(ctx context.Context, arg UpdateHashParams) error
	UpdateUsername(ctx context.Context, arg UpdateUsernameParams) error
}

var _ Querier = (*Queries)(nil)
