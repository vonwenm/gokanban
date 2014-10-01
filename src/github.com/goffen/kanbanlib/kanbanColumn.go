package kanbanlib

import (
	"fmt"
)

type KanbanColumn struct {
	name   string
	maxWip int
	tasks  []Task
}

func (col KanbanColumn) String() string {

	colDesc := fmt.Sprintf("%v (wip: %d, %#v)", col.name, col.maxWip, col.tasks)
	
	return colDesc
}

func (col *KanbanColumn) removeTask(taskName string) {
	i := -1
	for index, t := range col.tasks {
		if(t.name == taskName) {
			i = index
			break
		}
	}
	if(i == -1) {
		panic(fmt.Sprintf("Could not remove task %v", taskName))
	}

	col.tasks = append(col.tasks[:i], col.tasks[i+1:]...)
}

func (col *KanbanColumn) taskCount() int {
	return len(col.tasks)
}

func (col *KanbanColumn) addTask(task *Task) {
	if len(col.tasks) == cap(col.tasks) {
		panic(fmt.Sprintf("Could not add more task to column!"))
	}
	col.tasks = col.tasks[:len(col.tasks)+1]
	col.tasks[len(col.tasks)-1] = *task
}

func NewColumn(name string, maxWip int) *KanbanColumn {
	return &KanbanColumn{name: name, maxWip: maxWip, tasks: make([]Task, 0, maxWip)}
}
