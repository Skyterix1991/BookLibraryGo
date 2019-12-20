package model

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id int
}

func NewUser(id int) *User {
	return &User{Id: id}
}

func (u User) MarshalJSON() []byte {
	e, err := json.Marshal(User{
		Id: u.Id,
	})
	if err != nil {
		fmt.Println(err)

		return nil
	}

	return e
}
