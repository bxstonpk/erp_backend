package usecase

import "erp/backend/internal/user/repository"

type Usecase interface {
	UserUsecase
	RoleUsecase
	PermissionUsecase
	AssignmentUsecase
	CompanyAccessUsecase
}

type userUsecase struct {
	repo repository.Repository
}

func New(repo repository.Repository) Usecase {
	return &userUsecase{repo: repo}
}
