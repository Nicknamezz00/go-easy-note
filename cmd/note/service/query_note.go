package service

import (
	"context"
	"go-easy-note/cmd/note/dal/db"
	"go-easy-note/cmd/note/pack"
	"go-easy-note/cmd/note/rpc"
	"go-easy-note/kitex_gen/note"
	"go-easy-note/kitex_gen/user"
)

type QueryNoteService struct {
	ctx context.Context
}

func NewQueryNoteService(ctx context.Context) *QueryNoteService {
	return &QueryNoteService{ctx: ctx}
}

func (s *QueryNoteService) QueryNote(req *note.QueryNoteRequest) ([]*note.Note, int64, error) {
	noteModels, total, err := db.QueryNote(s.ctx, req.UserId, req.SearchKey, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, 0, err
	}
	umap, err := rpc.GetUsers(s.ctx, &user.GetUsersRequest{UserIds: []int64{req.UserId}})
	if err != nil {
		return nil, 0, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := umap[notes[i].UserId]; ok {
			notes[i].UserName = u.UserName
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, total, nil
}
