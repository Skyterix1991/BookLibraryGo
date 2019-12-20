package main

import (
	"./controller"
	"./model"
	"./model/repository"
)

func main() {
	initializeTestingValues()
	controller.InitializeRestController()
}

func initializeTestingValues() {
	repository.AddBookToLibrary(model.NewBook(0, "The Hobbit"))
	repository.AddBookToLibrary(model.NewBook(1, "The Fellowship of the Ring"))
	repository.AddBookToLibrary(model.NewBook(2, "The Two Towers"))
	repository.AddBookToLibrary(model.NewBook(3, "The Return of the King"))

	repository.AddUser(model.NewUser(0))
	repository.AddUser(model.NewUser(1))
	repository.AddUser(model.NewUser(2))
}
