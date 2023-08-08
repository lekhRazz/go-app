package logindbSvc

import "database/sql"

func (proto *dbproto) GetFailedLoginAttempt(userId int) (error, int) {
	query := `SELECT COUNT(id) as count FROM admin_failed_login_attempts WHERE user_id = $1 AND deleted = false`

	var count int = 0

	row := proto.db.QueryRow(query, userId)

	if err := row.Scan(&count); err != nil && err != sql.ErrNoRows {
		return err, count
	}
	return nil, count
}
