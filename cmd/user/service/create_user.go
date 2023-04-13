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

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *user.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}
	hs := md5.New()
	if _, err := io.WriteString(hs, req.Password); err != nil {
		return err
	}
	pwd := fmt.Sprintf("%x", hs.Sum(nil))
	return db.CreateUser(s.ctx, []*db.User{{
		UserName: req.UserName,
		Password: pwd,
	}})
}
