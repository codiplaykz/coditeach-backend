package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
)

var questionDAO = dao.QuestionDAO{Logger: logmatic.NewLogger()}

func CreateQuestion(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	test_id, err := strconv.Atoi(data["test_id"])

	question := models.Question{
		Test_id: uint(test_id),
		Text:    data["text"],
	}

	err = questionDAO.Create(&question)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create question, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(question)
}

func DeleteQuestion(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	question := models.Question{
		Id: uint(id),
	}

	err = questionDAO.Delete(&question)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Question not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete question, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Question deleted",
	})
}

func UpdateQuestion(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	test_id, err := strconv.Atoi(data["test_id"])

	question := models.Question{
		Id:      uint(id),
		Test_id: uint(test_id),
		Text:    data["text"],
	}

	err = questionDAO.Update(&question)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Question not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update question, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(question)
}

func GetQuestion(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	question := models.Question{
		Id: uint(id),
	}

	err = questionDAO.GetById(&question)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Question not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get question, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(question)
}
