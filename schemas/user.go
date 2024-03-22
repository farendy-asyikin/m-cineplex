package schemas

type CreateUserRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Role            string `json:"role" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type UpdateUserRequest struct {
	Name  *string `form:"name"`
	Email *string `form:"email"`
}

type DetailUserResponse struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}
