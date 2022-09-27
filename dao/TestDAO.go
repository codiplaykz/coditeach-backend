package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type TestDAO struct {
	Logger *logmatic.Logger
}

func (t *TestDAO) Create(test *models.Test) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into tests (name, description, duration, created_at, teacher_id) VALUES($1,$2,$3,$4,$5) returning id",
		test.Name,
		test.Description,
		test.Duration,
		test.Created_at,
		test.Teacher_id).Scan(&test.Id)

	if err != nil {
		t.Logger.Error("Unable to create test.")
		return err
	}

	t.Logger.Info("Test created.")

	return nil
}

func (t *TestDAO) Update(test *models.Test) error {
	err := database.DB.QueryRow(context.Background(),
		"update tests set name=$1, description=$2, duration=$3, created_at=$4, teacher_id=$5 where id=$6 returning id",
		test.Name,
		test.Description,
		test.Duration,
		test.Created_at,
		test.Teacher_id,
		test.Id).Scan(&test.Id)

	if err == pgx.ErrNoRows {
		t.Logger.Error("Test not found")
		return err
	}

	if err != nil {
		t.Logger.Error("Unable to update test.")
		return err
	}

	t.Logger.Info("Updated test with id: %v", test.Id)

	return nil
}

func (t *TestDAO) Delete(test *models.Test) error {
	err := database.DB.QueryRow(context.Background(), "delete from tests where id=$1 returning id", test.Id).Scan(&test.Id)

	if err == pgx.ErrNoRows {
		t.Logger.Error("Test not found")
		return err
	}

	if err != nil {
		t.Logger.Error("Unable to delete test with id: %v", test.Id)
		return err
	}
	t.Logger.Info("Deleted test with id: %v", test.Id)

	return nil
}

func (t *TestDAO) GetById(test *models.Test) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from tests where id=$1",
		test.Id)

	err := row.Scan(&test.Id, &test.Name, &test.Description, &test.Duration, &test.Created_at, &test.Teacher_id)

	if err != nil {
		t.Logger.Error("Unable to get test with id: %v.", test.Id)
		return err
	}

	return nil
}
