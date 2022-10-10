package controllers

import (
	"coditeach/dao"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
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

	result := fiber.Map{
		"id":       school.Id,
		"name":     school.Name,
		"location": school.Location,
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(result)
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

	result := fiber.Map{
		"id":       school.Id,
		"name":     school.Name,
		"location": school.Location,
	}

	c.Status(fiber.StatusAccepted)
	return c.JSON(result)
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

	result := fiber.Map{
		"id":       school.Id,
		"name":     school.Name,
		"location": school.Location,
	}

	c.Status(fiber.StatusOK)
	return c.JSON(result)
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
