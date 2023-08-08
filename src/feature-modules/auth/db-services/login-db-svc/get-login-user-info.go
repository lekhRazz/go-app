package logindbSvc

import (
	"database/sql"
)

type LoginUserInfo struct {
	Id            int
	Uuid          string
	Password      string
	Email         string
	Phone         string
	Role          string
	Suspended     bool
	SuspendReason string
}

func (proto *dbproto) GetUserInfo(userEmail string) (error, LoginUserInfo) {
	query := `SELECT 
							admins.id,
							admins.uuid,
							admins.password,
							admins.email,
							admins.phone,
							roles.code AS role,
							admins.suspended,
							COALESCE(admins.suspend_reason, ' ') as suspendReason
						FROM 
							admins
						INNER JOIN 
							roles
						ON roles.id = admins.role_id
						WHERE 
							admins.email = $1 
						AND 
							admins.deleted = false 
						AND
							admins.verified = true`

	var user LoginUserInfo

	row := proto.db.QueryRow(query, userEmail)

	if err := row.Scan(&user.Id, &user.Uuid, &user.Password, &user.Email, &user.Phone, &user.Role, &user.Suspended, &user.SuspendReason); err != nil && err != sql.ErrNoRows {
		return err, user
	}
	return nil, user
}
