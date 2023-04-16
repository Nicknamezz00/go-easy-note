package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"go-easy-note/cmd/user/dal/db"
	"go-easy-note/kitex_gen/user"
	"go-easy-note/pkg/errno"
	"io"
)

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

func (s *CheckUserService) CheckUser(req *user.CheckUserRequest) (int64, error) {
	hs := md5.New()
	if _, err := io.WriteString(hs, req.Password); err != nil {
		return 0, err
	}

	pwd := fmt.Sprintf("%x", hs.Sum(nil))
	uname := req.UserName
	users, err := db.QueryUser(s.ctx, uname)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.AuthorizationFailedErr
	}
	if users[0].Password != pwd {
		return 0, errno.AuthorizationFailedErr
	}
	return int64(users[0].ID), nil
}
