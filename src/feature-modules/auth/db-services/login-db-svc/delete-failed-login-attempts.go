package logindbSvc

func (proto *dbproto) DeleteFailedLoginAttempts(userId int) error {
	query := `update admin_failed_login_attempts set deleted=true where user_id = $1`

	_, err := proto.db.Exec(query, userId)

	if err != nil {
		return err
	}
	return nil
}
