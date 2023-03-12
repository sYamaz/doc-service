package gateway

import (
	"doc-api/api/gateway/model"
	"doc-api/api/usecase"

	"errors"

	"gorm.io/gorm"
)

type (
	userRepository struct {
		db *gorm.DB
	}
)

func NewLoginRepository(db *gorm.DB) usecase.LoginRepository {
	return &userRepository{
		db: db,
	}
}

func NewAuthRepository(db *gorm.DB) usecase.AuthRepository {
	return &userRepository{
		db: db,
	}
}

func NewUserRepository(db *gorm.DB) usecase.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) ValidateUser(id string, pass string) (ok bool, err error) {
	user := model.User{}
	if err := r.db.Where("id = ? and password = ?", id, pass).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepository) ExistUser(id string) (exist bool, err error) {
	user := model.User{}
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (r *userRepository) IsAdmin(id string) (isAdmin bool, err error) {
	user := model.User{}
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return false, err
	}

	return user.Admin, nil
}

func (r *userRepository) CreateUser(info *usecase.UserInfo) error {
	user := model.User{
		ManualIDModel: model.ManualIDModel{
			ID: info.Id,
		},
		Password: info.Password,
		Admin:    false,
	}
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUser(id string) (out *usecase.UserInfo, notFound bool, err error) {
	user := model.User{
		ManualIDModel: model.ManualIDModel{ID: id},
	}

	err = r.db.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, true, err
		} else {
			return nil, false, err
		}
	}

	out = &usecase.UserInfo{Id: user.ID, Password: user.Password}
	return out, false, nil
}

func (r *userRepository) DeleteUser(id string) (notFound bool, err error) {
	user := model.User{
		ManualIDModel: model.ManualIDModel{ID: id},
	}

	err = r.db.Delete(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, err
		} else {
			return false, err
		}
	}

	return false, nil
}

func (r *userRepository) UpdateUser(info *usecase.UserInfo) (notFound bool, err error) {
	user := model.User{
		ManualIDModel: model.ManualIDModel{ID: info.Id},
	}

	err = r.db.Model(&user).Updates(model.User{ManualIDModel: model.ManualIDModel{ID: info.Id}, Password: info.Password}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true, err
		} else {
			return false, err
		}
	}

	return false, nil
}
