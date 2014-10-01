package kanbanlib

import (
	"fmt"
	"time"
)

type Task struct {
	Name    string
	Created time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("%v (Created %v)", t.Name, t.Created)
}
func NewTask(name string) *Task {
	return &Task{Name: name, Created: time.Now()}
}
