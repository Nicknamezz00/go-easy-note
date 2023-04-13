package pack

import (
	"errors"
	"go-easy-note/kitex_gen/note"
	"go-easy-note/pkg/errno"
	"time"
)

func BuildBaseResp(err error) *note.BaseResp {
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

func baseResponse(err errno.ErrNo) *note.BaseResp {
	return &note.BaseResp{
		StatusCode: err.ErrCode,
		Message:    err.ErrMsg,
		Timestamp:  time.Now().Unix(),
	}
}
