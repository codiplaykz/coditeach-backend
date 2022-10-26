package controllers

import (
	"coditeach/dao"
	"coditeach/helpers"
	"coditeach/models"
	"github.com/gofiber/fiber/v2"
	"github.com/mborders/logmatic"
	"github.com/rs/xid"
	"path/filepath"
)

var projectDAO = dao.ProjectDAO{Logger: logmatic.NewLogger()}

func CreateProject(c *fiber.Ctx) error {

	name := c.FormValue("name")
	desc := c.FormValue("desc")
	p_type := c.FormValue("p_type")
	level := c.FormValue("level")
	t_comp := c.FormValue("tech_components")
	cover_image, err := c.FormFile("cover_image")
	schema_image, err := c.FormFile("schema_image")

	cover_image_url := xid.New().String() + filepath.Ext(cover_image.Filename)
	schema_image_url := xid.New().String() + filepath.Ext(schema_image.Filename)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get Buffer from file
	cover_buffer, err := cover_image.Open()
	schema_buffer, err := schema_image.Open()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	defer cover_buffer.Close()
	defer schema_buffer.Close()

	helpers.UploadImageToSpace(cover_image_url, cover_buffer)
	helpers.UploadImageToSpace(schema_image_url, schema_buffer)

	url1 := "https://cp-space.fra1.digitaloceanspaces.com/projects-images/" + cover_image_url
	url2 := "https://cp-space.fra1.digitaloceanspaces.com/projects-images/" + schema_image_url

	project := models.Project{
		Name:            name,
		Type:            p_type,
		Level:           level,
		Tech_components: t_comp,
		Description:     desc,
		Creator_Id:      0,
		Source_Code:     "",
		Block_code:      "",
		Cover_img_url:   url1,
		Scheme_img_url:  url2,
	}

	err = projectDAO.Create(&project)

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create project, try later.",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(project)
}
