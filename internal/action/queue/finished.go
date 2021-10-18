package queue

import "github.com/JavaHutt/arithmetic-progression-queue/internal/model"

type node struct {
	value model.Task
	next  *node
}

type FinishedList struct {
	head   *node
	length int
}

func (list *FinishedList) Insert(task model.Task) {
	current := list.head
	list.head = &node{task, current}
	list.length++
}

func (list *FinishedList) delete(ID string) {
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
