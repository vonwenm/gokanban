package kanbanlib

import (
	"fmt"
	"time"
)

type Task struct {
	name    string
	created time.Time
}

func (t Task) String() string {
	return fmt.Sprintf("%v (created %v)", t.name, t.created)
}
func NewTask(name string) *Task {
	return &Task{name: name, created: time.Now()}
}
