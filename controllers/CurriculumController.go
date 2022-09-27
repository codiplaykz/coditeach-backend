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

var curriculumDAO = dao.CurriculumDAO{Logger: logmatic.NewLogger()}

func CreateCurriculum(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	teacher_id, err := strconv.Atoi(data["teacher_id"])

	curriculum := models.Curriculum{
		Teacher_id:  uint(teacher_id),
		Title:       data["title"],
		Description: data["description"],
		Created_at:  time.Now(),
	}

	err = curriculumDAO.Create(&curriculum)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create curriculum",
		})
	}

	result := fiber.Map{
		"id":          curriculum.Id,
		"teacher_id":  curriculum.Teacher_id,
		"title":       curriculum.Title,
		"description": curriculum.Description,
		"created_at":  curriculum.Created_at,
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(result)
}

func DeleteCurriculum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	curriculum := models.Curriculum{
		Id: uint(id),
	}

	err = curriculumDAO.Delete(&curriculum)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete curriculum",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Curriculum deleted",
	})
}

func UpdateCurriculum(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update",
		})
	}

	id, err := strconv.Atoi(data["id"])
	teacher_id, err := strconv.Atoi(data["teacher_id"])

	curriculum := models.Curriculum{
		Id:          uint(id),
		Teacher_id:  uint(teacher_id),
		Title:       data["title"],
		Description: data["description"],
		Created_at:  time.Now(),
	}

	err = curriculumDAO.Update(&curriculum)

	if err == pgx.ErrNoRows {
		logger.Error("%s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Curriculum not found",
		})
	}

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update curriculum",
		})
	}

	result := fiber.Map{
		"id":          curriculum.Id,
		"teacher_id":  curriculum.Teacher_id,
		"title":       curriculum.Title,
		"description": curriculum.Description,
		"created_at":  curriculum.Created_at,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(result)
}

func GetCurriculum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	curriculum := models.Curriculum{
		Id: uint(id),
	}

	err = curriculumDAO.GetById(&curriculum)

	if err == pgx.ErrNoRows {
		logger.Error("%s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Curriculum not found",
		})
	}

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get curriculum",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(curriculum)
}
