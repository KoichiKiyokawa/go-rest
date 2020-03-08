package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

type Todo struct {
  ID    uint    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
  Done bool `json:"done"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
  db.LogMode(true)
	defer db.Close()

	db.AutoMigrate(&Todo{})

	var todos []Todo
	db.Find(&todos)
	fmt.Printf("%v\n", todos)

	w.Header().Set("Content-Type", "application/json")
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
