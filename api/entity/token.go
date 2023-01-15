package entity

import (
	"doc-api/env"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	JwtToken interface {
		Generate(id string, t *time.Time) (token string, err error)
		Validate(tokenString string) (id string, err error)
		ExtractUserId(token *jwt.Token) (id string, err error)
	}
	jwtToken struct {
		secret string
	}
)

const (
	map_userid = "user_id"
	map_exp    = "exp"
)

func NewJwtToken(secret env.JWT_SECRET_KEY) JwtToken {
	return &jwtToken{
		secret: string(secret),
	}
}

func (j *jwtToken) Generate(id string, t *time.Time) (token string, err error) {
	// jwtトークンの生成
	claims := jwt.MapClaims{
		map_userid: id,
		map_exp:    t.Add(time.Hour * 24 * 7).Unix(),
	}

	tokenStr := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return tokenStr.SignedString([]byte(j.secret))
}

func (j *jwtToken) ExtractUserId(token *jwt.Token) (id string, err error) {
	c := token.Claims.(jwt.MapClaims)
	id = c[map_userid].(string)

	return id, nil
}

func (j *jwtToken) Validate(tokenString string) (id string, err error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(j.secret), nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	if e := token.Claims.Valid(); e != nil {
		return "", e
	}

	c := token.Claims.(jwt.MapClaims)
	id = c[map_userid].(string)

	return id, nil
}
