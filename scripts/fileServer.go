package main

import (
	"log"
	"net/http"
)

func HandleMe(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]

	_, err := http.Dir("/home/chirag/hold/code/go-play/scripts").Open(filePath)
	if err != nil {
		http.NotFound(w, r)
	} else {
		http.ServeFile(w, r, filePath)
	}
}

func main() {
	PORT := ":8088"
	http.HandleFunc("/static/", HandleMe)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
