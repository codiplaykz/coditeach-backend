package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
)

var curriculumLessonDAO = dao.CurriculumLessonDAO{Logger: logmatic.NewLogger()}

func CreateCurriculumLesson(c *fiber.Ctx) error {
	curriculumLesson := new(models.CurriculumLesson)

	err := c.BodyParser(curriculumLesson)

	if err != nil {
		return err
	}

	err = curriculumLessonDAO.Create(curriculumLesson)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create curriculum lesson",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(curriculumLesson)
}

func DeleteCurriculumLesson(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	curriculumLesson := models.CurriculumLesson{
		Id: uint(id),
	}

	err = curriculumLessonDAO.Delete(&curriculumLesson)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete curriculum lesson",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Curriculum lesson deleted",
	})
}

func UpdateCurriculumLesson(c *fiber.Ctx) error {
	curriculumLesson := new(models.CurriculumLesson)

	err := c.BodyParser(curriculumLesson)

	if err != nil {
		return err
	}

	err = curriculumLessonDAO.Update(curriculumLesson)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update curriculum lesson",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(curriculumLesson)
}

func GetCurriculumLesson(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	curriculumLesson := models.CurriculumLesson{
		Id: uint(id),
	}

	err = curriculumLessonDAO.GetById(&curriculumLesson)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Lesson not found",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get curriculum lesson",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(curriculumLesson)
}
