package controller

import (
	"../model"
	"../model/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func deleteBookById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	bookId, _ := strconv.ParseInt(params["bookId"], 10, 64)

	if !repository.CheckIfBookExists(int(bookId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)

	repository.DeleteBookFromLibrary(int(bookId))
}

func createBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	decoder := json.NewDecoder(request.Body)

	var partialBook model.Book

	_ = decoder.Decode(&partialBook)

	book := model.NewBook(partialBook.Id, partialBook.Name)
	repository.AddBookToLibrary(book)
}

func getBookById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	bookId, _ := strconv.ParseInt(params["bookId"], 10, 64)

	if !repository.CheckIfBookExists(int(bookId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)

	var book *model.Book

	book = repository.GetBookById(int(bookId))

	jsonResponse := book.MarshalJSON()

	response.Write(jsonResponse)
}

func getMostIssuedBook(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	jsonResponse := repository.GetMostIssuedBook().MarshalJSON()

	response.Write(jsonResponse)
}

func isBookIssued(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	bookId, _ := strconv.ParseInt(params["bookId"], 10, 64)

	if !repository.CheckIfBookExists(int(bookId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	data := make(map[string]interface{})
	data["issued"] = repository.IsBookIssued(int(bookId))

	jsonResponse, _ := json.Marshal(data)

	response.Write(jsonResponse)
}

func issueBookToUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	bookId, _ := strconv.ParseInt(params["bookId"], 10, 64)
	userId, _ := strconv.ParseInt(params["userId"], 10, 64)

	if !repository.CheckIfBookExists(int(bookId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	if !repository.CheckIfUserExists(int(userId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	var book *model.Book
	var user *model.User

	book = repository.GetBookById(int(bookId))
	user = repository.GetUserById(int(userId))

	if repository.IsBookIssued(int(bookId)) {
		response.WriteHeader(http.StatusConflict)
	} else {
		response.WriteHeader(http.StatusOK)

		repository.IssueBook(book, user)
	}
}

func getAllBooks(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	allBooks := repository.GetAllBooksInLibrary()

	jsonResponse, _ := json.Marshal(allBooks)

	response.Write(jsonResponse)
}

func returnBookToLibrary(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	bookId, _ := strconv.ParseInt(params["bookId"], 10, 64)

	if !repository.CheckIfBookExists(int(bookId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	var book *model.Book

	book = repository.GetBookById(int(bookId))

	if repository.IsBookIssued(int(bookId)) {
		response.WriteHeader(http.StatusOK)

		repository.ReturnBook(book)
	} else {
		response.WriteHeader(http.StatusConflict)
	}
}

func getAllUsers(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)

	allBooks := repository.GetAllUsers()

	jsonResponse, _ := json.Marshal(allBooks)

	response.Write(jsonResponse)
}

func createUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)

	decoder := json.NewDecoder(request.Body)

	var partialUser model.User

	_ = decoder.Decode(&partialUser)

	user := model.NewUser(partialUser.Id)
	repository.AddUser(user)
}

func getUserById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	userId, _ := strconv.ParseInt(params["userId"], 10, 64)

	if !repository.CheckIfUserExists(int(userId)) {
		response.WriteHeader(http.StatusNotFound)
		return
	}

	response.WriteHeader(http.StatusOK)

	var user *model.User

	user = repository.GetUserById(int(userId))

	jsonResponse := user.MarshalJSON()

	response.Write(jsonResponse)
}

func InitializeRestController() {
	r := mux.NewRouter()
	r.HandleFunc("/library/users/", getAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/library/users/", createUser).Methods(http.MethodPost)
	r.HandleFunc("/library/users/{userId}/", getUserById).Methods(http.MethodGet)

	r.HandleFunc("/library/books/", getAllBooks).Methods(http.MethodGet)
	r.HandleFunc("/library/books/", createBook).Methods(http.MethodPost)
	r.HandleFunc("/library/books/{bookId}/", deleteBookById).Methods(http.MethodDelete)
	r.HandleFunc("/library/books/{bookId}/", getBookById).Methods(http.MethodGet)
	r.HandleFunc("/library/books/issue/most/", getMostIssuedBook).Methods(http.MethodGet)
	r.HandleFunc("/library/books/{bookId}/issue/{userId}/", issueBookToUser).Methods(http.MethodPost)
	r.HandleFunc("/library/books/{bookId}/return/", returnBookToLibrary).Methods(http.MethodPost)
	r.HandleFunc("/library/books/{bookId}/issued/", isBookIssued).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
