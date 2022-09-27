package controllers

import (
	"coditeach/dao"
	"coditeach/database"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mborders/logmatic"
	"strconv"
	"time"
)

var statisticsDAO = dao.StatisticsDAO{Logger: logmatic.NewLogger()}

func GenerateTotalStatistics(c *fiber.Ctx) error {
	schools_count := statisticsDAO.GetTotalSchoolsStatistics()
	students_count := statisticsDAO.GetTotalStudentsStatistics()
	teachers_count := statisticsDAO.GetTotalTeachersStatistics()
	parents_count := statisticsDAO.GetTotalParentStatistics()

	stats := fmt.Sprintf("total_schools: %v, total_students: %v, total_teachers: %v, total_parents: %v", schools_count, students_count, teachers_count, parents_count)

	_, err := database.DB.Exec(context.Background(), "INSERT INTO STATISTICS (information, created_at) VALUES($1,$2)", stats, time.Now())

	if err != nil {
		return nil
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"total_schools":  schools_count,
		"total_students": students_count,
		"total_teachers": teachers_count,
		"total_parents":  parents_count,
	})
}

func GenerateSchoolStatistics(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		return nil
	}

	students_count := statisticsDAO.GetStudentsStatisticsBySchool(id)
	teachers_count := statisticsDAO.GetTeachersStatisticsBySchool(id)
	parents_count := statisticsDAO.GetParentsStatisticsBySchool(id)

	stats := fmt.Sprintf("total_students: %v, total_teachers: %v, total_parents: %v by school id %v", students_count, teachers_count, parents_count, id)

	_, err = database.DB.Exec(context.Background(), "INSERT INTO STATISTICS (information, created_at) VALUES($1,$2)", stats, time.Now())

	if err != nil {
		return nil
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"total_students": students_count,
		"total_teachers": teachers_count,
		"total_parents":  parents_count,
	})
}

func GenerateReport(c *fiber.Ctx) error {
	return c.JSON(statisticsDAO.GenerateReport())
}
