package utils

import (
	logindbSvc "sample_go_app/src/feature-modules/auth/db-services/login-db-svc"
)

func ValidateUser(user logindbSvc.LoginUserInfo) (bool, string) {
	if user.Suspended == true {
		return false, user.SuspendReason
	}

	return true, ""
}
