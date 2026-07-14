package usecase

import (
	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	CreateUser(req dto.CreateUserRequest) (*entity.User, error)
	GetUser(userUUID string) (*entity.User, error)
	GetUserByUsername(username string) (*entity.User, error)
	ListUsers() ([]*entity.User, error)
	UpdateUser(userUUID string, req dto.UpdateUserRequest) (*entity.User, error)
	DeleteUser(userUUID string) error
}

func (u *userUsecase) CreateUser(req dto.CreateUserRequest) (*entity.User, error) {
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		ID:       uuid.NewString(),
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		FullName: req.FullName,
		Phone:    req.Phone,
		Status:   req.Status,
	}

	if err := u.repo.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// hashPassword bcrypt-hashes a plaintext password. entity.User.Password (and
// the persisted "hashed_password" column) are only ever meant to hold the
// hash, never the plaintext the caller submitted.
func hashPassword(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (u *userUsecase) GetUser(userUUID string) (*entity.User, error) {
	return u.repo.GetUserByUUID(userUUID)
}

func (u *userUsecase) GetUserByUsername(username string) (*entity.User, error) {
	return u.repo.GetUserByUsername(username)
}

func (u *userUsecase) ListUsers() ([]*entity.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecase) UpdateUser(userUUID string, req dto.UpdateUserRequest) (*entity.User, error) {
	user, err := u.repo.GetUserByUUID(userUUID)
	if err != nil {
		return nil, err
	}

	user.Username = req.Username
	user.Email = req.Email
	user.FullName = req.FullName
	user.Phone = req.Phone
	user.Status = req.Status

	if req.Password != "" {
		hashedPassword, err := hashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		user.Password = hashedPassword
	}

	if err := u.repo.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) DeleteUser(userUUID string) error {
	return u.repo.DeleteUser(userUUID)
}
