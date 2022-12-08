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

var curriculumDAO = dao.CurriculumDAO{Logger: logmatic.NewLogger()}

func CreateCurriculum(c *fiber.Ctx) error {
	curriculum := new(models.Curriculum)

	err := c.BodyParser(curriculum)

	if err != nil {
		return err
	}

	err = curriculumDAO.Create(curriculum)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create curriculum",
		})
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(curriculum)
}

func DeleteCurriculum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	curriculum := models.Curriculum{
		Id: uint(id),
	}

	err = curriculumDAO.Delete(&curriculum)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to delete curriculum",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "Curriculum deleted",
	})
}

func UpdateCurriculum(c *fiber.Ctx) error {
	curriculum := new(models.Curriculum)

	err := c.BodyParser(curriculum)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update",
		})
	}

	err = curriculumDAO.Update(curriculum)

	if err == pgx.ErrNoRows {
		logger.Error("%s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Curriculum not found",
		})
	}

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to update curriculum",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(curriculum)
}

func GetCurriculum(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	curriculum := models.Curriculum{
		Id: uint(id),
	}

	err = curriculumDAO.GetById(&curriculum)

	if err == pgx.ErrNoRows {
		logger.Error("%s", err)
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "Curriculum not found",
		})
	}

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get curriculum",
		})
	}

	c.Status(fiber.StatusOK)
	return c.JSON(curriculum)
}

type ComposeCurriculumResponse struct {
	User_id     int
	Title       string
	Description string
	Modules     []struct {
		Title       string
		Description string
		Blocks      []struct {
			Title       string
			Description string
			Lessons     []struct {
				Title       string
				Description string
				Duration    int
				Content     string
			}
		}
	}
}

func ComposeCurriculum(c *fiber.Ctx) error {
	response := new(ComposeCurriculumResponse)

	err := c.BodyParser(response)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create curriculum",
		})
	}

	curriculum := models.Curriculum{
		User_id:     uint(response.User_id),
		Title:       response.Title,
		Description: response.Description,
		Created_at:  time.Now(),
	}

	err = curriculumDAO.Create(&curriculum)

	if err != nil {
		logger.Error("%s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to create curriculum",
		})
	}

	for _, module := range response.Modules {
		moduleToInsert := models.Module{
			Curriculum_id: curriculum.Id,
			Title:         module.Title,
			Description:   module.Description,
			Created_at:    time.Now(),
		}
		err = moduleDAO.Create(&moduleToInsert)

		if err != nil {
			logger.Error("%s", err)
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Unable to create curriculum",
			})
		}

		for _, block := range module.Blocks {
			blockToInsert := models.Block{
				Module_id:   moduleToInsert.Id,
				Title:       block.Title,
				Description: block.Description,
				Created_at:  time.Now(),
			}
			err = blockDAO.Create(&blockToInsert)

			if err != nil {
				logger.Error("%s", err)
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(fiber.Map{
					"message": "Unable to create curriculum",
				})
			}

			for _, lesson := range block.Lessons {
				lessonToInsert := models.CurriculumLesson{
					Block_id:    blockToInsert.Id,
					Title:       lesson.Title,
					Description: lesson.Description,
					Duration:    lesson.Duration,
					Content:     lesson.Content,
					Created_at:  time.Now(),
				}
				err = curriculumLessonDAO.Create(&lessonToInsert)

				if err != nil {
					logger.Error("%s", err)
					c.Status(fiber.StatusInternalServerError)
					return c.JSON(fiber.Map{
						"message": "Unable to create curriculum",
					})
				}
			}
		}
	}

	c.Status(fiber.StatusCreated)
	return c.JSON(fiber.Map{
		"message": "Curriculum created",
	})
}

func GetAllFullCurriculums(c *fiber.Ctx) error {
	curriculums, err := curriculumDAO.GetAll()

	if err != nil {
		logger.Error("ERROR: %s", err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Unable to get projects, try later.",
		})
	}

	for _, curriculum := range curriculums {
		modules, err := moduleDAO.GetAllByCurriculumId(int(curriculum["id"].(int32)))

		if err != nil {
			logger.Error("ERROR: %s", err)
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "Unable to get projects, try later.",
			})
		}

		for _, module := range modules {
			blocks, err := blockDAO.GetAllByModuleId(int(module["id"].(int32)))

			if err != nil {
				logger.Error("ERROR: %s", err)
				c.Status(fiber.StatusInternalServerError)
				return c.JSON(fiber.Map{
					"message": "Unable to get projects, try later.",
				})
			}

			for _, block := range blocks {
				lessons, err := curriculumLessonDAO.GetAllByBlockId(int(block["id"].(int32)))
				if err != nil {
					logger.Error("ERROR: %s", err)
					c.Status(fiber.StatusInternalServerError)
					return c.JSON(fiber.Map{
						"message": "Unable to get projects, try later.",
					})
				}
				block["lessons"] = lessons
			}
			module["blocks"] = blocks
		}

		curriculum["modules"] = modules
	}

	return c.JSON(curriculums)
}
