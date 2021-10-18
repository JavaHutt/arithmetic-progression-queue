package action

import (
	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/sirupsen/logrus"
)

type ArithmeticProcessor interface {
	StartWorkers()
}

type arithmeticProcessor struct {
	log              logrus.Logger
	concurrencyLimit int
	waitingQueue     WaitingQueue
	inProgressChan   chan model.TaskInfo
	inProgressList   []*model.TaskInfo
	finishedList     FinishedList
}

func NewArithmeticProcessor(log logrus.Logger, concurrencyLimit int) ArithmeticProcessor {
	return &arithmeticProcessor{
		log:              log,
		concurrencyLimit: concurrencyLimit,
		waitingQueue:     NewWaitingQueue(),
		inProgressList:   make([]*model.TaskInfo, 0, concurrencyLimit),
		finishedList:     NewFinishedList(),
	}
}

func (a arithmeticProcessor) StartWorkers() {
	for i := 0; i < a.concurrencyLimit; i++ {
		go a.worker(i)
	}
}

func (a arithmeticProcessor) worker(i int) {
	a.log.Infof("#%d worker has started", i)
	for task := range a.inProgressChan {
		a.inProgressList[i] = &task
		a.log.Infof("task %s is now in progress, handled by worker #%d", task.ID, i)
		// TODO implement process method
		a.inProgressList[i] = nil
		a.finishedList.Insert(task)
		a.log.Infof("#%d worker has done with task %s", i, task.ID)
	}
}
