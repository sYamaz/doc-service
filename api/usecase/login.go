package usecase

import (
	"doc-api/api/entity"
	"time"
)

type (
	LoginService interface {
		Login(id string, pass string, out LoginOutputPort) error
	}

	LoginOutputPort interface {
		Success(token string, admin bool, err error) error
		InfoMissmatched(err error) error
		Failure(err error) error
	}

	LoginRepository interface {
		ValidateUser(id string, pass string) (ok bool, err error)
		IsAdmin(id string) (admin bool, err error)
	}

	loginService struct {
		hashing entity.Hashing
		token   entity.JwtToken
		repo    LoginRepository
	}
)

func NewLoginService(hashing entity.Hashing, token entity.JwtToken, repo LoginRepository) LoginService {
	return &loginService{
		hashing: hashing,
		token:   token,
		repo:    repo,
	}
}

func (s *loginService) Login(id string, pass string, out LoginOutputPort) error {
	// id, passの認証
	hashedPass := s.hashing.Hash(pass)

	// db照合
	ok, err := s.repo.ValidateUser(id, hashedPass)
	if err != nil {
		return out.Failure(err)
	}
	if !ok {
		return out.InfoMissmatched(err)
	}

	// jwtトークンの生成
	t := time.Now()
	tokenString, err := s.token.Generate(id, &t)
	if err != nil {
		return out.Failure(err)
	}

	// admin情報取得で失敗しても基本的には支障ない
	admin, err := s.repo.IsAdmin(id)

	return out.Success(tokenString, admin, err)
}
