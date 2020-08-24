package main

import (
	"encoding/json"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

type Message struct{
	Text string `json:"text"`
}

func showMessage(w http.ResponseWriter, r *http.Request){
	message:= Message{"World"}
	output,err:= json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func main() {
	folder:= packr.New("Goapp","./ui/build")
	http.Handle("/", http.FileServer(folder))
	http.HandleFunc("/hello", showMessage)
	go http.ListenAndServe(":8000", nil)
}