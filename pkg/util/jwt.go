package util

import (
	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type PayLoad struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Scope    string `json:"scope"`
	IsSuper  bool   `json:"is_super"`
}

type Claims struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Scope    string `json:"scope"`
	IsSuper  bool   `json:"is_super"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(payload PayLoad, expireTime int64) (string, error) {

	claims := Claims{
		payload.ID,
		payload.Account,
		EncodeMD5(payload.Password),
		payload.Scope,
		payload.IsSuper,
		jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "liaoliao",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken parsing token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
