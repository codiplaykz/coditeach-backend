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

var testResultDAO = dao.TestResultDAO{Logger: logmatic.NewLogger()}

func CreateTestResult(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	test_id, err := strconv.Atoi(data["test_id"])
	student_id, err := strconv.Atoi(data["student_id"])
	incorrect_answers, err := strconv.Atoi(data["incorrect_answers"])
	correct_answers, err := strconv.Atoi(data["correct_answers"])
	time_spent, err := strconv.Atoi(data["time_spent"])

	testResult := models.TestResult{
		Test_id:           uint(test_id),
		Student_id:        uint(student_id),
		Incorrect_answers: incorrect_answers,
		Correct_answers:   correct_answers,
		Time_spent:        time_spent,
		Pass_date:         time.Now(),
	}

	err = testResultDAO.Create(&testResult)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create test result, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(testResult)
}

func DeleteTestResult(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	testResult := models.TestResult{
		Id: uint(id),
	}

	err = testResultDAO.Delete(&testResult)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Test result not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete test result, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Test result deleted",
	})
}

func UpdateTestResult(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	test_id, err := strconv.Atoi(data["test_id"])
	student_id, err := strconv.Atoi(data["student_id"])
	incorrect_answers, err := strconv.Atoi(data["incorrect_answers"])
	correct_answers, err := strconv.Atoi(data["correct_answers"])
	time_spent, err := strconv.Atoi(data["time_spent"])

	testResult := models.TestResult{
		Id:                uint(id),
		Test_id:           uint(test_id),
		Student_id:        uint(student_id),
		Incorrect_answers: incorrect_answers,
		Correct_answers:   correct_answers,
		Time_spent:        time_spent,
		Pass_date:         time.Now(),
	}

	err = testResultDAO.Update(&testResult)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Test result not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update test result, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(testResult)
}

func GetTestResult(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	testResult := models.TestResult{
		Id: uint(id),
	}

	err = testResultDAO.GetById(&testResult)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Test result not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get test result, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(testResult)
}
