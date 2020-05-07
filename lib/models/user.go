package models

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type User struct {
	id        int64
	firstName string
	lastName  string
}

func NewUser(firstName, lastName string) User {
	return User{
		id:        rand.Int63(),
		firstName: firstName,
		lastName:  lastName,
	}
}

func (u User) GetFullName() string {
	return u.firstName + " " + u.lastName
}
