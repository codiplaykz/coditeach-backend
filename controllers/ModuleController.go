package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
)

var moduleDAO = dao.ModuleDAO{Logger: logmatic.NewLogger()}

func CreateModule(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	curriculumId, err := strconv.Atoi(data["curriculum_id"])

	module := models.Module{
		Curriculum_id: uint(curriculumId),
		Title:         data["title"],
		Description:   data["description"],
	}

	err = moduleDAO.Create(&module)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create module",
		})
	}

	result := fiber.Map{
		"id":            module.Id,
		"curriculum_id": module.Curriculum_id,
		"title":         module.Title,
		"description":   module.Description,
		"created_at":    module.Created_at,
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(result)
}

func DeleteModule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	module := models.Module{
		Id: uint(id),
	}

	err = moduleDAO.Delete(&module)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete module",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Module deleted",
	})
}

func UpdateModule(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	curriculumId, err := strconv.Atoi(data["curriculum_id"])

	module := models.Module{
		Id:            uint(id),
		Curriculum_id: uint(curriculumId),
		Title:         data["title"],
		Description:   data["description"],
	}

	err = moduleDAO.Update(&module)

	if err == pgx.ErrNoRows {
		logger.Error("%s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Module not found.",
		})
	}

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update module.",
		})
	}

	result := fiber.Map{
		"id":            module.Id,
		"curriculum_id": module.Curriculum_id,
		"title":         module.Title,
		"description":   module.Description,
		"created_at":    module.Created_at,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(result)
}

func GetModule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	module := models.Module{
		Id: uint(id),
	}

	err = moduleDAO.GetById(&module)

	if err == pgx.ErrNoRows {
		logger.Error("%s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Module not found.",
		})
	}

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update module.",
		})
	}

	result := fiber.Map{
		"id":            module.Id,
		"curriculum_id": module.Curriculum_id,
		"title":         module.Title,
		"description":   module.Description,
		"created_at":    module.Created_at,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(result)
}
