package kanbanlib

import (
	"fmt"
)

type KanbanColumn struct {
	Id int
	Name   string
	MaxWip int
	Tasks  []Task
}

func (col KanbanColumn) String() string {

	colDesc := fmt.Sprintf("%v (wip: %d, %#v)", col.Name, col.MaxWip, col.Tasks)

	return colDesc
}

func (col *KanbanColumn) removeTask(taskName string) {
	i := -1
	for index, t := range col.Tasks {
		if t.Name == taskName {
			i = index
			break
		}
	}
	if i == -1 {
		panic(fmt.Sprintf("Could not remove task %v", taskName))
	}

	col.Tasks = append(col.Tasks[:i], col.Tasks[i+1:]...)
}

func (col *KanbanColumn) taskCount() int {
	return len(col.Tasks)
}

func (col *KanbanColumn) addTask(task *Task) {
	if len(col.Tasks) == cap(col.Tasks) {
		panic(fmt.Sprintf("Could not add more task to column!"))
	}
	col.Tasks = col.Tasks[:len(col.Tasks)+1]
	col.Tasks[len(col.Tasks)-1] = *task
}

func NewColumn(name string, maxWip int, id int) *KanbanColumn {
	return &KanbanColumn{Name: name, MaxWip: maxWip, Tasks: make([]Task, 0, maxWip), Id: id}
}

