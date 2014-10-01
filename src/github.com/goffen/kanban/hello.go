package main

import (
	"fmt"
	l "github.com/goffen/kanbanlib"
)

func main() {
	kanban := l.NewKanban()
	kanban.AddColumn("Col 1", 2)
	kanban.AddColumn("Col 2", 2)
	kanban.AddColumn("Col 3", 1)

	kanban.AddTask(2, "Städa inomhus")
	fmt.Printf("%v\n", kanban)
	//fmt.Print(colAdded.String())

	l.Start(kanban)
}
