package service

import (
	"time"

	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/internal/repository"
	"github.com/DavidAfdal/workfinder/pkg/token"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// TODO: Create User Service Struct and Interface

// TODO: Create User Service Implementation

type UserService interface {
	Login(email string, password string) (string, error)
	CreateUser(user *entity.User) (*entity.User, error)
	FindById(id uuid.UUID) (*entity.User, error)
	FindAllUser() ([]entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id uuid.UUID) (bool, error)
}

type userService struct {
  userRepo repository.UserRepository
  tokenUseCase token.TokenUseCase
}

func NewUserService(userRepo repository.UserRepository, tokenUseCase token.TokenUseCase) UserService {
	return &userService{userRepo, tokenUseCase}
}


func (s *userService) Login(email string, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)

	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", err
	}

	expiredTime := time.Now().Local().Add(time.Minute * 5)

	claims := token.JwtCustomClaims{
		ID: user.ID,
		Email: user.Email,
		Role: user.Role,
		Address: user.Address,
		PhoneNumber: user.PhoneNumber,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	}

	token, err := s.tokenUseCase.GenerateAccessToken(claims)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) CreateUser(user *entity.User) (*entity.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hashedPassword)
	return s.userRepo.CreateUser(user)
}

func (s *userService) UpdateUser(user *entity.User) (*entity.User, error) {
	if user.Password != ""{
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return user, err
		}
		user.Password = string(hashedPassword)
	}

	return s.userRepo.UpdateUser(user)
}

func (s *userService) FindAllUser() ([]entity.User, error) {
	return s.userRepo.FindAllUser()
}

func (s *userService) FindById(id uuid.UUID) (*entity.User, error) {
	return s.userRepo.FindById(id)
}

func (s *userService) DeleteUser(id uuid.UUID) (bool, error) {
	user, err := s.userRepo.FindById(id)

	if err != nil {
		return false, err
	}

	return s.userRepo.DeleteUser(user)
}



