package models

import "time"

type TaskStatus string

const (
	InProgress TaskStatus = "in_progress"
	Waiting    TaskStatus = "waiting"
	Finished   TaskStatus = "finished"
)

type Task struct {
	Count    uint    `json:"n"`
	Delta    float32 `json:"d"`
	First    float32 `json:"n1"`
	Interval float32 `json:"I"`
	TTL      float32 `json:"TTL"`
}

type TaskInfo struct {
	Task
	QueueNumber      int        `json:"queue_number"`
	Status           TaskStatus `json:"status"`
	CurrentIteration int        `json:"current_iteration"`
	CreatedAt        time.Time  `json:"created_at"`
	StartedAt        time.Time  `json:"started_at"`
	FilnishedAt      time.Time  `json:"finished_at"`
}
