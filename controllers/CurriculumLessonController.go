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
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	block_id, err := strconv.Atoi(data["block_id"])

	curriculumLesson := models.CurriculumLesson{
		Block_id:    uint(block_id),
		Title:       data["title"],
		Description: data["description"],
		Lesson_type: data["type"],
		Content:     data["content"],
	}

	err = curriculumLessonDAO.Create(&curriculumLesson)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create curriculum lesson",
		})
	}

	result := fiber.Map{
		"id":          curriculumLesson.Id,
		"block_id":    curriculumLesson.Block_id,
		"title":       curriculumLesson.Title,
		"description": curriculumLesson.Description,
		"type":        curriculumLesson.Lesson_type,
		"content":     curriculumLesson.Content,
		"created_at":  curriculumLesson.Created_at,
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(result)
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
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	block_id, err := strconv.Atoi(data["block_id"])

	curriculumLesson := models.CurriculumLesson{
		Id:          uint(id),
		Block_id:    uint(block_id),
		Title:       data["title"],
		Description: data["description"],
		Lesson_type: data["type"],
		Content:     data["content"],
	}

	err = curriculumLessonDAO.Update(&curriculumLesson)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update curriculum lesson",
		})
	}

	result := fiber.Map{
		"id":          curriculumLesson.Id,
		"block_id":    curriculumLesson.Block_id,
		"title":       curriculumLesson.Title,
		"description": curriculumLesson.Description,
		"type":        curriculumLesson.Lesson_type,
		"content":     curriculumLesson.Content,
		"created_at":  curriculumLesson.Created_at,
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(result)
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
