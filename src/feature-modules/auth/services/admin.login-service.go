package authService

import (
	"fmt"
	"net/http"
	"sample_go_app/src/config"
	"sample_go_app/src/utilities"

	"database/sql"

	logindbSvc "sample_go_app/src/feature-modules/auth/db-services/login-db-svc"

	"sample_go_app/src/feature-modules/auth/utils"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

var pool *sql.DB

const (
	loginSuccessMessage                    = "Login success"
	loginFailedMessage                     = "Login failed. Check login credentials."
	failedLoginAttemptAndAccBlockedMessage = "Account blocked because of multiple failed login attempt."
	accountBlockedMessage                  = "Your account is blocked"
	attemptRemainingMessage                = "Login failed. %d attempt remaining"
	maxFailedLoginAttemptCount             = 5
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type HttpResponse struct {
	Status  int          `json:"status"`
	Token   string       `json:"token"`
	Message string       `json:"message"`
	Debug   config.Debug `json:"debug"`
}

/*
* @author lekhrazz
* @param {Context} c
* @return {Error} error
* @return {HTTPResposne}
 */

func AdminLoginHandler(c echo.Context) error {

	//@bind login request to struct
	loginInfo := new(LoginRequest)
	err := c.Bind(loginInfo)
	checkerror(err)
	loginRequest := *&loginInfo

	debug := c.Get("debug").(config.Debug)
	response := &HttpResponse{Debug: debug}

	//@check if login data is valid or not
	if loginInfo == nil || len(loginRequest.Email) == 0 || len(loginRequest.Password) == 0 {
		response.Message = loginFailedMessage
		response.Status = http.StatusBadRequest
		response.Token = ""
		return c.JSON(http.StatusBadRequest, response)
	}

	//@get db connection pool
	pool := c.Get("dbConnection").(config.DBConnection).DBConnection

	dbSvcs := logindbSvc.NewLoginDBSvcs(pool)

	//@get user info from db
	err, adminInfo := dbSvcs.GetUserInfo(loginRequest.Email)
	checkerror(err)

	//@check user exits or not, if not exists then return 400 status
	if adminInfo == (logindbSvc.LoginUserInfo{}) {
		response.Message = loginFailedMessage
		response.Status = http.StatusBadRequest
		response.Token = ""
		return c.JSON(http.StatusBadRequest, response)
	}

	//@verify user password if failed then count failed login and suspend/block account if failed login limit crosses
	if matchPassword := verifyPassword(adminInfo.Password, loginRequest.Password); matchPassword == false {
		err, failedLoginAttempt := dbSvcs.GetFailedLoginAttempt(adminInfo.Id)
		checkerror(err)

		var totalFailedLoginAttempt = failedLoginAttempt + 1

		//@update failed login attempt
		if failedLoginAttempt < maxFailedLoginAttemptCount {
			err := dbSvcs.SaveFailedLoginAttempt(adminInfo.Id, debug)
			checkerror(err)
		}

		//@block user
		if totalFailedLoginAttempt >= maxFailedLoginAttemptCount {
			dbSvcs.BlockUser(adminInfo.Id, failedLoginAttemptAndAccBlockedMessage)
			response.Message = failedLoginAttemptAndAccBlockedMessage
			response.Status = http.StatusBadRequest
			response.Token = ""
			return c.JSON(http.StatusBadRequest, response)
		}

		response.Message = fmt.Sprintf(attemptRemainingMessage, (maxFailedLoginAttemptCount - totalFailedLoginAttempt))
		response.Status = http.StatusBadRequest
		response.Token = ""
		return c.JSON(http.StatusBadRequest, response)
	}

	if isValid, validationMsg := utils.ValidateUser(adminInfo); isValid == false {
		response.Message = validationMsg
		response.Status = http.StatusBadRequest
		response.Token = ""
		return c.JSON(http.StatusBadRequest, response)
	}

	//@generate token
	secret, err := utilities.GenerateUUID()
	checkerror(err)

	//@generate session-id
	sessionId, err := utilities.GenerateUUID()
	checkerror(err)

	//@generate jwt token
	jwt, err := utilities.GenerateJwtToken(sessionId, secret)
	checkerror(err)

	//@delete failed logins attempts
	err = dbSvcs.DeleteFailedLoginAttempts(adminInfo.Id)
	checkerror(err)

	//@save login session
	err = dbSvcs.SaveLoginSession(sessionId, adminInfo.Id, debug, jwt, secret)
	checkerror(err)

	response.Message = loginSuccessMessage
	response.Status = http.StatusOK
	response.Token = jwt
	return c.JSON(http.StatusOK, response)
}

func checkerror(err error) {
	if err != nil {
		panic(err)
	}
}

func verifyPassword(hashPwd, plainPassword string) bool {
	return true
}
