package main

import (
	"gin_boot/pkg/jwts"
	"log"
	"testing"
)

var jwt = jwts.NewJWTHandler()

func TestLogin(t *testing.T) {
	// 模拟用户信息
	userID := int64(1)
	username := "admin"

	// 生成 JWT Token
	token, err := jwt.SetJWTToken(userID, username)
	if err != nil {
		log.Panicln(err.Error())
	}
	log.Println(token)
}

func TestGetJwt(t *testing.T) {
	//token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhZG1pbiIsImV4cCI6MTc1NTQ2MzIzOSwibmJmIjoxNzU1Mzc2ODM5LCJpYXQiOjE3NTUzNzY4MzksInVzZXJfaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.gYN8fDVsxZ4Cs4jJIWyn4JjiKPloe5B5aHOmc8FYSqE"
	//jwt.ParseToken()
}
