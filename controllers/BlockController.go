package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
	"time"
)

var blockDAO = dao.BlockDAO{Logger: logmatic.NewLogger()}

func CreateBlock(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	module_id, err := strconv.Atoi(data["module_id"])

	block := models.Block{
		Module_id:  uint(module_id),
		Title:      data["title"],
		Created_at: time.Now(),
	}

	err = blockDAO.Create(&block)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create block, try later.",
		})
	}

	result := fiber.Map{
		"id":         block.Id,
		"module_id":  block.Module_id,
		"title":      block.Title,
		"created_at": block.Created_at,
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(result)
}

func DeleteBlock(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	block := models.Block{
		Id: uint(id),
	}

	err = blockDAO.Delete(&block)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete block, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Block deleted",
	})
}

func UpdateBlock(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	module_id, err := strconv.Atoi(data["module_id"])

	block := models.Block{
		Id:         uint(id),
		Module_id:  uint(module_id),
		Title:      data["title"],
		Created_at: time.Now(),
	}

	err = blockDAO.Update(&block)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update block, try later.",
		})
	}

	result := fiber.Map{
		"id":         block.Id,
		"module_id":  block.Module_id,
		"title":      block.Title,
		"created_at": block.Created_at,
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(result)
}

func GetBlock(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	block := models.Block{
		Id: uint(id),
	}

	err = blockDAO.GetById(&block)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Block not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get block, try later.",
		})
	}

	result := fiber.Map{
		"id":         block.Id,
		"module_id":  block.Module_id,
		"title":      block.Title,
		"created_at": block.Created_at,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(result)
}
