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

// @Summary      delete school
// @Tags         school
// @Description  delete school
// @ID           delete-school
// @Accept       json
// @Produce      json
// @Param        id   query     number  true  "school id"
// @Success      200  {object}  object{message=string}
// @Failure      500   {object}  object{message=string}
// @Failure      404  {object}  object{message=string}
// @Router       /api/school/delete [delete]
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

// @Summary      update school
// @Tags         school
// @Description  update school
// @ID           update-school
// @Accept       json
// @Produce      json
// @Param        input body object{id=number,name=string,location=string} true "school info"
// @Success      202   {object}  object{id=number,name=string,location=string}
// @Failure      500    {object}  object{message=string}
// @Failure      404   {object}  object{message=string}
// @Router       /api/school/update [put]
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

// @Summary      get school
// @Tags         school
// @Description  get school
// @ID           get-school
// @Accept       json
// @Produce      json
// @Param        id   query     number  true  "school id"
// @Success      200   {object}  object{id=number,name=string,location=string}
// @Failure      500    {object}  object{message=string}
// @Failure      404   {object}  object{message=string}
// @Router       /api/school/get [get]
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
