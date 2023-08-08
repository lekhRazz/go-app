package config

import (
	"database/sql"
)

type TablesConfig struct {
	UsersTable                   string
	UsersOnboardingTable         string
	UsersEmailVerificationTable  string
	UsersMobileVerificationTable string
	UserFailedLoginAttemptTable  string
	UserLoginSessionTable        string
}

var DefaultTablesConfigData = TablesConfig{
	UsersTable:                   "users",
	UsersOnboardingTable:         "users_onboarding",
	UsersEmailVerificationTable:  "users_email_verification",
	UsersMobileVerificationTable: "users_mobile_verification",
	UserLoginSessionTable:        "user_login_session",
	UserFailedLoginAttemptTable:  "user_failed_login_attempt",
}

type DBConnection struct {
	DBConnection *sql.DB
}
