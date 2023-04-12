package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"go-easy-note/cmd/api/rpc"
	"go-easy-note/kitex_gen/user"
	"go-easy-note/pkg/errno"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var registerVal UserParam
	if err := c.Bind(&registerVal); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(registerVal.UserName) == 0 || len(registerVal.Password) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err := rpc.CreateUser(context.Background(), &user.CreateUserRequest{
		UserName: registerVal.UserName,
		Password: registerVal.Password,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}
