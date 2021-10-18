package action

import (
	"time"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/helpers"
	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
	"github.com/sirupsen/logrus"
)

type FinishedList interface {
	Insert(task *model.TaskInfo)
	Delete(ID string)
	GetTasks() []model.TaskInfo
}

type node struct {
	value *model.TaskInfo
	next  *node
}

type finishedList struct {
	log    logrus.Logger
	head   *node
	length int
}

func NewFinishedList(log logrus.Logger) FinishedList {
	return &finishedList{
		log: log,
	}
}

func (list finishedList) GetTasks() []model.TaskInfo {
	var result []model.TaskInfo

	current := list.head

	if current == nil {
		return result
	}

	for current.next != nil {
		result = append(result, *current.value)

		current = current.next
	}
	return result
}

func (list *finishedList) Insert(task *model.TaskInfo) {
	now := time.Now()
	task.FilnishedAt = &now
	task.Status = model.Finished

	current := list.head
	newNode := &node{task, current}
	list.head = newNode
	list.length++

	go list.deleteAfterTTL(newNode)
}

func (list *finishedList) Delete(ID string) {
	previous := list.head
	if previous.value.ID == ID {
		list.log.Infof("%s time came and it is gone forever", ID)
		list.head = previous.next
		list.length--
		return
	}
	current := previous.next

	for current.value.ID != ID {
		if current.next == nil {
			return
		}

		previous = current
		current = current.next
	}
	list.log.Infof("%s time came and it is gone forever", ID)
	previous.next = current.next
	list.length--
}

func (list *finishedList) deleteAfterTTL(n *node) {
	timeout := helpers.GetTimeDuration(n.value.TTL)
	<-time.After(timeout)
	list.Delete(n.value.ID)
}
