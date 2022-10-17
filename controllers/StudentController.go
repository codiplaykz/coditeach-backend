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

var studentDAO = dao.StudentDAO{Logger: logmatic.NewLogger()}

//Admin role

func CreateStudent(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create student, try later.",
		})
	}

	user_id, err := strconv.Atoi(data["user_id"])
	class_id, err := strconv.Atoi(data["class_id"])

	student := models.Student{
		User_id:  uint(user_id),
		Class_id: uint(class_id),
	}

	err = studentDAO.Create(&student)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create student, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(student)
}

func DeleteStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete student, try later.",
		})
	}

	student := models.Student{
		Id: uint(id),
	}

	err = studentDAO.Delete(&student)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Student not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete student, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Student deleted",
	})
}

func UpdateStudent(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	id, err := strconv.Atoi(data["id"])
	user_id, err := strconv.Atoi(data["user_id"])
	class_id, err := strconv.Atoi(data["class_id"])

	student := models.Student{
		Id:       uint(id),
		User_id:  uint(user_id),
		Class_id: uint(class_id),
	}

	err = studentDAO.Update(&student)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Student not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update student, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(student)
}

func GetStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	student := models.Student{
		Id: uint(id),
	}

	err = studentDAO.GetById(&student)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Student not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get student, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(student)
}

//Teacher role

func CreateStudentAccount(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	//Crypting password
	pass, err := bcrypt.GenerateFromPassword([]byte(data["pass"]), 14)
	class_id, err := strconv.Atoi(data["class_id"])

	student := dao.StudentAccount{
		Role:     "student",
		Login:    data["login"],
		Password: string(pass),
		Name:     data["name"],
		Surname:  data["surname"],
		Email:    data["email"],
		Class_id: uint(class_id),
	}

	err = studentDAO.CreateAccount(&student)

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
				"message": "Unable to create student account, try later.",
			})
		}
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(student)
}

func RegisterStudentAccount(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}

	class := models.Class{}

	err = classDAO.GetByParam("code", data["code"], &class)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Class with current code not found",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get student, try later.",
		})
	}

	//Crypting password
	pass, err := bcrypt.GenerateFromPassword([]byte(data["pass"]), 14)

	student := dao.StudentAccount{
		Role:     "student",
		Login:    data["login"],
		Password: string(pass),
		Name:     data["name"],
		Surname:  data["surname"],
		Email:    data["email"],
		Class_id: class.Id,
	}

	err = studentDAO.CreateAccount(&student)

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
				"message": "Unable to create student account, try later.",
			})
		}
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(student)
}

func GetAllStudents(c *fiber.Ctx) error {
	students, err := studentDAO.GetAll()

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get students, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(&students)
}
