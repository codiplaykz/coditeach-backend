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

var scheduleLessonDAO = dao.ScheduleLessonDAO{Logger: logmatic.NewLogger()}

func CreateScheduleLesson(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	schedule_id, err := strconv.Atoi(data["schedule_id"])

	scheduleLesson := models.ScheduleLesson{
		Schedule_id: uint(schedule_id),
		Start_time:  time.Now(),
		End_time:    time.Now(),
	}

	err = scheduleLessonDAO.Create(&scheduleLesson)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create schedule lesson, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(scheduleLesson)
}

func DeleteScheduleLesson(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	scheduleLesson := models.ScheduleLesson{
		Id: uint(id),
	}

	err = scheduleLessonDAO.Delete(&scheduleLesson)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Schedule lesson not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete schedule lesson, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Schedule lesson deleted",
	})
}

func UpdateScheduleLesson(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	schedule_id, err := strconv.Atoi(data["schedule_id"])

	scheduleLesson := models.ScheduleLesson{
		Id:          uint(id),
		Schedule_id: uint(schedule_id),
		Start_time:  time.Now(),
		End_time:    time.Now(),
	}

	err = scheduleLessonDAO.Update(&scheduleLesson)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Schedule lesson not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update schedule lesson, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(scheduleLesson)
}

func GetScheduleLesson(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	scheduleLesson := models.ScheduleLesson{
		Id: uint(id),
	}

	err = scheduleLessonDAO.GetById(&scheduleLesson)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Schedule lesson not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get schedule lesson, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(scheduleLesson)
}
