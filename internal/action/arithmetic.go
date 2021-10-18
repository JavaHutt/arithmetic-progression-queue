package action

import (
	"fmt"
	"time"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/helpers"
	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/sirupsen/logrus"
)

type ArithmeticProcessor interface {
	StartWorkers()
	AddTask(task *model.TaskInfo)
	GetTasks() []model.TaskInfo
}

type arithmeticProcessor struct {
	log              logrus.Logger
	concurrencyLimit int
	waitingQueue     WaitingQueue
	inProgress       inProgress
	finishedList     FinishedList
}

func NewArithmeticProcessor(log logrus.Logger, concurrencyLimit int) ArithmeticProcessor {
	inProgress := NewInProgress(concurrencyLimit)

	return &arithmeticProcessor{
		log:              log,
		concurrencyLimit: concurrencyLimit,
		waitingQueue:     NewWaitingQueue(inProgress),
		inProgress:       *inProgress,
		finishedList:     NewFinishedList(),
	}
}

func (a arithmeticProcessor) AddTask(task *model.TaskInfo) {
	a.waitingQueue.Enqueue(task)
	a.inProgress.InProgressChan <- task
}

func (a arithmeticProcessor) GetTasks() []model.TaskInfo {
	var result []model.TaskInfo

	result = append(result, a.inProgress.Get()...)
	result = append(result, a.waitingQueue.GetTasks()...)
	result = append(result, a.finishedList.GetTasks()...)

	return result
}

func (a arithmeticProcessor) StartWorkers() {
	for i := 0; i < a.concurrencyLimit; i++ {
		go a.worker(i)
	}
}

func (a *arithmeticProcessor) worker(i int) {
	a.log.Infof("#%d worker has started", i)
	for task := range a.inProgress.InProgressChan {
		a.inProgress.Put(i, task)
		a.waitingQueue.Dequeue()
		a.log.Infof("task %s is now in progress, handled by worker #%d", task.ID, i)
		// TODO implement process method
		a.inProgress.Remove(i)
		a.finishedList.Insert(task)
		a.log.Infof("#%d worker has done with task %s", i, task.ID)
	}
}

func (a arithmeticProcessor) processTask(task *model.TaskInfo) {
	ticker := time.NewTicker(helpers.GetTimeDuration(task.Interval))
	stop := make(chan struct{})

	go func() {
		for range ticker.C {
			if task.Count == 0 {
				stop <- struct{}{}
				break
			}

			task.Count--
		}
	}()

	go func() {
		<-stop
		fmt.Println("stopped")
		ticker.Stop()
	}()
}
