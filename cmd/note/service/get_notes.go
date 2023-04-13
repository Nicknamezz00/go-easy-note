package service

import (
	"context"
	"go-easy-note/cmd/note/dal/db"
	"go-easy-note/cmd/note/pack"
	"go-easy-note/cmd/note/rpc"
	"go-easy-note/kitex_gen/note"
	"go-easy-note/kitex_gen/user"
)

type GetNotesService struct {
	ctx context.Context
}

func NewGetNotesService(ctx context.Context) *GetNotesService {
	return &GetNotesService{ctx: ctx}
}

func (s *GetNotesService) GetNotes(req *note.GetNotesRequest) ([]*note.Note, error) {
	noteModels, err := db.GetNotes(s.ctx, req.NoteIds)
	if err != nil {
		return nil, err
	}
	uIDs := pack.UserIDs(noteModels)
	umap, err := rpc.GetUsers(s.ctx, &user.GetUsersRequest{UserIds: uIDs})
	if err != nil {
		return nil, err
	}
	notes := pack.Notes(noteModels)
	for i := 0; i < len(notes); i++ {
		if u, ok := umap[notes[i].UserId]; ok {
			notes[i].UserName = u.UserName
			notes[i].UserAvatar = u.Avatar
		}
	}
	return notes, nil
}
