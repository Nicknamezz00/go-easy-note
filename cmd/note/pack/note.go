package pack

import (
	"go-easy-note/cmd/note/dal/db"
	"go-easy-note/kitex_gen/note"
)

// Note pack note info
func Note(n *db.Note) *note.Note {
	if n == nil {
		return nil
	}
	return &note.Note{
		NoteId:     int64(n.ID),
		UserId:     n.UserID,
		Title:      n.Title,
		Content:    n.Content,
		CreateTime: n.CreatedAt.Unix(),
	}
}

// Notes pack list of notes
func Notes(ns []*db.Note) []*note.Note {
	notes := make([]*note.Note, 0)
	for _, n := range ns {
		if nt := Note(n); nt != nil {
			notes = append(notes, nt)
		}
	}
	return notes
}

// UserIDs pack list of notes to list of user ids
func UserIDs(ns []*db.Note) []int64 {
	uIDs := make([]int64, 0)
	if len(ns) == 0 {
		return uIDs
	}
	umap := make(map[int64]struct{})
	for _, n := range ns {
		if n != nil {
			umap[n.UserID] = struct{}{}
		}
	}
	for uID := range umap {
		uIDs = append(uIDs, uID)
	}
	return uIDs
}
