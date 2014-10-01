package kanbanlib

import (
	//"fmt"
	"testing"
)

func TestKanbanColumn(t *testing.T) {
	col := NewColumn("name", 1)
	if col.taskCount() != 0 {
		t.Error("Har tasks")
	}
	if col.name != "name" {
		t.Error("Fel namn")
	}
	col.addTask(NewTask("task 1"))

	if col.taskCount() != 1 {
		t.Error("Har inte en task")
	}

	col.removeTask("task 1")

	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("work did fail (is ok):", err)
	//	}
	//}()

	col.addTask(NewTask("task 2"))
}
