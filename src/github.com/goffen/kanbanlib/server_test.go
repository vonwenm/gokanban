package kanbanlib

import (
	"fmt"
	"testing"
	"net/http"
	"io/ioutil"
)

func TestServe(t *testing.T) {

	go Start()

	resp, err := http.Get("http://localhost:8080")	
	if err != nil {
		t.Error(fmt.Sprintf("Got error %s", err))
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s",body)
}


