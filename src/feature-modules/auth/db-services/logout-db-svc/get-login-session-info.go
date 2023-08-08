package logoutDBSvc

import (
	"database/sql"
)

type LoginUserInfo struct {
	Id     int
	Uuid   string
	Role   string
	Secret string
}

func (proto *dbproto) GetloginSession(sessionId string) (error, LoginUserInfo) {
	query := `SELECT 
							admins.id,
							admins.uuid,
							roles.code AS role,
							als.session_secret as secret
						FROM 
							admins
						INNER JOIN 
							roles
						ON roles.id = admins.role_id
						INNER JOIN
							admin_login_session als
						ON admins.id = als.user_id
						WHERE 
							als.uuid = $1 
						AND 
							als.deleted = false
						AND 
							admins.deleted = false 
						AND
							admins.verified = true`

	var user LoginUserInfo

	row := proto.db.QueryRow(query, sessionId)

	if err := row.Scan(&user.Id, &user.Uuid, &user.Role, &user.Secret); err != nil && err != sql.ErrNoRows {
		return err, user
	}
	return nil, user
}
