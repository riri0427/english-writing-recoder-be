package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/sentences", getAllSentences)
	http.HandleFunc("/sentence", sentence)
	http.ListenAndServe(":8080", nil)
}

func getAllSentences(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It's cloudy today.")
}

func sentence(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getSentence(w, r)
	} else if r.Method == http.MethodPost {
		createSentence(w, r)
	}
}

func getSentence(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It's sunny today.")
}

func createSentence(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}
