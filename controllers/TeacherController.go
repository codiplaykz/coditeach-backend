package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

var teacherDAO = dao.TeacherDAO{Logger: logmatic.NewLogger()}

//Admin role

func CreateTeacher(c *fiber.Ctx) error {
	teacher := new(models.Teacher)

	err := c.BodyParser(teacher)

	if err != nil {
		return err
	}

	err = teacherDAO.Create(teacher)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create teacher, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(teacher)
}

func DeleteTeacher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	teacher := models.Teacher{
		Id: uint(id),
	}

	err = teacherDAO.Delete(&teacher)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Teacher not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete teacher, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Teacher deleted",
	})
}

func UpdateTeacher(c *fiber.Ctx) error {
	teacher := new(models.Teacher)

	err := c.BodyParser(teacher)

	if err != nil {
		return err
	}

	err = teacherDAO.Update(teacher)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Teacher not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update teacher, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(teacher)
}

func GetTeacher(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	teacher := models.Teacher{
		Id: uint(id),
	}

	err = teacherDAO.GetById(&teacher)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Teacher not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get teacher, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(teacher)
}

func GetAllTeachers(c *fiber.Ctx) error {
	teachers, err := teacherDAO.GetAll()

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get teachers, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(&teachers)
}

//School admin role

func CreateTeacherAccount(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	//Crypting password
	pass, err := bcrypt.GenerateFromPassword([]byte(data["pass"]), 14)
	school_id, err := strconv.Atoi(data["school_id"])

	teacher := dao.TeacherAccount{
		Role:      "teacher",
		Login:     data["login"],
		Password:  string(pass),
		Name:      data["name"],
		Surname:   data["surname"],
		Email:     data["email"],
		School_id: uint(school_id),
	}

	err = teacherDAO.CreateAccount(&teacher)

	if err != nil {
		if err.Error() == "email already exists" {
			logger.Error("ERROR: %s", err)
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Email already exists",
			})
		} else {
			logger.Error("ERROR: %s", err)
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Unable to get teacher, try later.",
			})
		}
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(teacher)
}
