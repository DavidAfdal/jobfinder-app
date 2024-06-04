package binder

// TODO : Create Type Input Request For User Handler

type LoginRequest struct {
	Email string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Address string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	Gender string `json:"gender"`
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
