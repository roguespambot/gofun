// a webserver than returns some json

package main

import (
	"encoding/json"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	myItems := map[string]string{"item1": "1", "item2": "2", "item3": "3"}
	a, _ := json.Marshal(myItems)

	w.Write(a)
	return
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
