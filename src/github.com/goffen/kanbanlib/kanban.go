package kanbanlib

import (
	"fmt"
)

type Kanban struct {
	columns []KanbanColumn
}

func NewKanban() *Kanban {
	return &Kanban{columns: make([]KanbanColumn, 0)}
}

func (k *Kanban) AddColumn(name string, maxWip int) KanbanColumn {
	col := NewColumn(name, maxWip)
	k.columns = append(k.columns, *col)
	fmt.Printf("Appended to columns. Now %d long\n", len(k.columns))
	return *col
}

// Adds task to column by index
func (k *Kanban) AddTask(col int, name string) {
	task := NewTask(name)
	k.columns[col].addTask(task)
}

func (k *Kanban) MoveTask(from int, to int, task *Task) {
	k.columns[from].removeTask(task.name)
	k.columns[to].addTask(task)
}

// implements Stringer()
func (k Kanban) String() string {

	s := fmt.Sprintf("I got %d columns\n", len(k.columns))
	for i := 0; i < len(k.columns); i++ {
		col := k.columns[i]

		s += fmt.Sprintf("%v\n", col)
	}
	return s
}
