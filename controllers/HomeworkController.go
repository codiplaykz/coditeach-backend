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

var homeworkDAO = dao.HomeworkDAO{Logger: logmatic.NewLogger()}

func CreateHomework(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	subject_id, err := strconv.Atoi(data["subject_id"])

	homework := models.Homework{
		Name:        data["name"],
		Description: data["description"],
		Deadline:    time.Now(),
		Subject_id:  uint(subject_id),
	}

	err = homeworkDAO.Create(&homework)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create homework, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(homework)
}

func DeleteHomework(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	homework := models.Homework{
		Id: uint(id),
	}

	err = homeworkDAO.Delete(&homework)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Homework not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete homework, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Homework deleted",
	})
}

func UpdateHomework(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	subject_id, err := strconv.Atoi(data["subject_id"])

	homework := models.Homework{
		Id:          uint(id),
		Name:        data["name"],
		Description: data["description"],
		Deadline:    time.Now(),
		Subject_id:  uint(subject_id),
	}

	err = homeworkDAO.Update(&homework)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Homework not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update homework, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(homework)
}

func GetHomework(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	homework := models.Homework{
		Id: uint(id),
	}

	err = homeworkDAO.GetById(&homework)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Homework not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get homework, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(homework)
}
