package action

import (
	"sync"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
)

type WaitingQueue interface {
	Enqueue(task *model.TaskInfo)
	Dequeue() *model.TaskInfo
	GetTasks() []model.TaskInfo
}

type waitingQueue struct {
	sync.Mutex
	tasks []*model.TaskInfo
}

func NewWaitingQueue() WaitingQueue {
	return &waitingQueue{}
}

func (q *waitingQueue) Enqueue(task *model.TaskInfo) {
	q.Lock()
	defer q.Unlock()
	task.Status = model.Waiting
	task.QueueNumber = len(q.tasks) + 1
	q.tasks = append(q.tasks, task)
}

func (q *waitingQueue) Dequeue() *model.TaskInfo {
	q.Lock()
	defer q.Unlock()
	task := q.tasks[0]
	task.QueueNumber = 0
	q.tasks = q.tasks[1:]
	for i := range q.tasks {
		q.tasks[i].QueueNumber--
	}

	return task
}

func (q *waitingQueue) GetTasks() []model.TaskInfo {
	var result []model.TaskInfo
	for _, v := range q.tasks {
		result = append(result, *v)
	}

	return result
}
