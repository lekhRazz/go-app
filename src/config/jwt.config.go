package config

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	SessionId string `json:"sessionId"`
	jwt.RegisteredClaims
}

type SessionInfo struct {
	SessionId string `json:"sessionId"`
	UserId    int    `json:"userId"`
	UserUUID  string `json:"userUUID"`
	Role      string `json:"role"`
}
