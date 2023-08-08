package logindbSvc

func (proto *dbproto) BlockUser(userId int, suspendMessage string) error {
	query := `update admins set suspended = true, suspend_reason=$1 where id = $2`

	_, err := proto.db.Exec(query, suspendMessage, userId)

	if err != nil {
		return err
	}
	return nil
}
