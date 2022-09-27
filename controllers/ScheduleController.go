package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"strconv"
)

var scheduleDAO = dao.ScheduleDAO{Logger: logmatic.NewLogger()}

func CreateSchedule(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	subject_id, err := strconv.Atoi(data["subject_id"])
	class_id, err := strconv.Atoi(data["class_id"])

	schedule := models.Schedule{
		Name:       data["name"],
		Subject_id: uint(subject_id),
		Class_id:   uint(class_id),
	}

	err = scheduleDAO.Create(&schedule)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create schedule, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(schedule)
}

func DeleteSchedule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	schedule := models.Schedule{
		Id: uint(id),
	}

	err = scheduleDAO.Delete(&schedule)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Schedule not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete schedule, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Schedule deleted",
	})
}

func UpdateSchedule(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	subject_id, err := strconv.Atoi(data["subject_id"])
	class_id, err := strconv.Atoi(data["class_id"])

	schedule := models.Schedule{
		Id:         uint(id),
		Name:       data["name"],
		Subject_id: uint(subject_id),
		Class_id:   uint(class_id),
	}

	err = scheduleDAO.Update(&schedule)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Schedule not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update schedule, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(schedule)
}

func GetSchedule(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	schedule := models.Schedule{
		Id: uint(id),
	}

	err = scheduleDAO.GetById(&schedule)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Schedule not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get schedule, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(schedule)
}
