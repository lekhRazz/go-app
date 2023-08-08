package logindbSvc

import (
	"sample_go_app/src/config"
	// _ "github.com/lib/pq"
)

func (proto *dbproto) SaveLoginSession(sessionId string, userId int, debug config.Debug, token string, sessionSecret string) error {
	query := `INSERT INTO admin_login_session(uuid, user_id, ip, user_agent, token, session_secret) VALUES($1, $2, $3, $4, $5, $6) returning id`

	_, err := proto.db.Exec(query, sessionId, userId, debug.Ip, debug.UserAgent, token, sessionSecret)
	if err != nil {
		return err
	}

	return nil
}
