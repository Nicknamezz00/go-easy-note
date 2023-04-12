package handlers

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"go-easy-note/cmd/api/rpc"
	"go-easy-note/kitex_gen/note"
	"go-easy-note/pkg/constant"
	"go-easy-note/pkg/errno"
	"strconv"
)

// CreateNote CreatNote create note
func CreateNote(ctx context.Context, c *app.RequestContext) {
	var noteVal NoteParam
	if err := c.Bind(&noteVal); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if len(noteVal.Content) == 0 || len(noteVal.Title) == 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constant.IdentityKey].(float64))
	err := rpc.CreateNote(context.Background(), &note.CreateNoteRequest{
		Title:   noteVal.Title,
		Content: noteVal.Content,
		UserId:  userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// DeleteNote delete note
func DeleteNote(ctx context.Context, c *app.RequestContext) {
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constant.IdentityKey].(float64))
	noteIDStr := c.Param(constant.NoteID)
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if noteIDInt <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	err = rpc.DeleteNote(context.Background(), &note.DeleteNoteRequset{
		NoteId: noteIDInt,
		UserId: userID,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// UpdateNote update note
func UpdateNote(ctx context.Context, c *app.RequestContext) {
	var noteVal NoteParam
	if err := c.Bind(&noteVal); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constant.IdentityKey].(float64))
	noteIDStr := c.Param(constant.NoteID)
	noteIDInt, err := strconv.ParseInt(noteIDStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	if noteIDInt <= 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &note.UpdateNoteRequest{NoteId: noteIDInt, UserId: userID}
	if len(noteVal.Title) != 0 {
		req.Title = &noteVal.Title
	}
	if len(noteVal.Content) != 0 {
		req.Content = &noteVal.Content
	}
	if err = rpc.UpdateNote(context.Background(), req); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, nil)
}

// QueryNote query list of notes
func QueryNote(ctx context.Context, c *app.RequestContext) {
	claims := jwt.ExtractClaims(ctx, c)
	userID := int64(claims[constant.IdentityKey].(float64))

	var queryVal struct {
		Limit     int64  `json:"limit" form:"limit" query:"limit"`
		Offset    int64  `json:"offset" form:"offset" query:"offset"`
		SearchKey string `json:"search_key" form:"search_key" query:"search_key"`
	}
	if err := c.Bind(&queryVal); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	if queryVal.Limit < 0 || queryVal.Offset < 0 {
		SendResponse(c, errno.ParamErr, nil)
		return
	}

	req := &note.QueryNoteRequest{UserId: userID, Limit: queryVal.Limit, Offset: queryVal.Offset}
	if len(queryVal.SearchKey) != 0 {
		req.SearchKey = &queryVal.SearchKey
	}
	notes, total, err := rpc.QueryNote(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, map[string]interface{}{
		constant.Total: total,
		constant.Notes: notes,
	})
}
