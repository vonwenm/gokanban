package kanbanlib

import (
	"fmt"
)

type KanbanBoard struct {
	Id int
	Columns []KanbanColumn
}

func NewKanban(id int) *KanbanBoard {
	return &KanbanBoard{Columns: make([]KanbanColumn, 0), Id: id }
}

func (k *KanbanBoard) GetColumn(id int) *KanbanColumn {
	for i := 0; i < len(k.Columns); i++ {
		col := k.Columns[i]
		if(col.Id == id) {
			return &col
		}
	}
	return nil
}

func (k *KanbanBoard) NrColumns() int {
	return len(k.Columns)
}

func (k *KanbanBoard) AddColumn(name string, maxWip int, id int) KanbanColumn {
	col := NewColumn(name, maxWip, id)
	k.Columns = append(k.Columns, *col)
	fmt.Printf("Appended to Columns. Now %d long\n", len(k.Columns))
	return *col
}

// Adds task to column by index
func (k *KanbanBoard) AddTask(col int, name string) {
	task := NewTask(name)
	k.Columns[col].addTask(task)
}

func (k *KanbanBoard) MoveTask(from int, to int, task *Task) {
	k.Columns[from].removeTask(task.Name)
	k.Columns[to].addTask(task)
}

// implements Stringer()
func (k KanbanBoard) String() string {

	s := fmt.Sprintf("I got %d Columns\n", len(k.Columns))
	for i := 0; i < len(k.Columns); i++ {
		col := k.Columns[i]

		s += fmt.Sprintf("%v\n", col)
	}
	return s
}
