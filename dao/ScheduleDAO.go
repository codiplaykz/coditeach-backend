package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type ScheduleDAO struct {
	Logger *logmatic.Logger
}

func (s *ScheduleDAO) Create(schedule *models.Schedule) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into schedules (name, subject_id, class_id) VALUES($1,$2,$3) returning id",
		schedule.Name,
		schedule.Subject_id,
		schedule.Class_id).Scan(&schedule.Id)

	if err != nil {
		s.Logger.Error("Unable to create schedule.")
		return err
	}

	s.Logger.Info("Schedule created.")

	return nil
}

func (s *ScheduleDAO) Update(schedule *models.Schedule) error {
	err := database.DB.QueryRow(context.Background(),
		"update schedules set name=$1, subject_id=$2, class_id=$3 where id=$4 returning id",
		schedule.Name,
		schedule.Subject_id,
		schedule.Class_id,
		schedule.Id).Scan(&schedule.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Schedule not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to update schedule.")
		return err
	}

	s.Logger.Info("Updated schedule with id: %v", schedule.Id)

	return nil
}

func (s *ScheduleDAO) Delete(schedule *models.Schedule) error {
	err := database.DB.QueryRow(context.Background(), "delete from schedules where id=$1 returning id", schedule.Id).Scan(&schedule.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Schedule not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to delete class with id: %v", schedule.Id)
		return err
	}
	s.Logger.Info("Deleted schedule with id: %v", schedule.Id)

	return nil
}

func (s *ScheduleDAO) GetById(schedule *models.Schedule) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from schedules where id=$1",
		schedule.Id)

	err := row.Scan(&schedule.Id, &schedule.Name, &schedule.Subject_id, &schedule.Class_id)

	if err != nil {
		s.Logger.Error("Unable to get schedule with id: %v.", schedule.Id)
		return err
	}

	return nil
}
