package dal

import "go-easy-note/cmd/user/dal/db"

// Init init user data access layer
func Init() {
	db.Init() // MySQL
}
