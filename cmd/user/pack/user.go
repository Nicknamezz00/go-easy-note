package pack

import (
	"go-easy-note/cmd/user/dal/db"
	"go-easy-note/kitex_gen/user"
)

func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	return &user.User{
		UserId:   int64(u.ID),
		UserName: u.UserName,
		Avatar:   "test avatar",
	}
}

func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if curUser := User(u); curUser != nil {
			users = append(users, curUser)
		}
	}
	return users
}
