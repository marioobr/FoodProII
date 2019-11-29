package models

import "github.com/dgrijalva/jwt-go"

type JwtClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Names string `json:"names"`
	LastNames string `json:"lastnames"`
	NextStep string `json:"next_step"`
	jwt.StandardClaims
}