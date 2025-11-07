package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (c *ListController) UpdateList(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	list := new(models.List)

	if err := ctx.BodyParser(list); err != nil {
		return utils.BadRequest(ctx, "Gagal parsing data", err.Error())
	}

	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "ID tidak valid", err.Error())
	}

	existingList, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List tidak ditemukan", err.Error())
	}

	list.InternalID = existingList.InternalID
	list.PublicID = existingList.PublicID

	if err := c.service.Update(list); err != nil {
		return utils.BadRequest(ctx, "Gagal update list", err.Error())
	}

	updatedList, err := c.service.GetByPublicID(publicID)
	if err != nil {
		return utils.NotFound(ctx, "List tidak ditemukan", err.Error())
	}

	return utils.Success(ctx, "Berhasil memperbaharui list", updatedList)
}

func (c *ListController) GetListOnBoard(ctx *fiber.Ctx) error {
	boardPublicID := ctx.Params("board_id")
	if _, err := uuid.Parse(boardPublicID); err != nil {
		return utils.BadRequest(ctx, "ID tidak valid", err.Error())
	}

	list, err := c.service.GetByBoardID(boardPublicID)
	if err != nil {
		return utils.NotFound(ctx, "List tidak ditemukan", err.Error())
	}

	return utils.Success(ctx, "Data berhasil diambil", list)
}

func (c *ListController) DeleteList(ctx *fiber.Ctx) error {
	publicId := ctx.Params("id")
	if _, err := uuid.Parse(publicId); err != nil {
		return utils.BadRequest(ctx, "ID tidak valid", err.Error())
	}

	list, err := c.service.GetByPublicID(publicId)
	if err != nil {
		return utils.NotFound(ctx, "List tidak ditemukan", err.Error())
	}

	if err := c.service.Delete(uint(list.InternalID)); err != nil {
		return utils.InternalServerError(ctx, "Gagal menghapus list", err.Error())
	}

	return utils.Success(ctx, "Berhasil menghapus list", publicId)
}
