func (p postgresUserRepository) GetUserByUUID(userUUID string) (*entity.User, error) {
	user := &entity.User{}
	err := p.db.Preload("UserRoles").Preload("UserPermissions").Preload("RolePermissions").Preload("UserCompanyAccess").Where("user_uuid = ?", userUUID).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p postgresUserRepository) GetUserByUsername(username string) (*entity.User, error) {
	user := &entity.User{}
	err := p.db.Preload("UserRoles").Preload("UserPermissions").Preload("RolePermissions").Preload("UserCompanyAccess").Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (p postgresUserRepository) GetAllUsers() ([]*entity.User, error) {
	users := []*entity.User{}
	err := p.db.Preload("UserRoles").Preload("UserPermissions").Preload("RolePermissions").Preload("UserCompanyAccess").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p postgresUserRepository) GetAllRoles() ([]*entity.Role, error) {
	roles := []*entity.Role{}
	err := p.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (p postgresUserRepository) GetAllUserRoles() ([]*entity.Role, error) {
	userRoles := []*entity.Role{}
	err := p.db.Find(&userRoles).Error
	if err != nil {
		return nil, err
	}

	return userRoles, nil
}

func (p postgresUserRepository) GetAllUserPermissions() ([]*entity.UserPermission, error) {
	userPermissions := []*entity.UserPermission{}
	err := p.db.Find(&userPermissions).Error
	if err != nil {
		return nil, err
	}

	return userPermissions, nil
}

func (p postgresUserRepository) GetAllRolePermissions() ([]*entity.RolePermission, error) {
	rolePermissions := []*entity.RolePermission{}
	err := p.db.Find(&rolePermissions).Error
	if err != nil {
		return nil, err
	}

	return rolePermissions, nil
}

func (p postgresUserRepository) GetAllUserCompanyAccesses() ([]*entity.CompanyAccess, error) {
	userCompanyAccesses := []*entity.CompanyAccess{}
	err := p.db.Find(&userCompanyAccesses).Error
	if err != nil {
		return nil, err
	}

	return userCompanyAccesses, nil
}