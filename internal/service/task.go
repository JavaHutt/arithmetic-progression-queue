package service

import (
	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/sirupsen/logrus"
)

type TaskService interface {
	AddTask(task model.Task) error
	GetTasks() ([]model.TaskInfo, error)
}

type service struct {
	log logrus.Logger
}

func NewTaskService(log logrus.Logger) TaskService {
	return &service{
		log: log,
	}
}

func (s service) AddTask(task model.Task) error {
	s.log.WithFields(logrus.Fields{
		"count":    task.Count,
		"delta":    task.Delta,
		"first":    task.First,
		"interval": task.Interval,
		"TTL":      task.TTL,
	}).Info("new task was recieved")

	return nil
}

func (s service) GetTasks() ([]model.TaskInfo, error) {
	s.log.Info("get tasks")

	return nil, nil
}
