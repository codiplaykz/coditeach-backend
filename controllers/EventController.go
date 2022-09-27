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

var eventDAO = dao.EventDAO{Logger: logmatic.NewLogger()}

func CreateEvent(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	event := models.Event{
		Title:       data["title"],
		Description: data["description"],
		Date:        time.Now(),
	}

	err = eventDAO.Create(&event)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create event, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(event)
}

func DeleteEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	event := models.Event{
		Id: uint(id),
	}

	err = eventDAO.Delete(&event)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Event not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete event, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Event deleted",
	})
}

func UpdateEvent(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])

	event := models.Event{
		Id:          uint(id),
		Title:       data["title"],
		Description: data["description"],
		Date:        time.Now(),
	}

	err = eventDAO.Update(&event)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Event not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update event, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(event)
}

func GetEvent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	event := models.Event{
		Id: uint(id),
	}

	err = eventDAO.GetById(&event)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Event not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get event, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(event)
}
