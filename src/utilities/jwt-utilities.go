package utilities

import (
	"database/sql"
	"fmt"
	"sample_go_app/src/config"
	"time"

	logoutDBSvc "sample_go_app/src/feature-modules/auth/db-services/logout-db-svc"

	"github.com/golang-jwt/jwt/v5"
)

// type Claims struct {
// 	SessionId string `json:"sessionId"`
// 	jwt.RegisteredClaims
// }

func GenerateJwtToken(sessionId, secret string) (string, error) {

	secretKey := []byte(secret)
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &config.Claims{
		SessionId: sessionId,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJwtToken(tokenString string, dbConnPool *sql.DB) (sessionInfo config.SessionInfo, err error) {

	dbSvcs := logoutDBSvc.NewLogoutDBSvcs(dbConnPool)
	loginSessionInfo := logoutDBSvc.LoginUserInfo{}
	var sessionId string
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		sessionId = token.Claims.(jwt.MapClaims)["sessionId"].(string)
		err, loginSessionInfo = dbSvcs.GetloginSession(sessionId)
		if err != nil {
			return nil, fmt.Errorf("Error occured when finding login user info:")
		}

		if loginSessionInfo == (logoutDBSvc.LoginUserInfo{}) {
			return nil, fmt.Errorf("Session Info not found:")
		}

		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(loginSessionInfo.Secret), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionInfo.SessionId = sessionId
		sessionInfo.UserId = loginSessionInfo.Id
		sessionInfo.UserUUID = loginSessionInfo.Uuid
		sessionInfo.Role = loginSessionInfo.Role
		return sessionInfo, nil
	} else {
		return sessionInfo, err
	}
}
