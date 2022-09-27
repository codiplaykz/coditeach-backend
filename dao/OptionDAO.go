package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type OptionDAO struct {
	Logger *logmatic.Logger
}

func (o *OptionDAO) Create(option *models.Option) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into options (question_id, text, is_correct) VALUES($1,$2,$3) returning id",
		option.Question_id,
		option.Text,
		option.Is_correct).Scan(&option.Id)

	if err != nil {
		o.Logger.Error("Unable to create option.")
		return err
	}

	o.Logger.Info("Option created.")

	return nil
}

func (o *OptionDAO) Update(option *models.Option) error {
	err := database.DB.QueryRow(context.Background(),
		"update options set question_id=$1, text=$2, is_correct=$3 where id=$4 returning id",
		option.Question_id,
		option.Text,
		option.Is_correct,
		option.Id).Scan(&option.Id)

	if err == pgx.ErrNoRows {
		o.Logger.Error("Option not found")
		return err
	}

	if err != nil {
		o.Logger.Error("Unable to update option.")
		return err
	}

	o.Logger.Info("Updated class with id: %v", option.Id)

	return nil
}

func (o *OptionDAO) Delete(option *models.Option) error {
	err := database.DB.QueryRow(context.Background(), "delete from options where id=$1 returning id", option.Id).Scan(&option.Id)

	if err == pgx.ErrNoRows {
		o.Logger.Error("Option not found")
		return err
	}

	if err != nil {
		o.Logger.Error("Unable to delete option with id: %v", option.Id)
		return err
	}
	o.Logger.Info("Deleted option with id: %v", option.Id)

	return nil
}

func (o *OptionDAO) GetById(option *models.Option) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from options where id=$1",
		option.Id)

	err := row.Scan(&option.Id, &option.Question_id, &option.Text, &option.Is_correct)

	if err != nil {
		o.Logger.Error("Unable to get option with id: %v.", option.Id)
		return err
	}

	return nil
}
