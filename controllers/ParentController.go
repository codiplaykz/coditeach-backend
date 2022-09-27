package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
)

var parentDAO = dao.ParentDAO{Logger: logmatic.NewLogger()}

func CreateParent(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	user_id, err := strconv.Atoi(data["user_id"])
	student_id, err := strconv.Atoi(data["student_id"])

	parent := models.Parent{
		User_id:    uint(user_id),
		Student_id: uint(student_id),
	}

	err = parentDAO.Create(&parent)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create parent, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(parent)
}

func DeleteParent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	parent := models.Parent{
		Id: uint(id),
	}

	err = parentDAO.Delete(&parent)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Parent not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete parent, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Parent deleted",
	})
}

func UpdateParent(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	user_id, err := strconv.Atoi(data["user_id"])
	student_id, err := strconv.Atoi(data["student_id"])

	parent := models.Parent{
		Id:         uint(id),
		User_id:    uint(user_id),
		Student_id: uint(student_id),
	}

	err = parentDAO.Update(&parent)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Parent not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update parent, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(parent)
}

func GetParent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	parent := models.Parent{
		Id: uint(id),
	}

	err = parentDAO.GetById(&parent)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Parent not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get parent, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(parent)
}
