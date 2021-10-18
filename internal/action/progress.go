package action

import (
	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
)

type InProgress interface {
	Put(i int, task *model.TaskInfo)
	Remove(i int)
	Get() <-chan *model.TaskInfo
}

type inProgress struct {
	inProgressChan chan *model.TaskInfo
	inProgressList []*model.TaskInfo
}

func NewInProgress(concurrencyLimit int) InProgress {
	return &inProgress{
		inProgressChan: make(chan *model.TaskInfo),
		inProgressList: make([]*model.TaskInfo, 0, concurrencyLimit),
	}
}

func (p *inProgress) Put(i int, task *model.TaskInfo) {
	task.Status = model.InProgress
	p.inProgressList[i] = task
	p.inProgressChan <- task
}

func (p *inProgress) Remove(i int) {
	p.inProgressList[i] = nil
}

func (p inProgress) Get() <-chan *model.TaskInfo {
	return p.inProgressChan
}
