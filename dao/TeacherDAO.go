package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type TeacherDAO struct {
	Logger *logmatic.Logger
}

type TeacherAccount struct {
	Id        uint   `json:"id"`
	Login     string `json:"login"`
	Role      string `json:"role"`
	Password  string `json:"-"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	School_id uint   `json:"school_id"`
}

func (t *TeacherDAO) Create(teacher *models.Teacher) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into teachers (user_id, school_id) VALUES($1,$2) returning id",
		teacher.User_id,
		teacher.School_id).Scan(&teacher.Id)

	if err != nil {
		t.Logger.Error("Unable to create teacher.")
		return err
	}

	t.Logger.Info("Teacher created.")

	return nil
}

func (t *TeacherDAO) Update(teacher *models.Teacher) error {
	err := database.DB.QueryRow(context.Background(),
		"update teachers set user_id=$1, school_id=$2 where id=$3 returning id",
		teacher.User_id,
		teacher.School_id,
		teacher.Id).Scan(&teacher.Id)

	if err == pgx.ErrNoRows {
		t.Logger.Error("Teacher not found")
		return err
	}

	if err != nil {
		t.Logger.Error("Unable to update teacher.")
		return err
	}

	t.Logger.Info("Updated teacher with id: %v", teacher.Id)

	return nil
}

func (t *TeacherDAO) Delete(teacher *models.Teacher) error {
	err := database.DB.QueryRow(context.Background(), "delete from teachers where id=$1 returning id", teacher.Id).Scan(&teacher.Id)

	if err == pgx.ErrNoRows {
		t.Logger.Error("Teacher not found")
		return err
	}

	if err != nil {
		t.Logger.Error("Unable to delete teacher with id: %v", teacher.Id)
		return err
	}

	t.Logger.Info("Deleted teacher with id: %v", teacher.Id)

	return nil
}

func (t *TeacherDAO) GetById(teacher *models.Teacher) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from teachers where id=$1",
		teacher.Id)

	err := row.Scan(&teacher.Id, &teacher.User_id, &teacher.School_id)

	if err == pgx.ErrNoRows {
		t.Logger.Error("Teacher not found")
		return err
	}

	if err != nil {
		t.Logger.Error("Unable to get teacher with id: %v.", teacher.Id)
		return err
	}

	return nil
}

func (t *TeacherDAO) CreateAccount(account *TeacherAccount) error {
	err := database.DB.QueryRow(context.Background(),
		"select id from users where email=$1",
		account.Email).Scan(&account.Id)

	if err != nil && err != pgx.ErrNoRows {
		t.Logger.Error("Unable to create teacher account.")
		return err
	}

	if account.Id != 0 {
		t.Logger.Error("Unable to create teacher account, email already exist")
		return errors.New("email already exists")
	}

	err = database.DB.QueryRow(context.Background(),
		"insert into users (login, role_id, password, name, surname, email) VALUES($1,$2,$3,$4,$5,$6) returning id",
		account.Login, 2, account.Password, account.Name, account.Surname, account.Email).Scan(&account.Id)

	if err != nil {
		t.Logger.Error("Unable to create teacher account.")
		return err
	}

	teacher := models.Teacher{
		User_id:   account.Id,
		School_id: account.School_id,
	}

	err = t.Create(&teacher)
	if err != nil {
		t.Logger.Error("Unable to create teacher account.")
		return err
	}

	t.Logger.Info("User created with login: %s, email: %s", account.Login, account.Email)

	return nil
}
