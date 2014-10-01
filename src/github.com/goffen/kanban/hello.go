package main

import (
	"fmt"
	l "github.com/goffen/kanbanlib"
	p "github.com/goffen/kanbanpersistence"
	s "github.com/goffen/restserver"
)

func main() {
	kanban := l.NewKanban(1)
	kanban.AddColumn("Col 1", 2, 1)
	kanban.AddColumn("Col 2", 2, 2)
	kanban.AddColumn("Col 3", 1, 3)

	kanban.AddTask(2, "Städa inomhus")
	fmt.Printf("%v\n", kanban)
	//fmt.Print(colAdded.String())

	go s.Start(kanban)
	p.Persist(kanban)
}
