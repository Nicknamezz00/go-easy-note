package db

import (
	"context"
	"go-easy-note/pkg/constants"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	UserID  int64  `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (n *Note) TableName() string {
	return constants.NoteTableName
}

// CreateNote create note
func CreateNote(ctx context.Context, notes *[]Note) error {
	if err := DB.WithContext(ctx).Create(notes).Error; err != nil {
		return err
	}
	return nil
}

// GetNotes get list of multiple notes
func GetNotes(ctx context.Context, noteIDs []int64) ([]*Note, error) {
	var ret []*Note
	if len(noteIDs) == 0 {
		return ret, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", noteIDs).Find(&ret).Error; err != nil {
		return ret, err
	}
	return ret, nil
}

// UpdateNote update note
func UpdateNote(ctx context.Context, noteID, userID int64, title, content *string) error {
	params := map[string]interface{}{}
	if title != nil {
		params["title"] = *title
	}
	if content != nil {
		params["content"] = *content
	}
	return DB.WithContext(ctx).Model(&Note{}).
		Where("id = ? and user_id = ?", noteID, userID).
		Updates(params).Error
}

// DeleteNote delete note
func DeleteNote(ctx context.Context, noteID, userID int64) error {
	return DB.WithContext(ctx).Where("id = ? and user_id = ?", noteID, userID).Delete(&Note{}).Error
}

func QueryNote(ctx context.Context, userID int64, searchKey *string, limit, offset int) ([]*Note, int64, error) {
	var ret []*Note
	var total int64
	conn := DB.WithContext(ctx).Model(&Note{}).Where("user_id = ?", userID)

	if searchKey != nil {
		conn.Where("title like ?", "%"+*searchKey+"%")
	}

	if err := conn.Count(&total).Error; err != nil {
		return ret, total, err
	}
	if err := conn.Limit(limit).Offset(offset).Find(&ret).Error; err != nil {
		return ret, total, err
	}
	return ret, total, nil
}
