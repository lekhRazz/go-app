package onboardService

type DefaultMessageConfig struct {
	SignUpSuccess       string
	SignUpFailed        string
	EmailVerifySuccess  string
	EmailVerifyFailed   string
	MobileVerifySuccess string
	MobileVerifyFailed  string
}

var DefaultMessageConfigData = DefaultMessageConfig{
	SignUpSuccess:       "Signup success.",
	SignUpFailed:        "Signup failed.",
	EmailVerifySuccess:  "Email verify success.",
	EmailVerifyFailed:   "Email verify failed.",
	MobileVerifySuccess: "Mobile verify success.",
	MobileVerifyFailed:  "Mobile verify failed.",
}
