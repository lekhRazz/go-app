package logindbSvc

import (
	"database/sql"
	"sample_go_app/src/config"
)

type dbproto struct {
	db *sql.DB
}

type LoginDBSvc interface {
	GetUserInfo(userEmail string) (error, LoginUserInfo)
	SaveLoginSession(sessionId string, userId int, debug config.Debug, token string, sessionSecret string) error
	GetFailedLoginAttempt(userId int) (error, int)
	SaveFailedLoginAttempt(userId int, debug config.Debug) error
	BlockUser(userId int, suspendMessage string) error
	DeleteFailedLoginAttempts(userId int) error
}

func NewLoginDBSvcs(dbConn *sql.DB) LoginDBSvc {
	return &dbproto{db: dbConn}
}
