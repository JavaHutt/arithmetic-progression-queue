package action

import (
	"math"
	"time"

	"github.com/JavaHutt/arithmetic-progression-queue/internal/model"
)

type FinishedList interface {
	Insert(task model.TaskInfo)
	Delete(ID string)
}

type node struct {
	value model.TaskInfo
	next  *node
}

type finishedList struct {
	head   *node
	length int
}

func NewFinishedList() FinishedList {
	return &finishedList{}
}

func (list *finishedList) Insert(task model.TaskInfo) {
	current := list.head
	newNode := &node{task, current}
	list.head = newNode
	list.length++
	go list.deleteAfterTTL(newNode)
}

func (list *finishedList) Delete(ID string) {
	previous := list.head
	if previous.value.ID == ID {
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

	previous.next = current.next
	list.length--
}

func (list *finishedList) deleteAfterTTL(n *node) {
	integer, float := math.Modf(float64(n.value.TTL))
	timeout := time.Second*time.Duration(integer) + time.Millisecond*time.Duration(float)
	<-time.After(timeout)
	list.Delete(n.value.ID)
}
