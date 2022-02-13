package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
}


func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}