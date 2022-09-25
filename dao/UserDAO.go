package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type UserDAO struct {
	Logger *logmatic.Logger
}

func (u *UserDAO) Create(user *models.User) error {
	roleId := 1

	if user.Role == "student" {
		roleId = 1
	} else if user.Role == "teacher" {
		roleId = 2
	} else if user.Role == "parent" {
		roleId = 3
	} else if user.Role == "school_admin" {
		roleId = 4
	} else if user.Role == "admin" {
		roleId = 5
	} else {
		u.Logger.Error("Unable to create user, role not found.")
		return errors.New("role not found")
	}

	row := database.DB.QueryRow(context.Background(),
		"insert into users (login, role_id, password, name, surname, email) VALUES($1,$2,$3,$4,$5,$6) returning id",
		user.Login, roleId, string(user.Password), user.Name, user.Surname, user.Email)

	err := row.Scan(&user.Id)

	if err != nil {
		u.Logger.Error("Unable to create user.")
		return err
	}

	u.Logger.Info("User created with login: %s, email: %s", user.Login, user.Email)

	return nil
}

func (u *UserDAO) GetById(user *models.User) error {
	row := database.DB.QueryRow(context.Background(),
		"select users.id, users.login, roles.name as role, users.password, users.name, users.surname, users.email from users inner join roles on roles.id = users.role_id where users.id = $1", user.Id)

	err := row.Scan(&user.Id, &user.Login, &user.Role, &user.Password, &user.Name, &user.Surname, &user.Email)

	if err == pgx.ErrNoRows {
		u.Logger.Warn("User not found.")
		return nil
	}

	if err != nil {
		u.Logger.Error("Unable to get user by id")
		return err
	}

	return nil
}

func (u *UserDAO) GetByEmail(user *models.User) error {
	row := database.DB.QueryRow(context.Background(),
		"select users.id, users.login, roles.name as role, users.password, users.name, users.surname, users.email from users inner join roles on roles.id = users.role_id where users.email = $1", user.Email)

	err := row.Scan(&user.Id, &user.Login, &user.Role, &user.Password, &user.Name, &user.Surname, &user.Email)

	if err == pgx.ErrNoRows {
		u.Logger.Warn("User not found.")
		return nil
	}

	if err != nil {
		u.Logger.Error("Unable to get user by email.")
		return err
	}

	return nil
}
