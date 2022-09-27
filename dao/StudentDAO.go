package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type StudentDAO struct {
	Logger *logmatic.Logger
}

type StudentAccount struct {
	Id       uint   `json:"id"`
	Login    string `json:"login"`
	Role     string `json:"role"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Class_id uint   `json:"Class_id"`
}

func (s *StudentDAO) Create(student *models.Student) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into students (user_id, class_id) VALUES($1,$2) returning id",
		student.User_id,
		student.Class_id).Scan(&student.Id)

	if err != nil {
		s.Logger.Error("Unable to create student.")
		return err
	}

	s.Logger.Info("Student created.")

	return nil
}

func (s *StudentDAO) Update(student *models.Student) error {
	err := database.DB.QueryRow(context.Background(),
		"update students set user_id=$1, class_id=$2 where id=$3 returning id",
		student.User_id,
		student.Class_id,
		student.Id).Scan(&student.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Student not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to update student.")
		return err
	}

	s.Logger.Info("Updated student with id: %v", student.Id)

	return nil
}

func (s *StudentDAO) Delete(student *models.Student) error {
	err := database.DB.QueryRow(context.Background(), "delete from students where id=$1 returning id", student.Id).Scan(&student.Id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Student not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to delete student with id: %v", student.Id)
		return err
	}

	s.Logger.Info("Deleted student with id: %v", student.Id)

	return nil
}

func (s *StudentDAO) GetById(student *models.Student) error {
	err := database.DB.QueryRow(context.Background(),
		"select * from students where id=$1",
		student.Id).Scan(&student.Id, &student.User_id, &student.Class_id)

	if err == pgx.ErrNoRows {
		s.Logger.Error("Student not found")
		return err
	}

	if err != nil {
		s.Logger.Error("Unable to get student with id: %v.", student.Id)
		return err
	}

	return nil
}

func (s *StudentDAO) CreateAccount(account *StudentAccount) error {
	err := database.DB.QueryRow(context.Background(),
		"select id from users where email=$1",
		account.Email).Scan(&account.Id)

	if err != nil && err != pgx.ErrNoRows {
		s.Logger.Error("Unable to create student account.")
		return err
	}

	if account.Id != 0 {
		s.Logger.Error("Unable to create teacher account, email already exist")
		return errors.New("email already exists")
	}

	err = database.DB.QueryRow(context.Background(),
		"insert into users (login, role_id, password, name, surname, email) VALUES($1,$2,$3,$4,$5,$6) returning id",
		account.Login, 1, account.Password, account.Name, account.Surname, account.Email).Scan(&account.Id)

	if err != nil {
		s.Logger.Error("Unable to create student account.")
		return err
	}

	student := models.Student{
		User_id:  account.Id,
		Class_id: account.Class_id,
	}

	err = s.Create(&student)

	if err != nil {
		s.Logger.Error("Unable to create student account.")
		return err
	}

	s.Logger.Info("User created with login: %s, email: %s", account.Login, account.Email)

	return nil
}
