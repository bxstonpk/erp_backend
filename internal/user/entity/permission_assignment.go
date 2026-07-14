package entity

type PermissionSource string

const (
	PermissionSourceUser PermissionSource = "user"
	PermissionSourceRole PermissionSource = "role"
)

// PermissionAssignment represents a Permission attached to a specific User or
// Role. Source and OwnerID say which one, and AssignmentID is the join row's
// own UUID, needed to update or delete the assignment itself.
type PermissionAssignment struct {
	AssignmentID string
	Source       PermissionSource
	OwnerID      string
	Permission   Permission
}
