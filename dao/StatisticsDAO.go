package dao

import (
	"coditeach/database"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/mborders/logmatic"
)

type StatisticsDAO struct {
	Logger *logmatic.Logger
}

func (s *StatisticsDAO) GetTotalSchoolsStatistics() int {
	count := 0

	err := database.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM schools").Scan(&count)

	if err != nil {
		s.Logger.Error("Unable to get schools statistics.")
		return 0
	}

	s.Logger.Info("Count of schools: %v", count)

	return count
}

func (s *StatisticsDAO) GetTotalTeachersStatistics() int {
	count := 0

	err := database.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM teachers").Scan(&count)

	if err != nil {
		s.Logger.Error("Unable to get teachers statistics.")
		return 0
	}

	s.Logger.Info("Count of teachers: %v", count)

	return count
}

func (s *StatisticsDAO) GetTotalStudentsStatistics() int {
	count := 0

	err := database.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM students").Scan(&count)

	if err != nil {
		s.Logger.Error("Unable to get students statistics.")
		return 0
	}

	s.Logger.Info("Count of students: %v", count)

	return count
}

func (s *StatisticsDAO) GetTotalParentStatistics() int {
	count := 0

	err := database.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM parents").Scan(&count)

	if err != nil {
		s.Logger.Error("Unable to get parents statistics.")
		return 0
	}

	s.Logger.Info("Count of parents: %v", count)

	return count
}

func (s *StatisticsDAO) GetStudentsStatisticsBySchool(id int) int {
	count := 0

	err := database.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM students inner join classes on students.class_id=classes.id where classes.school_id=$1", id).Scan(&count)

	if err != nil {
		s.Logger.Error("Unable to get students statistics by school.")
		return 0
	}

	s.Logger.Info("Count of students: %v", count)

	return count
}

func (s *StatisticsDAO) GetTeachersStatisticsBySchool(id int) int {
	count := 0

	err := database.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM teachers where school_id=$1", id).Scan(&count)

	if err != nil {
		s.Logger.Error("Unable to get teachers statistics by school.")
		return 0
	}

	s.Logger.Info("Count of teachers: %v", count)

	return count
}

func (s *StatisticsDAO) GetParentsStatisticsBySchool(id int) int {
	count := 0

	err := database.DB.QueryRow(context.Background(),
		"SELECT COUNT(*) FROM parents inner join students on parents.student_id=students.id inner join classes on classes.id=students.class_id where classes.school_id=$1", id).Scan(&count)

	if err != nil {
		s.Logger.Error("Unable to get parents statistics by school.")
		return 0
	}

	s.Logger.Info("Count of parents: %v", count)

	return count
}

func (s *StatisticsDAO) GenerateReport() fiber.Map {
	rows, err := database.DB.Query(context.Background(), "select concat(information, '| CREATION DATE:', created_at) from statistics")

	if err != nil {
		return nil
	}

	defer rows.Close()

	var res []string

	for rows.Next() {
		str := ""
		err = rows.Scan(&str)
		if err != nil {
			return nil
		}
		res = append(res, str)
	}

	return fiber.Map{
		"stats": res,
	}
}
