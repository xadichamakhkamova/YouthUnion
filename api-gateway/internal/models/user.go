package models

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`             // HTTP status kodi
	Message string `json:"message" example:"invalid data"` // Xatolik tavsifi
	Details any    `json:"details,omitempty"`              // Qo‘shimcha ma’lumot (ixtiyoriy)
}

type User struct {
	ID          string `json:"id" db:"id"`
	Identifier  int32  `json:"identifier" db:"identifier"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Faculty     string `json:"faculty" db:"faculty"`
	Course      int32  `json:"course" db:"course"`
	BirthDate   string `json:"birth_date" db:"birth_date"`
	Gender      string `json:"gender" db:"gender"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

type CreateUserRequest struct {
	Identifier   int32  `json:"identifier" db:"identifier"`
	FirstName    string `json:"first_name" db:"first_name"`
	LastName     string `json:"last_name" db:"last_name"`
	PhoneNumber  string `json:"phone_number" db:"phone_number"`
	PasswordHash string `json:"password" db:"password_hash"`
	Faculty      string `json:"faculty" db:"faculty"`
	Course       int32  `json:"course" db:"course"`
	BirthDate    string `json:"birth_date" db:"birth_date"`
	Gender       string `json:"gender" db:"gender"`
}

type GetUserByIdentifierRequest struct {
	Identifier int32  `json:"identifier" db:"identifier"`
	Password   string `json:"password" db:"password"`
}

type GetUserByIdentifierResponse struct {
	Status int32 `json:"status"`
}

type UpdateUserRequest struct {
	ID          string `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	PhoneNumber string `json:"phone_number" db:"phone_number"`
	Faculty     string `json:"faculty" db:"faculty"`
	Course      int32  `json:"course" db:"course"`
	BirthDate   string `json:"birth_date" db:"birth_date"`
	Gender      string `json:"gender" db:"gender"`
}

type ChangePasswordRequest struct {
	UserID      string `json:"user_id" db:"user_id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponse struct {
	Status    int32  `json:"status"`
	UpdatedAt string `json:"updated_at"`
}

type GetUserByIdRequest struct {
	ID string `json:"id"`
}

type ListUsersRequest struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

type UserList struct {
	Users      []User `json:"users"`
	TotalCount int32  `json:"total_count"`
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}

type DeleteUserResponse struct {
	Status        int32  `json:"status"`
	DeletedUserID string `json:"deleted_user_id"`
	DeletedAt     string `json:"deleted_at"`
}

type UserSession struct {
	ID           string `json:"id" db:"id"`
	UserID       string `json:"user_id" db:"user_id"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	ExpiresAt    string `json:"expires_at" db:"expires_at"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}

type CreateSessionRequest struct {
	UserID       string `json:"user_id" db:"user_id"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	ExpiresAt    string `json:"expires_at" db:"expires_at"`
}

type GetSessionByTokenRequest struct {
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
}

type RoleType struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}

type CreateRoleRequest struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type UpdateRoleRequest struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type GetRoleByIdRequest struct {
	ID string `json:"id"`
}

type DeleteRoleRequest struct {
	ID string `json:"id"`
}

type DeleteRoleResponse struct {
	Status        int32  `json:"status"`
	DeletedRoleID string `json:"deleted_role_id"`
	DeletedAt     string `json:"deleted_at"`
}

type ListRolesRequest struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}

type RoleTypeList struct {
	Roles      []RoleType `json:"roles"`
	TotalCount int32      `json:"total_count"`
}

type UserRole struct {
	ID         string `json:"id" db:"id"`
	UserID     string `json:"user_id" db:"user_id"`
	RoleID     string `json:"role_id" db:"role_id"`
	AssignedAt string `json:"assigned_at" db:"assigned_at"`
}

type AssignRoleRequest struct {
	UserID string `json:"user_id" db:"user_id"`
	RoleID string `json:"role_id" db:"role_id"`
}

type RemoveRoleRequest struct {
	ID string `json:"id" db:"id"`
}

type RemoveRoleResponse struct {
	Status        int32  `json:"status"`
	RemovedRoleID string `json:"removed_role_id"`
	RemovedAt     string `json:"removed_at"`
}

type ListUserRolesRequest struct {
	UserID string `json:"user_id"`
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
}

type UserRoleList struct {
	UserRoles  []UserRole `json:"user_roles"`
	TotalCount int32      `json:"total_count"`
}
