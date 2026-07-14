package usecase

import (
	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"

	"github.com/google/uuid"
)

type RoleUsecase interface {
	CreateRole(req dto.CreateRoleRequest) (*entity.Role, error)
	GetRole(roleUUID string) (*entity.Role, error)
	ListRoles() ([]*entity.Role, error)
	UpdateRole(roleUUID string, req dto.UpdateRoleRequest) (*entity.Role, error)
	DeleteRole(roleUUID string) error
}

func (u *userUsecase) CreateRole(req dto.CreateRoleRequest) (*entity.Role, error) {
	role := &entity.Role{
		ID:          uuid.NewString(),
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := u.repo.CreateRole(role); err != nil {
		return nil, err
	}

	return role, nil
}

func (u *userUsecase) GetRole(roleUUID string) (*entity.Role, error) {
	return u.repo.GetRoleByUUID(roleUUID)
}

func (u *userUsecase) ListRoles() ([]*entity.Role, error) {
	return u.repo.GetAllRoles()
}

func (u *userUsecase) UpdateRole(roleUUID string, req dto.UpdateRoleRequest) (*entity.Role, error) {
	role := &entity.Role{
		ID:          roleUUID,
		CompanyID:   req.CompanyID,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := u.repo.UpdateRole(role); err != nil {
		return nil, err
	}

	return role, nil
}

func (u *userUsecase) DeleteRole(roleUUID string) error {
	return u.repo.DeleteRole(roleUUID)
}
