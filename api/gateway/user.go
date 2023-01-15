package gateway

import (
	"doc-api/api/usecase"

	"errors"

	"gorm.io/gorm"
)

type (
	userRepository struct {
		db *gorm.DB
	}

	User struct {
		Id       *int
		UserId   string
		Password string
		Admin    bool
	}
)

func NewUserRepository(db *gorm.DB) usecase.LoginRepository {
	return &userRepository{
		db: db,
	}
}

func NewSignupRepository(db *gorm.DB) usecase.SignupRepository {
	return &userRepository{
		db: db,
	}
}

func NewAuthRepository(db *gorm.DB) usecase.AuthRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) ValidateUser(id string, pass string) (ok bool, err error) {
	user := User{}
	if err := r.db.Where("user_id = ? and password = ?", id, pass).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepository) ExistUser(id string) (exist bool, err error) {
	user := User{}
	if err := r.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepository) IsAdmin(id string) (isAdmin bool, err error) {
	user := User{}
	if err := r.db.Where("user_id = ?", id).Find(&user).Error; err != nil {
		return false, err
	}

	return user.Admin, nil
}

func (r *userRepository) CreateUser(id string, pass string) error {
	user := User{
		UserId:   id,
		Password: pass,
		Admin:    false,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
