package models

import "time"

type TestResult struct {
	Id                uint      `json:"id"`
	Test_id           uint      `json:"test_id"`
	Student_id        uint      `json:"student_id"`
	Incorrect_answers int       `json:"incorrect_answers"`
	Correct_answers   int       `json:"correct_answers"`
	Time_spent        int       `json:"time_spent"`
	Pass_date         time.Time `json:"pass_date"`
}
