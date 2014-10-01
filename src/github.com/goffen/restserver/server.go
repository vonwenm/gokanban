package restserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	l "github.com/goffen/kanbanlib"
)

func Start(kanban *l.KanbanBoard) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(kanban)
		if err != nil {
			fmt.Println("error:", err)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Write(b)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
