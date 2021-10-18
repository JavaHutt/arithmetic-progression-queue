package action

import "github.com/JavaHutt/arithmetic-progression-queue/internal/model"

type WaitingQueue interface {
	Enqueue(task model.TaskInfo)
	Dequeue() model.TaskInfo
	GetTasks() []model.TaskInfo
}

type waitingQueue []model.TaskInfo

func NewWaitingQueue() WaitingQueue {
	return &waitingQueue{}
}

func (q *waitingQueue) Enqueue(task model.TaskInfo) {
	*q = append(*q, task)
}

func (q *waitingQueue) Dequeue() model.TaskInfo {
	task := (*q)[0]
	*q = (*q)[1:]
	return task
}

func (q *waitingQueue) GetTasks() []model.TaskInfo {
	var result []model.TaskInfo
	for _, v := range *q {
		result = append(result, v)
	}

	return result
}
