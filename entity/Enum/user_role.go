package enum

type UserRole uint8

const (
	RoleAdmin     UserRole = 0
	RoleUser      UserRole = 1
	RoleModerator UserRole = 2
)

func (r UserRole) GetLabel() string {
	switch r {
	case RoleAdmin:
		return "admin"
	case RoleUser:
		return "user"
	case RoleModerator:
		return "moderator"
	default:
		return "unknown"
	}
}

func UserRoleFromString(s string) UserRole {
	switch s {
	case "admin":
		return RoleAdmin
	case "moderator":
		return RoleModerator
	case "user":
		return RoleUser
	default:
		return RoleUser
	}
}
