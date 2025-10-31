package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hfishere/wpu-project-management/models"
	"github.com/hfishere/wpu-project-management/services"
	"github.com/hfishere/wpu-project-management/utils"
	"github.com/jinzhu/copier"
)

type UserController struct {
	service services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{service: s}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return utils.BadRequest(ctx, "Gagal Parsing Data", err.Error())
	}

	if err := c.service.Register(user); err != nil {
		return utils.BadRequest(ctx, "Registrasi Gagal", err.Error())
	}

	var userResponse models.UserResponse
	_ = copier.Copy(&userResponse, &user)

	return utils.Success(ctx, "Register Success", userResponse)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.BodyParser(&body); err != nil {
		return utils.BadRequest(ctx, "Invalid Request", err.Error())
	}

	user, err := c.service.Login(body.Email, body.Password)
	if err != nil {
		return utils.Unauthorized(ctx, "Login Failed", err.Error())
	}

	token, _ := utils.GenerateToken(user.InternalID, user.Role, user.Email, user.PublicID)
	refreshToken, _ := utils.GenerateRefreshToken(user.InternalID)

	var userResponse models.UserResponse
	_ = copier.Copy(&userResponse, &user)
	return utils.Success(ctx, "Login Succesful", fiber.Map{
		"access_token":  token,
		"refresh_token": refreshToken,
		"user":          userResponse,
	})
}
