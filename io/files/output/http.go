package main

import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http.ResponseWriter Sample"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
