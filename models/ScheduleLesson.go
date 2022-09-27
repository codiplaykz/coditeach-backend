package models

import (
	"time"
)

type ScheduleLesson struct {
	Id          uint      `json:"id"`
	Schedule_id uint      `json:"schedule_id"`
	Start_time  time.Time `json:"start_time"`
	End_time    time.Time `json:"end_time"`
}
