package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
)

var subjectDAO = dao.SubjectDAO{Logger: logmatic.NewLogger()}

func CreateSubject(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	teacher_id, err := strconv.Atoi(data["teacher_id"])

	subject := models.Subject{
		Teacher_id:  uint(teacher_id),
		Name:        data["name"],
		Description: data["description"],
	}

	err = subjectDAO.Create(&subject)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create subject, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(subject)
}

func DeleteSubject(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	subject := models.Subject{
		Id: uint(id),
	}

	err = subjectDAO.Delete(&subject)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Subject not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete subject, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Subject deleted",
	})
}

func UpdateSubject(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	teacher_id, err := strconv.Atoi(data["teacher_id"])

	subject := models.Subject{
		Id:          uint(id),
		Teacher_id:  uint(teacher_id),
		Name:        data["name"],
		Description: data["description"],
	}

	err = subjectDAO.Update(&subject)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Subject not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update subject, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(subject)
}

func GetSubject(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	subject := models.Subject{
		Id: uint(id),
	}

	err = subjectDAO.GetById(&subject)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Subject not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get subject, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(subject)
}
