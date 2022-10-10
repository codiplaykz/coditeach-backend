package dao

import (
	"coditeach/database"
	"coditeach/helpers"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type SchoolDAO struct {
	Logger *logmatic.Logger
}

func (c *SchoolDAO) Create(school *models.School) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into schools (name, location) VALUES($1,$2) returning id",
		school.Name,
		school.Location).Scan(&school.Id)

	if err != nil {
		c.Logger.Error("Unable to create school.")
		return err
	}

	c.Logger.Info("School created with name: %s, location: %s", school.Name, school.Location)

	return nil
}

func (c *SchoolDAO) Update(school *models.School) error {
	err := database.DB.QueryRow(context.Background(),
		"update schools set name=$1, location=$2 where id=$3 returning id",
		school.Name,
		school.Location,
		school.Id).Scan(&school.Id)

	if err == pgx.ErrNoRows {
		c.Logger.Error("School not found")
		return err
	}

	if err != nil {
		c.Logger.Error("Unable to update school.")
		return err
	}

	c.Logger.Info("School updated with id: %v", school.Id)

	return nil
}

func (c *SchoolDAO) Delete(school *models.School) error {
	err := database.DB.QueryRow(context.Background(), "delete from schools where id=$1 returning id", school.Id).Scan(&school.Id)

	if err == pgx.ErrNoRows {
		c.Logger.Error("School not found")
		return err
	}

	if err != nil {
		c.Logger.Error("Unable to delete school with id: %v", school.Id)
		return err
	}

	c.Logger.Info("Deleted school with id: %v", school.Id)

	return nil
}

func (c *SchoolDAO) GetById(school *models.School) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from schools where id=$1",
		school.Id)

	err := row.Scan(&school.Id, &school.Name, &school.Location)

	if err == pgx.ErrNoRows {
		c.Logger.Error("School not found")
		return err
	}

	if err != nil {
		c.Logger.Error("Unable to get school with id: %v.", school.Id)
		return err
	}

	return nil
}

func (c *SchoolDAO) GetAll() ([]map[string]interface{}, error) {
	rows, err := database.DB.Query(context.Background(),
		"select * from schools")

	if err != nil {
		c.Logger.Error("Could not get schools")
		return nil, err
	}

	json := helpers.PgSqlRowsToJson(rows)

	if err == pgx.ErrNoRows {
		c.Logger.Error("Schools not found")
		return nil, err
	}

	if err != nil {
		c.Logger.Error("Unable to get schools")
		return nil, err
	}

	return json, nil
}
