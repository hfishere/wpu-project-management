package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hfishere/wpu-project-management/models"
	"github.com/hfishere/wpu-project-management/services"
	"github.com/hfishere/wpu-project-management/utils"
)

type ListController struct {
	service services.ListService
}

func NewListController(s services.ListService) *ListController {
	return &ListController{service: s}
}

func (c *ListController) CreateList(ctx *fiber.Ctx) error {
	list := new(models.List)
	if err := ctx.BodyParser(list); err != nil {
		return utils.BadRequest(ctx, "Gagal membaca request", err.Error())
	}

	if err := c.service.Create(list); err != nil {
		return utils.BadRequest(ctx, "Gagal membuat list", err.Error())
	}

	return utils.Success(ctx, "Berhasil membuat list", list)
}
