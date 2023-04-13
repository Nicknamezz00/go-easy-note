package service

import (
	"context"
	"go-easy-note/cmd/note/dal/db"
	"go-easy-note/kitex_gen/note"
)

type DeleteNoteService struct {
	ctx context.Context
}

func NewDeleteNoteService(ctx context.Context) *DeleteNoteService {
	return &DeleteNoteService{ctx: ctx}
}

func (s *DeleteNoteService) DeleteNote(req *note.DeleteNoteRequset) error {
	return db.DeleteNote(s.ctx, req.NoteId, req.UserId)
}
