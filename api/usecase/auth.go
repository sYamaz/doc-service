package usecase

import (
	"doc-api/api/entity"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type (
	AuthUser interface {
		Auth(j *jwt.Token, out AuthOutputPort) error
		AuthAsAdmin(j *jwt.Token, out AuthOutputPort) error
	}
	auth struct {
		t   entity.JwtToken
		rep AuthRepository
	}

	AuthRepository interface {
		ExistUser(userId string) (exist bool, err error)
		IsAdmin(userId string) (isAdmin bool, err error)
	}

	AuthOutputPort interface {
		Success(userId string) error
		Failure(err error) error
	}
)

func NewAuthUserService(t entity.JwtToken, rep AuthRepository) AuthUser {
	return &auth{
		t:   t,
		rep: rep,
	}
}

func (a *auth) Auth(j *jwt.Token, out AuthOutputPort) error {
	userId, err := a.t.ExtractUserId(j)
	if err != nil {
		return out.Failure(err)
	}

	exist, err := a.rep.ExistUser(userId)
	if err != nil {
		return out.Failure(err)
	}
	if !exist {
		return out.Failure(errors.New("invalid user_id"))
	}

	return out.Success(userId)
}

func (a *auth) AuthAsAdmin(j *jwt.Token, out AuthOutputPort) error {
	userId, err := a.t.ExtractUserId(j)
	if err != nil {
		return out.Failure(err)
	}

	exist, err := a.rep.ExistUser(userId)
	if err != nil {
		return out.Failure(err)
	}
	if !exist {
		return out.Failure(errors.New("invalid user_id"))
	}

	isAdmin, err := a.rep.IsAdmin(userId)
	if err != nil {
		return out.Failure(err)
	}
	if !isAdmin {
		return out.Failure(errors.New("user is not admin"))
	}

	return out.Success(userId)
}
