package logoutDBSvc

func (proto *dbproto) LogoutUser(userId int, sessionId string) error {
	query := `update admin_login_session set deleted=true where uuid = $1 AND user_id = $2`

	_, err := proto.db.Exec(query, sessionId, userId)

	if err != nil {
		return err
	}
	return nil
}
