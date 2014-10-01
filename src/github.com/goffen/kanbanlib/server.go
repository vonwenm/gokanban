package kanbanlib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Start(kanban *KanbanBoard) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(kanban)
		if err != nil {
			fmt.Println("error:", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
