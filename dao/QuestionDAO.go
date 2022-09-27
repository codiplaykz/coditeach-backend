package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type QuestionDAO struct {
	Logger *logmatic.Logger
}

func (q *QuestionDAO) Create(question *models.Question) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into questions (test_id, text) VALUES($1,$2) returning id",
		question.Test_id,
		question.Text).Scan(&question.Id)

	if err != nil {
		q.Logger.Error("Unable to create question.")
		return err
	}

	q.Logger.Info("Question created.")

	return nil
}

func (q *QuestionDAO) Update(question *models.Question) error {
	err := database.DB.QueryRow(context.Background(),
		"update questions set test_id=$1, text=$2 where id=$3 returning id",
		question.Test_id,
		question.Text,
		question.Id).Scan(&question.Id)

	if err == pgx.ErrNoRows {
		q.Logger.Error("Question not found")
		return err
	}

	if err != nil {
		q.Logger.Error("Unable to update question.")
		return err
	}

	q.Logger.Info("Updated question with id: %v", question.Id)

	return nil
}

func (q *QuestionDAO) Delete(question *models.Question) error {
	err := database.DB.QueryRow(context.Background(), "delete from questions where id=$1 returning id", question.Id).Scan(&question.Id)

	if err == pgx.ErrNoRows {
		q.Logger.Error("Question not found")
		return err
	}

	if err != nil {
		q.Logger.Error("Unable to delete question with id: %v", question.Id)
		return err
	}
	q.Logger.Info("Deleted question with id: %v", question.Id)

	return nil
}

func (q *QuestionDAO) GetById(question *models.Question) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from questions where id=$1",
		question.Id)

	err := row.Scan(&question.Id, &question.Test_id, &question.Text)

	if err != nil {
		q.Logger.Error("Unable to get question with id: %v.", question.Id)
		return err
	}

	return nil
}
