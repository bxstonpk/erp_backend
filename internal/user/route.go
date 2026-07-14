package user

import (
	"net/http"

	"erp/backend/internal/user/handler"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router, h *handler.Handler) {
	users := router.PathPrefix("/users").Subrouter()
	users.HandleFunc("", h.ListUsers).Methods(http.MethodGet)
	users.HandleFunc("", h.CreateUser).Methods(http.MethodPost)
	users.HandleFunc("/username/{username}", h.GetUserByUsername).Methods(http.MethodGet)
	users.HandleFunc("/{userUUID}", h.GetUser).Methods(http.MethodGet)
	users.HandleFunc("/{userUUID}", h.UpdateUser).Methods(http.MethodPut)
	users.HandleFunc("/{userUUID}", h.DeleteUser).Methods(http.MethodDelete)
	users.HandleFunc("/{userUUID}/roles", h.GetUserRoles).Methods(http.MethodGet)
	users.HandleFunc("/{userUUID}/roles", h.AssignUserRole).Methods(http.MethodPost)
	users.HandleFunc("/{userUUID}/permissions", h.GetUserPermissionAssignments).Methods(http.MethodGet)
	users.HandleFunc("/{userUUID}/permissions", h.AssignUserPermission).Methods(http.MethodPost)
	users.HandleFunc("/{userUUID}/company-access", h.GetUserCompanyAccesses).Methods(http.MethodGet)
	users.HandleFunc("/{userUUID}/company-access", h.AssignUserCompanyAccess).Methods(http.MethodPost)

	userRoles := router.PathPrefix("/user-roles").Subrouter()
	userRoles.HandleFunc("/{userRoleUUID}", h.UpdateUserRole).Methods(http.MethodPut)
	userRoles.HandleFunc("/{userRoleUUID}", h.DeleteUserRole).Methods(http.MethodDelete)

	userPermissions := router.PathPrefix("/user-permissions").Subrouter()
	userPermissions.HandleFunc("/{userPermissionUUID}", h.UpdateUserPermission).Methods(http.MethodPut)
	userPermissions.HandleFunc("/{userPermissionUUID}", h.DeleteUserPermission).Methods(http.MethodDelete)

	userCompanyAccess := router.PathPrefix("/user-company-access").Subrouter()
	userCompanyAccess.HandleFunc("/{userCompanyUUID}", h.UpdateUserCompanyAccess).Methods(http.MethodPut)
	userCompanyAccess.HandleFunc("/{userCompanyUUID}", h.DeleteUserCompanyAccess).Methods(http.MethodDelete)

	roles := router.PathPrefix("/roles").Subrouter()
	roles.HandleFunc("", h.ListRoles).Methods(http.MethodGet)
	roles.HandleFunc("", h.CreateRole).Methods(http.MethodPost)
	roles.HandleFunc("/{roleUUID}", h.GetRole).Methods(http.MethodGet)
	roles.HandleFunc("/{roleUUID}", h.UpdateRole).Methods(http.MethodPut)
	roles.HandleFunc("/{roleUUID}", h.DeleteRole).Methods(http.MethodDelete)
	roles.HandleFunc("/{roleUUID}/permissions", h.GetRolePermissionAssignments).Methods(http.MethodGet)
	roles.HandleFunc("/{roleUUID}/permissions", h.AssignRolePermission).Methods(http.MethodPost)

	rolePermissions := router.PathPrefix("/role-permissions").Subrouter()
	rolePermissions.HandleFunc("/{rolePermissionUUID}", h.UpdateRolePermission).Methods(http.MethodPut)
	rolePermissions.HandleFunc("/{rolePermissionUUID}", h.DeleteRolePermission).Methods(http.MethodDelete)

	permissions := router.PathPrefix("/permissions").Subrouter()
	permissions.HandleFunc("", h.ListPermissions).Methods(http.MethodGet)
	permissions.HandleFunc("", h.CreatePermission).Methods(http.MethodPost)
	permissions.HandleFunc("/{permissionUUID}", h.GetPermission).Methods(http.MethodGet)
	permissions.HandleFunc("/{permissionUUID}", h.UpdatePermission).Methods(http.MethodPut)
	permissions.HandleFunc("/{permissionUUID}", h.DeletePermission).Methods(http.MethodDelete)
}
