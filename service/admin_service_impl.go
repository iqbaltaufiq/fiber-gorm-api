package service

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iqbaltaufiq/go-fiber-restapi/helper"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/domain"
	"github.com/iqbaltaufiq/go-fiber-restapi/model/web"
	"github.com/iqbaltaufiq/go-fiber-restapi/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	DB              *gorm.DB
	Validate        *validator.Validate
}

func NewAdminService(AdminRepository repository.AdminRepository, DB *gorm.DB, validate *validator.Validate) AdminService {
	return &AdminServiceImpl{
		AdminRepository: AdminRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (s *AdminServiceImpl) Create(ctx context.Context, request web.CreateAdminRequest) fiber.Map {
	if errVal := s.Validate.Struct(request); errVal != nil {
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

	// hash password
	hash, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	// mapping
	admin := domain.Admin{
		Username: request.Username,
		Password: string(hash),
		Role:     "admin",
	}

	admin, errSave := s.AdminRepository.Save(ctx, tx, admin)
	if errSave != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errSave.Error(),
		}
	}

	return fiber.Map{
		"error":  false,
		"result": admin,
	}
}

func (s *AdminServiceImpl) Login(ctx context.Context, credential web.LoginAdminRequest) fiber.Map {
	if errVal := s.Validate.Struct(credential); errVal != nil {
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

	admin, errFind := s.AdminRepository.FindByUsername(ctx, tx, credential.Username)
	if errFind != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errFind.Error(),
		}
	}

	errCheck := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(credential.Password))
	if errCheck != nil {
		return fiber.Map{
			"error":   true,
			"err_msg": errCheck.Error(),
		}
	}

	return fiber.Map{
		"error": false,
	}
}
