package usecase

import (
	"erp/backend/internal/user/dto"
	"erp/backend/internal/user/entity"
)

type CompanyAccessUsecase interface {
	AssignUserCompanyAccess(userUUID string, req dto.CompanyAccessRequest) error
	UpdateUserCompanyAccess(userCompanyUUID string, req dto.UpdateCompanyAccessRequest) error
	DeleteUserCompanyAccess(userCompanyAccessUUID string) error
	GetUserCompanyAccesses(userUUID string) ([]*entity.CompanyAccess, error)
}

func (u *userUsecase) AssignUserCompanyAccess(userUUID string, req dto.CompanyAccessRequest) error {
	access := &entity.CompanyAccess{
		CompanyID: req.CompanyID,
		BranchID:  req.BranchID,
		IsDefault: req.IsDefault,
		ReadOnly:  req.ReadOnly,
		Active:    req.Active,
	}

	return u.repo.AssignUserCompanyAccess(userUUID, access)
}

func (u *userUsecase) UpdateUserCompanyAccess(userCompanyUUID string, req dto.UpdateCompanyAccessRequest) error {
	access := &entity.CompanyAccess{
		CompanyID: req.CompanyID,
		BranchID:  req.BranchID,
		IsDefault: req.IsDefault,
		ReadOnly:  req.ReadOnly,
		Active:    req.Active,
	}

	return u.repo.UpdateUserCompanyAccess(userCompanyUUID, req.UserUUID, access)
}

func (u *userUsecase) DeleteUserCompanyAccess(userCompanyAccessUUID string) error {
	return u.repo.DeleteUserCompanyAccess(userCompanyAccessUUID)
}

func (u *userUsecase) GetUserCompanyAccesses(userUUID string) ([]*entity.CompanyAccess, error) {
	return u.repo.GetUserCompanyAccesses(userUUID)
}
