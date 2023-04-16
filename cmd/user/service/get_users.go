package service

import (
	"context"
	"go-easy-note/cmd/user/dal/db"
	"go-easy-note/cmd/user/pack"
	"go-easy-note/kitex_gen/user"
)

type GetUsersService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUsersService {
	return &GetUsersService{ctx: ctx}
}

func (s *GetUsersService) GetUsers(req *user.GetUsersRequest) ([]*user.User, error) {
	userModels, err := db.GetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(userModels), nil
}
