package db

import (
	"context"
	"go-easy-note/pkg/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

// CreateUser create user
func CreateUser(ctx context.Context, users []*User) error {
	return DB.WithContext(ctx).Create(users).Error
}

// GetUsers get list of multiple users
func GetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	ret := make([]*User, 0)
	if len(userIDs) == 0 {
		return ret, nil
	}

	if err := DB.WithContext(ctx).Where("id in ?", userIDs).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}

func QueryUser(ctx context.Context, username string) ([]*User, error) {
	ret := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("user_name = ?", username).Find(&ret).Error; err != nil {
		return nil, err
	}
	return ret, nil
}
