package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/helper"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"github.com/iqbaltaufiq/go-fiber-restapi/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

// constructor
func NewUserService(UserRepo repository.UserRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: UserRepo,
		DB:             DB,
		Validate:       validate,
	}
}

func (s *UserServiceImpl) Register(ctx context.Context, request web.CreateUserRequest) fiber.Map {
	errVal := s.Validate.Struct(request)
	if errVal != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errVal.Error(),
		}
	}

	tx := s.DB.Begin()
	if tx.Error != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": tx.Error.Error(),
		}
	}
	defer helper.CommitOrRollback(tx)

	// hash user password
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	userReq := domain.User{
		Username: request.Username,
		Password: string(hash),
		Role:     "user",
	}

	user, errCreate := s.UserRepository.CreateUser(ctx, tx, userReq)
	if errCreate != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errCreate.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": user,
	}
}

func (s *UserServiceImpl) Login(ctx context.Context, request web.LoginUserRequest) fiber.Map {
	errVal := s.Validate.Struct(request)
	if errVal != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errVal.Error(),
		}
	}

	tx := s.DB.Begin()
	if tx.Error != nil {
		panic(tx.Error)
	}
	defer helper.CommitOrRollback(tx)

	user, errFind := s.UserRepository.FindUserByUsername(ctx, tx, request.Username)
	if errFind != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errFind.Error(),
		}
	}

	errCheck := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if errCheck != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errCheck.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": true,
	}
}

func (s *UserServiceImpl) GenerateApiKey(ctx context.Context, username string) fiber.Map {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": tx.Error.Error(),
		}
	}
	defer helper.CommitOrRollback(tx)

	user, errFind := s.UserRepository.FindUserByUsername(ctx, tx, username)
	if errFind != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errFind.Error(),
		}
	}

	// the api key is an MD5 hash of current nanosecond
	// concatenated with user id
	timeNow := strconv.Itoa(time.Now().Nanosecond())
	userIdForKey := strconv.Itoa(int(user.Id))
	hash := fmt.Sprintf("%x", md5.Sum([]byte(timeNow+userIdForKey)))

	key := domain.ApiKey{
		Username: user.Username,
		ApiKey:   hash,
		Expires:  time.Now().Add(30 * 24 * time.Hour),
	}

	returnedKey, errKey := s.UserRepository.CreateApiKey(ctx, tx, key)
	if errKey != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errKey.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": returnedKey,
	}
}
