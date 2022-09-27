package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type TestResultDAO struct {
	Logger *logmatic.Logger
}

func (t *TestResultDAO) Create(testResult *models.TestResult) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into test_results (test_id, student_id, incorrect_answers, correct_answers, time_spent, pass_date) VALUES($1,$2,$3,$4,$5,$6) returning id",
		testResult.Test_id,
		testResult.Student_id,
		testResult.Incorrect_answers,
		testResult.Correct_answers,
		testResult.Time_spent,
		testResult.Pass_date).Scan(&testResult.Id)

	if err != nil {
		t.Logger.Error("Unable to create test result.")
		return err
	}

	t.Logger.Info("Test result created.")

	return nil
}

func (t *TestResultDAO) Update(testResult *models.TestResult) error {
	err := database.DB.QueryRow(context.Background(),
		"update test_results set test_id=$1, student_id=$2, incorrect_answers=$3, correct_answers=$4, time_spent=$5, pass_date=$6 where id=$7 returning id",
		testResult.Test_id,
		testResult.Student_id,
		testResult.Incorrect_answers,
		testResult.Correct_answers,
		testResult.Time_spent,
		testResult.Pass_date,
		testResult.Id).Scan(&testResult.Id)

	if err == pgx.ErrNoRows {
		t.Logger.Error("Test result not found")
		return err
	}

	if err != nil {
		t.Logger.Error("Unable to update test result.")
		return err
	}

	t.Logger.Info("Updated test result with id: %v", testResult.Id)

	return nil
}

func (t *TestResultDAO) Delete(testResult *models.TestResult) error {
	err := database.DB.QueryRow(context.Background(), "delete from test_results where id=$1 returning id", testResult.Id).Scan(&testResult.Id)

	if err == pgx.ErrNoRows {
		t.Logger.Error("Test result not found")
		return err
	}

	if err != nil {
		t.Logger.Error("Unable to delete test result with id: %v", testResult.Id)
		return err
	}
	t.Logger.Info("Deleted test result with id: %v", testResult.Id)

	return nil
}

func (t *TestResultDAO) GetById(testResult *models.TestResult) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from test_results where id=$1",
		testResult.Id)

	err := row.Scan(&testResult.Id,
		&testResult.Test_id,
		&testResult.Student_id,
		&testResult.Incorrect_answers,
		&testResult.Correct_answers,
		&testResult.Time_spent,
		&testResult.Pass_date)

	if err != nil {
		t.Logger.Error("Unable to get test result with id: %v.", testResult.Id)
		return err
	}

	return nil
}
