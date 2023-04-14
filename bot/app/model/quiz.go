package model

import "time"

type Quiz struct {
	ID        int64
	Topic     string
	Questions []Question
}

type QuizNew struct {
	ID         int64
	Topic      string
	Creator_id int64
	Created_at time.Time
}
