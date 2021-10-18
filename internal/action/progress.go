package action

import (
	"time"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
)

type InProgress interface {
	Get() []model.TaskInfo
	Put(i int, task *model.TaskInfo)
	Remove(i int)
}

type inProgress struct {
	InProgressChan chan *model.TaskInfo
	inProgressList []*model.TaskInfo
}

func NewInProgress(concurrencyLimit int) *inProgress {
	return &inProgress{
		InProgressChan: make(chan *model.TaskInfo),
		inProgressList: make([]*model.TaskInfo, concurrencyLimit),
	}
}

func (p inProgress) Get() []model.TaskInfo {
	var result []model.TaskInfo
	for _, v := range p.inProgressList {
		if v != nil {
			result = append(result, *v)
		}
	}
	return result
}

func (p *inProgress) Put(i int, task *model.TaskInfo) {
	now := time.Now()
	task.StartedAt = &now
	task.Status = model.InProgress

	p.inProgressList[i] = task
}

func (p *inProgress) Remove(i int) {
	p.inProgressList[i] = nil
}
