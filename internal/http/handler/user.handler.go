package handler

import (
	"net/http"

	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/internal/http/binder"
	"github.com/DavidAfdal/workfinder/internal/service"
	"github.com/DavidAfdal/workfinder/pkg/response"
	"github.com/DavidAfdal/workfinder/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)


type UserHandler interface {
	FindAllUser(ctx echo.Context) error
	CreateUser(ctx echo.Context) error
	Login(ctx echo.Context) error
	UpdateUser(ctx echo.Context) error
	DeleteUser(ctx echo.Context) error
	FindByUserID(ctx echo.Context) error
	ProfileUser(ctx echo.Context) error
	Logout(ctx echo.Context) error
}

type userHandler struct {
   userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{userService: userService}
}

// @Summary Get All Users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [get]
func (h *userHandler) FindAllUser(ctx echo.Context) error {

	data, err := h.userService.FindAllUser()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}
	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get all users", data))
}

// @Summary Login
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body binder.LoginRequest true "Login Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /login [post]
func (h *userHandler) Login(ctx echo.Context) error {
	var input binder.LoginRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	accessToken, err := h.userService.Login(input.Email, input.Password)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success login", map[string]interface{}{
		"access_token": accessToken,
	}))
}

// @Summary Create User
// @Description Create a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body binder.CreateUserRequest true "Create User Request" Example({"name":"John Doe","email":"john.doe@example.com","password":"password123","address":"123 Main St","phone_number":"08123451672","gender":"male","role":"applicant"})
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /register [post]
func (h *userHandler) CreateUser(ctx echo.Context) error {
	var input binder.CreateUserRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	newUser := entity.NewUser(input.Name, input.Email, input.Password, input.Address, input.PhoneNumber, input.Gender, input.Role)
	user, err := h.userService.CreateUser(newUser)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success create user", user))
}

// @Summary Update User
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body binder.UpdateUserRequest true "Update User Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [put]
func (h *userHandler) UpdateUser(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)

	var input binder.UpdateUserRequest
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	updateUser := entity.UpdateUser(claims.ID, input.Name, input.Email, input.Password, input.Address, input.PhoneNumber, input.Gender)

	updatedUser, err := h.userService.UpdateUser(updateUser)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success update user", updatedUser))
}

// @Summary Delete User
// @Description Delete an existing user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users [delete]
func (h *userHandler) DeleteUser(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)


	isDeleted, err := h.userService.DeleteUser(claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success delete user", isDeleted))
}

// @Summary Get User Profile
// @Description Get the profile of the authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/profile [get]
func (h *userHandler) ProfileUser(ctx echo.Context) error {
	dataUser, _ := ctx.Get("user").(*jwt.Token)
	claims := dataUser.Claims.(*token.JwtCustomClaims)


	user, err := h.userService.FindById(claims.ID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get user", user))
}

// @Summary Get User By ID
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param request body binder.UserFindByIDRequest true "Find User by ID Request"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /users/{id} [get]
func (h *userHandler) FindByUserID(ctx echo.Context) error {
	var input binder.UserFindByIDRequest

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	id := uuid.MustParse(input.ID)

	user, err := h.userService.FindById(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return ctx.JSON(http.StatusOK, response.SuccessResponse(http.StatusOK, "success get user", user))
}

func (h *userHandler) Logout(ctx echo.Context) error {
	return nil
}

