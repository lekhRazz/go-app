package logoutDBSvc

import (
	"database/sql"
)

type dbproto struct {
	db *sql.DB
}

type LogoutDBSvc interface {
	GetloginSession(userEmail string) (error, LoginUserInfo)
	LogoutUser(userId int, sessionId string) error
}

func NewLogoutDBSvcs(dbConn *sql.DB) LogoutDBSvc {
	return &dbproto{db: dbConn}
}
