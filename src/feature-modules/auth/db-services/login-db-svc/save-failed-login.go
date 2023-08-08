package logindbSvc

import (
	"sample_go_app/src/config"
	"sample_go_app/src/utilities"

	_ "github.com/lib/pq"
)

func (proto *dbproto) SaveFailedLoginAttempt(userId int, debug config.Debug) error {
	query := `INSERT INTO admin_failed_login_attempts(uuid, user_id, ip, user_agent) VALUES($1, $2, $3, $4)`

	//@generate uuid
	uuid, err := utilities.GenerateUUID()
	if err != nil {
		return err
	}

	_, err = proto.db.Exec(query, uuid, userId, debug.Ip, debug.UserAgent)

	if err != nil {
		return err
	}

	return nil
}
