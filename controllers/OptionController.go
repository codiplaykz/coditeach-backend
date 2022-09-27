package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
)

var optionDAO = dao.OptionDAO{Logger: logmatic.NewLogger()}

func CreateOption(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	question_id, err := strconv.Atoi(data["question_id"])
	is_correct, err := strconv.ParseBool(data["is_correct"])

	option := models.Option{
		Question_id: uint(question_id),
		Text:        data["text"],
		Is_correct:  is_correct,
	}

	err = optionDAO.Create(&option)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create option, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(option)
}

func DeleteOption(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	option := models.Option{
		Id: uint(id),
	}

	err = optionDAO.Delete(&option)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Option not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete option, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Option deleted",
	})
}

func UpdateOption(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	question_id, err := strconv.Atoi(data["question_id"])
	is_correct, err := strconv.ParseBool(data["is_correct"])

	option := models.Option{
		Id:          uint(id),
		Question_id: uint(question_id),
		Text:        data["text"],
		Is_correct:  is_correct,
	}

	err = optionDAO.Update(&option)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Option not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update option, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(option)
}

func GetOption(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	option := models.Option{
		Id: uint(id),
	}

	err = optionDAO.GetById(&option)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Option not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get option, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(option)
}
