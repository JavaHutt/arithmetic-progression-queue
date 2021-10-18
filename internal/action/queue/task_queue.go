package queue

type TaskQueue interface{}

type taskQueue struct {
	finishedList FinishedList
}

func NewTaskQueue() TaskQueue {
	return &taskQueue{}
}
