package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/hfishere/wpu-project-management/models"
	"github.com/hfishere/wpu-project-management/services"
	"github.com/hfishere/wpu-project-management/utils"
)

type CardController struct {
	service services.CardService
}

func NewCardController(s services.CardService) *CardController {
	return &CardController{service: s}
}

func (c *CardController) CreateCard(ctx *fiber.Ctx) error {
	type CreateCardRequest struct {
		ListPublicID string    `json:"list_id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		DueDate      time.Time `json:"due_date"`
		Position     int       `json:"position"`
	}

	var req CreateCardRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "Gagal Mengambil Data", err.Error())
	}

	card := &models.Card{
		Title:       req.Title,
		Description: req.Description,
		DueDate:     &req.DueDate,
		Position:    int64(req.Position),
	}

	if err := c.service.Create(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "Gagal Membuat Card", err.Error())
	}

	return utils.Success(ctx, "Berhasil Membuat Card", card)
}

func (c *CardController) GetCardOnList(ctx *fiber.Ctx) error {
	listPublicID := ctx.Params("list_id")
	if _, err := uuid.Parse(listPublicID); err != nil {
		return utils.BadRequest(ctx, "List ID tidak valid", err.Error())
	}

	card, err := c.service.GetByListID(listPublicID)
	if err != nil {
		return utils.NotFound(ctx, "Card tidak ditemukan", err.Error())
	}

	return utils.Success(ctx, "Data Card berhasil diambil", card)
}
