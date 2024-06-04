package repository

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/DavidAfdal/workfinder/internal/entity"
	"github.com/DavidAfdal/workfinder/pkg/cache"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


type UserRepository interface {
	FindAllUser() ([]entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindById(id uuid.UUID) (*entity.User, error)
	CreateUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(user *entity.User) (bool, error)
}

type userRepository struct {
	db *gorm.DB
	cahce cache.Cacheable
}

func NewUserRepository(db *gorm.DB, cahce cache.Cacheable) UserRepository {
	return &userRepository{db, cahce}
}


func (r *userRepository) FindAllUser() ([]entity.User, error) {
	users := make([]entity.User, 0)


	key:= "GetAllUsers"

	data := r.cahce.Get(key)

	fmt.Println(data)

	if data == "" {
		if err := r.db.Find(&users).Error; err != nil {
			return users, err
		}

		marshalJob, _:= json.Marshal(users)
		err := r.cahce.Set(key, marshalJob, 2 * time.Minute)

		if err != nil {
			return users, err
		}
	} else {
		err := json.Unmarshal([]byte(data), &users)
		if err != nil {
			return users, err
		}
	}



	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	user := new(entity.User)

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindById(id uuid.UUID) (*entity.User, error) {
	user := new(entity.User)

	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user *entity.User) (*entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}


func (r *userRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	fields := make(map[string]interface{})

	if user.Password != "" {
		fields["password"] = user.Password
	}

	if user.Address != "" {
		fields["address" ] = user.Address
	}

	if user.PhoneNumber != "" {
		fields["phone_number" ] = user.PhoneNumber
	}

	if err := r.db.Model(&user).Updates(fields).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) DeleteUser(user *entity.User) (bool, error){
	if err:= r.db.Delete(&user).Error; err != nil {
		return false, nil
	}

	return true, nil
}

