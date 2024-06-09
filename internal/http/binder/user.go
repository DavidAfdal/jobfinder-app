package binder

// TODO : Create Type Input Request For User Handler

type LoginRequest struct {
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserRequest struct {
	Name        string `json:"name" example:"John Doe"`
    Email       string `json:"email" example:"john.doe@example.com"`
    Password    string `json:"password" example:"password123"`
    Address     string `json:"address" example:"123 Main St"`
    PhoneNumber string `json:"phone_number" example:"123-456-7890"`
    Gender      string `json:"gender" example:"male"`
    Role        string `json:"role" example:"applicant"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Gender string `json:"gender"`
}

type UserFindByIDRequest struct {
	ID string `param:"id" validate:"required"`
}
