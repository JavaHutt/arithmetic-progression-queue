package model

import "time"

type TaskStatus string

const (
	InProgress TaskStatus = "in_progress"
	Waiting    TaskStatus = "waiting"
	Finished   TaskStatus = "finished"
)

type Task struct {
	ID       string  `json:"id"`
	Count    uint    `json:"count" validate:"required,gte=0"`
	Delta    float32 `json:"delta" validate:"required"`
	First    float32 `json:"first" validate:"required"`
	Interval float32 `json:"interval" validate:"required,gte=0"`
	TTL      float32 `json:"ttl" validate:"required,gte=0"`
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
