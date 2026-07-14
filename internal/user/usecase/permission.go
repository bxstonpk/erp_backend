package usecase

import (
	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"

	"github.com/google/uuid"
)

type PermissionUsecase interface {
	CreatePermission(req dto.CreatePermissionRequest) (*entity.Permission, error)
	GetPermission(permissionUUID string) (*entity.Permission, error)
	ListPermissions() ([]*entity.Permission, error)
	UpdatePermission(permissionUUID string, req dto.UpdatePermissionRequest) (*entity.Permission, error)
	DeletePermission(permissionUUID string) error
}

func (u *userUsecase) CreatePermission(req dto.CreatePermissionRequest) (*entity.Permission, error) {
	permission := &entity.Permission{
		ID:          uuid.NewString(),
		CompanyID:   req.CompanyID,
		Code:        req.Code,
		Module:      req.Module,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := u.repo.CreatePermission(permission); err != nil {
		return nil, err
	}

	return permission, nil
}

func (u *userUsecase) GetPermission(permissionUUID string) (*entity.Permission, error) {
	return u.repo.GetPermissionByUUID(permissionUUID)
}

func (u *userUsecase) ListPermissions() ([]*entity.Permission, error) {
	return u.repo.GetAllPermissions()
}

func (u *userUsecase) UpdatePermission(permissionUUID string, req dto.UpdatePermissionRequest) (*entity.Permission, error) {
	permission := &entity.Permission{
		ID:          permissionUUID,
		CompanyID:   req.CompanyID,
		Code:        req.Code,
		Module:      req.Module,
		Name:        req.Name,
		Description: req.Description,
	}

	if err := u.repo.UpdatePermission(permission); err != nil {
		return nil, err
	}

	return permission, nil
}

func (u *userUsecase) DeletePermission(permissionUUID string) error {
	return u.repo.DeletePermission(permissionUUID)
}
