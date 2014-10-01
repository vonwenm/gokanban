package kanbanlib

import (
	"fmt"
)

type Kanban struct {
	Columns []KanbanColumn
}

func NewKanban() *Kanban {
	return &Kanban{Columns: make([]KanbanColumn, 0)}
}

func (k *Kanban) AddColumn(name string, maxWip int) KanbanColumn {
	col := NewColumn(name, maxWip)
	k.Columns = append(k.Columns, *col)
	fmt.Printf("Appended to Columns. Now %d long\n", len(k.Columns))
	return *col
}

// Adds task to column by index
func (k *Kanban) AddTask(col int, name string) {
	task := NewTask(name)
	k.Columns[col].addTask(task)
}

func (k *Kanban) MoveTask(from int, to int, task *Task) {
	k.Columns[from].removeTask(task.name)
	k.Columns[to].addTask(task)
}

// implements Stringer()
func (k Kanban) String() string {

	s := fmt.Sprintf("I got %d Columns\n", len(k.Columns))
	for i := 0; i < len(k.Columns); i++ {
		col := k.Columns[i]

		s += fmt.Sprintf("%v\n", col)
	}
	return s
}
