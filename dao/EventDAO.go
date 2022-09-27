package dao

import (
	"coditeach/database"
	"coditeach/models"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/mborders/logmatic"
)

type EventDAO struct {
	Logger *logmatic.Logger
}

func (e *EventDAO) Create(event *models.Event) error {
	err := database.DB.QueryRow(context.Background(),
		"insert into events (title, description, date) VALUES($1,$2,$3) returning id",
		event.Title,
		event.Description,
		event.Date).Scan(&event.Id)

	if err != nil {
		e.Logger.Error("Unable to create event.")
		return err
	}

	e.Logger.Info("Event created.")

	return nil
}

func (e *EventDAO) Update(event *models.Event) error {
	err := database.DB.QueryRow(context.Background(),
		"update events set title=$1, description=$2, date=$3 where id=$4 returning id",
		event.Title,
		event.Description,
		event.Date,
		event.Id).Scan(&event.Id)

	if err == pgx.ErrNoRows {
		e.Logger.Error("Event not found")
		return err
	}

	if err != nil {
		e.Logger.Error("Unable to update event.")
		return err
	}

	e.Logger.Info("Updated event with id: %v", event.Id)

	return nil
}

func (e *EventDAO) Delete(event *models.Event) error {
	err := database.DB.QueryRow(context.Background(), "delete from events where id=$1 returning id", event.Id).Scan(&event.Id)

	if err == pgx.ErrNoRows {
		e.Logger.Error("Event not found")
		return err
	}

	if err != nil {
		e.Logger.Error("Unable to delete event with id: %v", event.Id)
		return err
	}
	e.Logger.Info("Deleted event with id: %v", event.Id)

	return nil
}

func (e *EventDAO) GetById(event *models.Event) error {
	row := database.DB.QueryRow(context.Background(),
		"select * from events where id=$1",
		event.Id)

	err := row.Scan(&event.Id, &event.Title, &event.Description, &event.Date)

	if err != nil {
		e.Logger.Error("Unable to get event with id: %v.", event.Id)
		return err
	}

	return nil
}
