package usecase

import "doc-api/api/entity"

type (
	UserInfo struct {
		Id       string
		Password string
	}

	UserPostOutputPort interface {
		Success() error
		Failure(err error) error
	}

	UserGetOutputPort interface {
		Success(info *UserInfo) error
		NotFound(err error) error
		Failure(err error) error
	}

	UserDeleteOutputPort interface {
		Success() error
		NotFound(err error) error
		Failure(err error) error
	}

	UserUpdateOutputPort interface {
		Success() error
		NotFound(err error) error
		Failure(err error) error
	}

	UserRepository interface {
		CreateUser(info *UserInfo) error
		UpdateUser(info *UserInfo) (notFound bool, err error)
		DeleteUser(id string) (notFound bool, err error)
		GetUser(id string) (out *UserInfo, notFound bool, err error)
	}

	UserService interface {
		SignupUser(info *UserInfo, out UserPostOutputPort) error
		GetUser(userId string, out UserGetOutputPort) error
		DeleteUser(userId string, out UserDeleteOutputPort) error
		UpdateUser(info *UserInfo, out UserUpdateOutputPort) error
	}

	userService struct {
		r UserRepository
		h entity.Hashing
	}
)

func NewUserService(rep UserRepository, hashing entity.Hashing) UserService {
	return &userService{r: rep, h: hashing}
}

func (s *userService) SignupUser(info *UserInfo, out UserPostOutputPort) error {

	hashed := UserInfo{
		Id:       info.Id,
		Password: s.h.Hash(info.Password),
	}

	if err := s.r.CreateUser(&hashed); err != nil {
		return out.Failure(err)
	}

	return out.Success()
}

func (s *userService) GetUser(userId string, out UserGetOutputPort) error {
	o, notFound, err := s.r.GetUser(userId)
	if notFound {
		return out.NotFound(err)
	}
	if err != nil {
		return out.Failure(err)
	}
	return out.Success(o)
}

func (s *userService) DeleteUser(userId string, out UserDeleteOutputPort) error {
	notFound, err := s.r.DeleteUser(userId)
	if notFound {
		return out.NotFound(err)
	}
	if err != nil {
		return out.Failure(err)
	}
	return out.Success()
}

func (s *userService) UpdateUser(info *UserInfo, out UserUpdateOutputPort) error {
	hashed := UserInfo{
		Id:       info.Id,
		Password: s.h.Hash(info.Password),
	}

	notFound, err := s.r.UpdateUser(&hashed)
	if notFound {
		return out.NotFound(err)
	}
	if err != nil {
		return out.Failure(err)
	}
	return out.Success()
}
