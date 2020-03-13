package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

func indexHandler(c echo.Context) error {
	db := connect()
	defer db.Close()

	var todos []Todo
	db.Find(&todos)
	return c.JSON(http.StatusOK, todos)
}

func newHandler(c echo.Context) error {
	db := connect()
	defer db.Close()

	newTodo := new(Todo)
	if err := c.Bind(newTodo); err != nil {
		return err
	}
	db.NewRecord(newTodo)
	db.Create(&newTodo)
	return c.JSON(http.StatusOK, newTodo)
}

func toggleHandler(c echo.Context) error {
	db := connect()
	defer db.Close()

	id, _ := strconv.Atoi(c.Param("id"))
	todo := Todo{ID: uint(id)}
	db.First(&todo)
	todo.Done = !todo.Done
	db.Save(&todo)
	return c.NoContent(http.StatusOK)
}

func deleteHandler(c echo.Context) error {
	db := connect()
	defer db.Close()

	id, _ := strconv.Atoi(c.Param(("id")))
	deleteTodo := Todo{ID: uint(id)}
	db.Delete(&deleteTodo)
	return c.NoContent(http.StatusOK)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "PUT", "POST", "DELETE"},
	}))

	// Routing
	e.GET("/todos", indexHandler)
	e.POST("/todos/new", newHandler)
	e.PUT("/todos/:id/toggle", toggleHandler)
	e.DELETE("/todos/:id", deleteHandler)
	e.Logger.Fatal(e.Start(":8080"))
}
