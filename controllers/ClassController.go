package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"math/rand"
	"strconv"
)

var classDAO = dao.ClassDAO{Logger: logmatic.NewLogger()}

func generateCode() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func CreateClass(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	teacher_id, err := strconv.Atoi(data["teacher_id"])
	school_id, err := strconv.Atoi(data["school_id"])

	class := models.Class{
		Teacher_id: uint(teacher_id),
		School_id:  uint(school_id),
		Name:       data["name"],
		Code:       generateCode(),
	}

	err = classDAO.Create(&class)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create class, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(class)
}

func DeleteClass(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	class := models.Class{
		Id: uint(id),
	}

	err = classDAO.Delete(&class)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Class not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete class, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Class deleted",
	})
}

func UpdateClass(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	teacher_id, err := strconv.Atoi(data["teacher_id"])
	school_id, err := strconv.Atoi(data["school_id"])

	class := models.Class{
		Id:         uint(id),
		Teacher_id: uint(teacher_id),
		School_id:  uint(school_id),
		Name:       data["name"],
		Code:       data["code"],
	}

	err = classDAO.Update(&class)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Class not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update class, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(class)
}

func GetClass(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	class := models.Class{
		Id: uint(id),
	}

	err = classDAO.GetById(&class)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Class not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get class, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(class)
}

func GetClassByCode(c *fiber.Ctx) error {
	code := c.Query("code")

	fmt.Println(code)

	class := models.Class{}

	err := classDAO.GetByParam("code", code, &class)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Class not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get class, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(class)
}
