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

var testDAO = dao.TestDAO{Logger: logmatic.NewLogger()}

func CreateTest(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	duration, err := strconv.Atoi(data["duration"])
	teacher_id, err := strconv.Atoi(data["teacher_id"])

	test := models.Test{
		Name:        data["name"],
		Description: data["description"],
		Duration:    duration,
		Created_at:  time.Now(),
		Teacher_id:  uint(teacher_id),
	}

	err = testDAO.Create(&test)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create test, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(test)
}

func DeleteTest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	test := models.Test{
		Id: uint(id),
	}

	err = testDAO.Delete(&test)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Test not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete test, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Test deleted",
	})
}

func UpdateTest(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	duration, err := strconv.Atoi(data["duration"])
	teacher_id, err := strconv.Atoi(data["teacher_id"])

	test := models.Test{
		Id:          uint(id),
		Name:        data["name"],
		Description: data["description"],
		Duration:    duration,
		Created_at:  time.Now(),
		Teacher_id:  uint(teacher_id),
	}

	err = testDAO.Update(&test)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Test not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update test, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(test)
}

func GetTest(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	test := models.Test{
		Id: uint(id),
	}

	err = testDAO.GetById(&test)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Test not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get test, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(test)
}
