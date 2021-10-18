package service

import (
	"time"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/sirupsen/logrus"
	"github.com/teris-io/shortid"
)

type arithmeticProcessor interface {
	AddTask(task *model.TaskInfo)
	GetTasks() []model.TaskInfo
}

type TaskService interface {
	AddTask(task model.Task) error
	GetTasks() []model.TaskInfo
}

type service struct {
	log       logrus.Logger
	shortID   *shortid.Shortid
	processor arithmeticProcessor
}

func NewTaskService(log logrus.Logger, processor arithmeticProcessor) TaskService {
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)

	return &service{
		log:       log,
		shortID:   sid,
		processor: processor,
	}
}

func (s service) AddTask(task model.Task) error {
	id, err := s.shortID.Generate()
	if err != nil {
		s.log.WithError(err).Info("unable to generate ID")
		return err
	}

	task.ID = id

	s.log.WithFields(logrus.Fields{
		"count":    task.Count,
		"delta":    task.Delta,
		"first":    task.First,
		"interval": task.Interval,
		"TTL":      task.TTL,
	}).Infof("new task was recieved, new ID given: %s", id)

	taskInfo := model.TaskInfo{
		Task: model.Task{
			ID:       task.ID,
			Count:    task.Count,
			Delta:    task.Delta,
			First:    task.First,
			Interval: task.Interval,
			TTL:      task.TTL,
		},
		QueueNumber:      0,
		Status:           model.Waiting,
		CurrentIteration: 0,
		CreatedAt:        time.Now(),
		StartedAt:        nil,
		FilnishedAt:      nil,
	}

	s.processor.AddTask(&taskInfo)

	return nil
}

func (s service) GetTasks() []model.TaskInfo {
	return s.processor.GetTasks()
}
