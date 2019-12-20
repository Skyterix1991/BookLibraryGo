package repository

import (
	"../../model"
)

var availableBooks = make(map[int]*model.Book)
var issuedBooks = make(map[int]*model.Book)
var users = make(map[int]*model.User)

func AddBookToLibrary(book *model.Book) {
	availableBooks[book.Id] = book
}

func AddUser(user *model.User) {
	users[user.Id] = user
}

func GetUserById(id int) *model.User {
	return users[id]
}

func GetBookById(id int) *model.Book {
	var book = availableBooks[id]

	if book != nil {
		return book
	} else {
		return issuedBooks[id]
	}
}

func ReturnBook(book *model.Book) {
	var bookId = book.Id

	if !IsBookIssued(bookId) {
		return
	}

	availableBooks[book.Id] = book
	issuedBooks[book.Id] = nil

	book.CurrentUser = nil
	book.Issued = false
}

func IssueBook(book *model.Book, user *model.User) {
	var bookId = book.Id

	if IsBookIssued(bookId) {
		return
	}

	book.CurrentUser = user
	book.Issued = true
	book.IssuedTimes = book.IssuedTimes + 1

	availableBooks[book.Id] = nil
	issuedBooks[book.Id] = book
}

func IsBookIssued(id int) bool {
	value, _ := availableBooks[id]

	return value == nil
}

func GetMostIssuedBook() *model.Book {
	var mostIssuedBook *model.Book

	for _, element := range availableBooks {
		if element == nil {
			continue
		}

		if mostIssuedBook == nil {
			mostIssuedBook = element
		}

		if element.IssuedTimes > mostIssuedBook.IssuedTimes {
			mostIssuedBook = element
		}
	}

	for _, element := range issuedBooks {
		if element == nil {
			continue
		}

		if mostIssuedBook == nil {
			mostIssuedBook = element
		}

		if element.IssuedTimes > mostIssuedBook.IssuedTimes {
			mostIssuedBook = element
		}
	}

	return mostIssuedBook
}

func CheckIfBookExists(bookId int) bool {
	return GetBookById(bookId) != nil
}

func CheckIfUserExists(bookId int) bool {
	return GetUserById(bookId) != nil
}

func DeleteBookFromLibrary(bookId int) {
	if !CheckIfBookExists(bookId) {
		return
	}

	availableBooks[bookId] = nil
	issuedBooks[bookId] = nil
}

func GetAllBooksInLibrary() []*model.Book {
	var allBooks []*model.Book

	for _, value := range availableBooks {
		if value == nil {
			continue
		}

		allBooks = append(allBooks, value)
	}

	for _, value := range issuedBooks {
		if value == nil {
			continue
		}

		allBooks = append(allBooks, value)
	}

	return allBooks
}

func GetAllUsers() []*model.User {
	var allUsers []*model.User

	for _, value := range users {
		if value == nil {
			continue
		}

		allUsers = append(allUsers, value)
	}

	return allUsers
}
