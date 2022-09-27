package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type HomeworkDAO struct {
	Logger *logmatic.Logger
}

func (h *HomeworkDAO) Create(homework *models.Homework) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into homeworks (name, description, deadline, subject_id) VALUES($1,$2,$3,$4) returning id",
		homework.Name,
		homework.Description,
		homework.Deadline,
		homework.Subject_id).Scan(&homework.Id)

	if err != nil {
		h.Logger.Error("Unable to create homework.")
		return err
	}

	h.Logger.Info("Homework created.")

	return nil
}

func (h *HomeworkDAO) Update(homework *models.Homework) error {
	err := database.DB.QueryRow(context.Background(),
		"update homeworks set name=$1, description=$2, deadline=$3, subject_id=$4 where id=$5 returning id",
		homework.Name,
		homework.Description,
		homework.Deadline,
		homework.Subject_id,
		homework.Id).Scan(&homework.Id)

	if err == pgx.ErrNoRows {
		h.Logger.Error("Homework not found")
		return err
	}

	if err != nil {
		h.Logger.Error("Unable to update homework.")
		return err
	}

	h.Logger.Info("Updated homework with id: %v", homework.Id)

	return nil
}

func (h *HomeworkDAO) Delete(homework *models.Homework) error {
	err := database.DB.QueryRow(context.Background(), "delete from homeworks where id=$1 returning id", homework.Id).Scan(&homework.Id)

	if err == pgx.ErrNoRows {
		h.Logger.Error("Homework not found")
		return err
	}

	if err != nil {
		h.Logger.Error("Unable to delete homework with id: %v", homework.Id)
		return err
	}
	h.Logger.Info("Deleted homework with id: %v", homework.Id)

	return nil
}

func (h *HomeworkDAO) GetById(homework *models.Homework) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from homeworks where id=$1",
		homework.Id)

	err := row.Scan(&homework.Id, &homework.Name, &homework.Description, &homework.Description, &homework.Subject_id)

	if err != nil {
		h.Logger.Error("Unable to get homework with id: %v.", homework.Id)
		return err
	}

	return nil
}
