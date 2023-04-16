package constants

const (
	IdentityKey             = "id"
	SecretKey               = "secret key"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	ApiServiceName          = "api"
	NoteServiceName         = "note"
	UserServiceName         = "user"
	UserTableName           = "user"
	NoteTableName           = "note"
	MySQLDefaultDSN         = "wrz:wrz@tcp(localhost:9901)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "127.0.0.1:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10
)
