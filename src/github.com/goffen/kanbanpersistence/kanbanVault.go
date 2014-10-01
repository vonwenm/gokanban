package kanbanpersistence

import (
	"encoding/json"
	"fmt"
	l "github.com/goffen/kanbanlib"
	"io"
	"os"
)

func Persist(kBoard *l.KanbanBoard) {

	jsonbytes, marshalError := json.Marshal(kBoard)

	if marshalError != nil {
		fmt.Println(marshalError)
	}

	filename := "test.txt"

	fmt.Println("writing: " + filename)
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}

	n, err := io.WriteString(f, string(jsonbytes))
	if err != nil {
		fmt.Println(n, err)
	}
	f.Close()

	fmt.Printf("Successfully wrote %s to file\n", f.Name())
}
