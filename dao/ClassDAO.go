package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type ClassDAO struct {
	Logger *logmatic.Logger
}

func (c *ClassDAO) Create(class *models.Class) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into classes (teacher_id, school_id, name, code) VALUES($1,$2,$3,$4) returning id",
		class.Teacher_id,
		class.School_id,
		class.Name,
		class.Code).Scan(&class.Id)

	if err != nil {
		c.Logger.Error("Unable to create class.")
		return err
	}

	c.Logger.Info("Class created.")

	return nil
}

func (c *ClassDAO) Update(class *models.Class) error {
	err := database.DB.QueryRow(context.Background(),
		"update classes set teacher_id=$1, school_id=$2, name=$3, code=$4 where id=$5 returning id",
		class.Teacher_id,
		class.School_id,
		class.Name,
		class.Code,
		class.Id).Scan(&class.Id)

	if err == pgx.ErrNoRows {
		c.Logger.Error("Class not found")
		return err
	}

	if err != nil {
		c.Logger.Error("Unable to update class.")
		return err
	}

	c.Logger.Info("Updated class with id: %v", class.Id)

	return nil
}

func (c *ClassDAO) Delete(class *models.Class) error {
	err := database.DB.QueryRow(context.Background(), "delete from classes where id=$1 returning id", class.Id).Scan(&class.Id)

	if err == pgx.ErrNoRows {
		c.Logger.Error("Class not found")
		return err
	}

	if err != nil {
		c.Logger.Error("Unable to delete class with id: %v", class.Id)
		return err
	}
	c.Logger.Info("Deleted class with id: %v", class.Id)

	return nil
}

func (c *ClassDAO) GetById(class *models.Class) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from classes where id=$1",
		class.Id)

	err := row.Scan(&class.Id, &class.Teacher_id, &class.School_id, &class.Name, &class.Code)

	if err != nil {
		c.Logger.Error("Unable to get class with id: %v.", class.Id)
		return err
	}

	return nil
}

func (c *ClassDAO) GetByParam(paramName, param string, class *models.Class) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from classes where "+paramName+"=$1",
		param)

	err := row.Scan(&class.Id, &class.Teacher_id, &class.School_id, &class.Name, &class.Code)

	if err != nil {
		c.Logger.Error("Unable to get class with %s: %s.", paramName, param)
		return err
	}

	return nil
}
