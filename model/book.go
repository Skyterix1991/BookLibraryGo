package model

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Id          int
	Name        string
	CurrentUser *User
	Issued      bool
	IssuedTimes int
}

func NewBook(id int, name string) *Book {
	return &Book{Id: id, Name: name, Issued: false, IssuedTimes: 0}
}

func (b Book) MarshalJSON() []byte {
	e, err := json.Marshal(Book{
		Id:          b.Id,
		Name:        b.Name,
		Issued:      b.Issued,
		IssuedTimes: b.IssuedTimes,
		CurrentUser: b.CurrentUser,
	})
	if err != nil {
		fmt.Println(err)

		return nil
	}

	return e
}
