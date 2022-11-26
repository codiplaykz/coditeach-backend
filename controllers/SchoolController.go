package controllers

import (
	"bytes"
	"coditeach/dao"
	"coditeach/helpers"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"strconv"
)

var schoolDAO = dao.SchoolDAO{Logger: logmatic.NewLogger()}

func CreateSchool(c *fiber.Ctx) error {
	school := new(models.School)

	err := c.BodyParser(school)

	if err != nil {
		return err
	}

	err = schoolDAO.Create(school)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(school)
}

func DeleteSchool(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	school := models.School{
		Id: uint(id),
	}

	err = schoolDAO.Delete(&school)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "School not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete school, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "School deleted",
	})
}

func UpdateSchool(c *fiber.Ctx) error {
	school := new(models.School)

	err := c.BodyParser(school)

	if err != nil {
		return err
	}

	err = schoolDAO.Update(school)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "School not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update school, try later.",
		})
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(school)
}

func GetSchool(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	school := models.School{
		Id: uint(id),
	}

	err = schoolDAO.GetById(&school)

	if err == pgx.ErrNoRows {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "School not found.",
		})
	}

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get school, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(school)
}

func GetAllSchools(c *fiber.Ctx) error {
	schools, err := schoolDAO.GetAll()

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get schools, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(&schools)
}

func GetSchoolAdmins(c *fiber.Ctx) error {
	schoolId, err := strconv.Atoi(c.Query("school_id"))

	school := models.SchoolAdmin{
		School_id: uint(schoolId),
	}

	school_admins, err := schoolAdminDAO.GetBySchoolId(&school)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get school admins, try later.",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(&school_admins)
}

//Admin routes

type CreateSchoolAdminAccountRequestParams struct {
	Email     string `json:"email"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	School_id uint   `json:"school_id"`
}

var schoolAdminDAO = dao.SchoolAdminDAO{Logger: logmatic.NewLogger()}

func CreateSchoolAdminAccount(c *fiber.Ctx) error {
	accountInfo := new(CreateSchoolAdminAccountRequestParams)

	if err := c.BodyParser(accountInfo); err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school admin account, try later.",
		})
	}

	generatedPassword, err := password.Generate(10, 2, 0, false, false)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school admin account, try later.",
		})
	}

	userToCheck := &models.User{
		Email: accountInfo.Email,
	}

	err = userDAO.GetByEmail(userToCheck)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school admin account, try later.",
		})
	}

	if userToCheck.Id != 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User with current email already exist",
		})
	}

	//Crypting password
	pass, err := bcrypt.GenerateFromPassword([]byte(generatedPassword), 14)

	user := &models.User{
		Login:    "",
		Role:     "school_admin",
		Password: pass,
		Name:     accountInfo.Name,
		Surname:  accountInfo.Surname,
		Email:    accountInfo.Email,
	}

	err = userDAO.Create(user)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school admin account, try later.",
		})
	}

	err = schoolAdminDAO.Create(&models.SchoolAdmin{
		User_id:   user.Id,
		School_id: accountInfo.School_id,
	})

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school admin account, try later.",
		})
	}

	//SEND USER DATA TO EMAIL

	// Get html
	var htmlBody bytes.Buffer
	t, err := template.ParseFiles("C:\\Users\\65\\Desktop\\Development\\Backend\\coditeach-backend\\helpers\\SchoolAdminAccountDetailsMail.html")

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school admin account, try later.",
		})
	}

	t.Execute(&htmlBody, struct {
		Name     string
		Surname  string
		Email    string
		Password string
	}{Name: accountInfo.Name, Surname: accountInfo.Surname, Email: accountInfo.Email, Password: generatedPassword})

	err = helpers.SendEmail(accountInfo.Email, "Данные для входа в платформу", htmlBody.String())

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create school admin account, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "School admin account created successfully!",
	})
}
