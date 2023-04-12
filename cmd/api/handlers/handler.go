package handlers

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"go-easy-note/pkg/errno"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserParam struct {
	UserName string `json:"user-name"`
	Password string `json:"password"`
}

type NoteParam struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func SendResponse(c *app.RequestContext, err error, data interface{}) {
	e := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, Response{
		Code:    e.ErrCode,
		Message: e.ErrMsg,
		Data:    data,
	})
}
