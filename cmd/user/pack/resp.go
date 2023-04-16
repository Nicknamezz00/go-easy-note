package pack

import (
	"errors"
	"go-easy-note/kitex_gen/user"
	"go-easy-note/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *user.BaseResp {
	if err == nil {
		return baseResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResponse(e)
	}

	se := errno.ServiceErr.WithMessage(err.Error())
	return baseResponse(se)
}

func baseResponse(err errno.ErrNo) *user.BaseResp {
	return &user.BaseResp{
		StatusCode: err.ErrCode,
		Message:    err.ErrMsg,
		Timestamp:  time.Now().Unix(),
	}
}
