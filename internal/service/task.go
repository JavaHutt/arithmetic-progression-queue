package service

import (
	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

type TaskService interface {
	AddTask(task model.Task) error
	GetTasks() ([]model.TaskInfo, error)
}

type service struct {
	log     logrus.Logger
	shortID *shortid.Shortid
}

func NewTaskService(log logrus.Logger) TaskService {
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)

	return &service{
		log:     log,
		shortID: sid,
	}
}

func (s service) AddTask(task model.Task) error {
	id, err := s.shortID.Generate()
	if err != nil {
		return err
	}

	task.ID = id

	s.log.WithFields(logrus.Fields{
		"count":    task.Count,
		"delta":    task.Delta,
		"first":    task.First,
		"interval": task.Interval,
		"TTL":      task.TTL,
	}).Infof("new task was recieved, new id given: %s", id)

	return nil
}

func (s service) GetTasks() ([]model.TaskInfo, error) {
	s.log.Info("get tasks")

	return nil, nil
}
