package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Todo struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}

func connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(true)

	db.AutoMigrate(&Todo{})
	return db
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	db := connect()
  defer db.Close()

	var todos []Todo
	db.Find(&todos)

	w.Header().Set("Content-Type", "application/json")
  resp, _ := json.Marshal(todos)
  w.Write(resp)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	db := connect()
  defer db.Close()

	title := r.URL.Query().Get("title")
	body := r.URL.Query().Get("body")
	newTodo := Todo{Title: title, Body: body, Done: false}
  db.NewRecord(newTodo)
	db.Create(&newTodo)
}

func main() {
  http.HandleFunc("/new",newHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
