package usecase

import (
	"doc-api/api/entity"
	"errors"
)

type (
	SignupService interface {
		Signup(id string, pass string, out SignupOutputPort) error
	}

	SignupOutputPort interface {
		Success() error
		Failure(err error) error
	}

	SignupRepository interface {
		ExistUser(id string) (exist bool, err error)
		CreateUser(id string, pass string) error
	}

	signupService struct {
		hashing entity.Hashing
		repo    SignupRepository
	}
)

func NewSignupService(hashing entity.Hashing, repo SignupRepository) SignupService {
	return &signupService{
		hashing: hashing,
		repo:    repo,
	}
}

func (s *signupService) Signup(id string, pass string, out SignupOutputPort) error {
	exist, err := s.repo.ExistUser(id)
	if err != nil {
		out.Failure(err)
	}
	if exist {
		// ID重複エラー
		out.Failure(errors.New("id exists"))
	}
	hashedPass := s.hashing.Hash(pass)
	if err := s.repo.CreateUser(id, hashedPass); err != nil {
		// アカウント作成失敗エラー
		out.Failure(err)
	}

	return out.Success()

}
