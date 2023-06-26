package main

import "math/rand"

type Account struct {
	ID            int    `json:"id"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	AccountNumber int64  `json:"accountNumber"`
	Balance       int64  `json:"balance"`
}

func NewAccount(firstname, lastname string) *Account {
	return &Account{
		ID:            rand.Intn(10000),
		FirstName:     firstname,
		LastName:      lastname,
		AccountNumber: int64(rand.Intn(10000000)),
	}
}
