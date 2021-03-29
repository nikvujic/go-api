package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

type bookHandlers struct {
	store map[string]Book
}

func (h *bookHandlers) books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	case "POST":
		h.post(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func (h *bookHandlers) get(w http.ResponseWriter, r *http.Request) {
	books := make([]Book, len(h.store))

	i := 0
	for _, book := range h.store {
		books[i] = book
		i++
	}

	jsonBytes, err := json.Marshal(books)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h *bookHandlers) post(w http.ResponseWriter, r *http.Request) {
	// 16

}

func newBookHandlers() *bookHandlers {
	return &bookHandlers{
		store: map[string]Book{
			"1": Book{
				ID:     "1",
				Title:  "The Hitchhiker's Guide to the Galaxy",
				Author: "Daglas Adams",
				Price:  40,
			},
			"2": Book{
				ID:     "2",
				Title:  "To Kill a Mockingbird",
				Author: "Harper Lee",
				Price:  20,
			},
			"3": Book{
				ID:     "3",
				Title:  "Don Quixote",
				Author: "Miguel de Cervantes",
				Price:  35,
			},
		},
	}
}

func main() {
	bookHandlers := newBookHandlers()

	http.HandleFunc("/books", bookHandlers.books)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
