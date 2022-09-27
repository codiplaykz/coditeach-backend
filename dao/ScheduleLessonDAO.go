package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type ScheduleLessonDAO struct {
	Logger *logmatic.Logger
}

func (s *ScheduleLessonDAO) Create(scheduleLesson *models.ScheduleLesson) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into schedule_lessons (schedule_id, start_time, end_time) VALUES($1,$2,$3) returning id",
		scheduleLesson.Schedule_id,
		scheduleLesson.Start_time,
		scheduleLesson.End_time).Scan(&scheduleLesson.Id)

	if err != nil {
		s.Logger.Error("Unable to create schedule lesson.")
		return err
	}

	s.Logger.Info("Schedule lesson created.")

	return nil
}

func (s *ScheduleLessonDAO) Update(scheduleLesson *models.ScheduleLesson) error {
	err := database.DB.QueryRow(context.Background(),
		"update schedule_lessons set schedule_id=$1, start_time=$2, end_time=$3 where id=$4 returning id",
		scheduleLesson.Schedule_id,
		scheduleLesson.Start_time,
		scheduleLesson.End_time,
		scheduleLesson.Id).Scan(&scheduleLesson.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Schedule lesson not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to update schedule lesson.")
		return err
	}

	s.Logger.Info("Updated schedule lesson with id: %v", scheduleLesson.Id)

	return nil
}

func (s *ScheduleLessonDAO) Delete(scheduleLesson *models.ScheduleLesson) error {
	err := database.DB.QueryRow(context.Background(), "delete from schedule_lessons where id=$1 returning id", scheduleLesson.Id).Scan(&scheduleLesson.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Schedule lesson not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to delete schedule lesson with id: %v", scheduleLesson.Id)
		return err
	}
	s.Logger.Info("Deleted schedule lesson with id: %v", scheduleLesson.Id)

	return nil
}

func (s *ScheduleLessonDAO) GetById(scheduleLesson *models.ScheduleLesson) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from schedule_lessons where id=$1",
		scheduleLesson.Id)

	err := row.Scan(&scheduleLesson.Id, &scheduleLesson.Schedule_id, &scheduleLesson.Start_time, &scheduleLesson.End_time)

	if err != nil {
		s.Logger.Error("Unable to get schedule lesson with id: %v.", scheduleLesson.Id)
		return err
	}

	return nil
}
