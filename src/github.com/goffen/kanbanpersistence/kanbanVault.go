package kanbanpersistence

import (
	"fmt"
	"io"
	"os"
)

func Persist(kBoard *KanbanBoard) {
	Write()
	kBoard.Name

	filename := "test.txt"

	fmt.Println("writing: " + filename)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	n, err := io.WriteString(f, "blahblahblah")
	if err != nil {
		fmt.Println(n, err)
	}
	f.Close()
}
