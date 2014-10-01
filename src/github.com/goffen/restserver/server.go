package restserver

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	l "github.com/goffen/kanbanlib"
)

func serveJson(k interface{}, w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(k)
	if err != nil {
		fmt.Println("error:", err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func Start(kanban *l.KanbanBoard) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveJson(kanban, w, r)
	})

	http.HandleFunc("/nrBoards", func(w http.ResponseWriter, r *http.Request) {
		serveJson(kanban.NrColumns(), w, r)
	})
	boardSegment := "/board/"
	http.HandleFunc(boardSegment, func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len(boardSegment):]
		splits := strings.Split(path, "/")
		columnId, err := strconv.Atoi(splits[0])
		if err != nil {
			fmt.Println("Inte så bra grej")
		}
		column := kanban.GetColumn(columnId)
		serveJson(column, w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
