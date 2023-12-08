package main

import (
	"fmt"
	"log"
)

func seedAccount(fname, lname, pw string, store Storage) *Account {
	acc, err := NewAccount(fname, lname, pw)

	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s Storage) {
	seedAccount("Exe", "Pine", "gobank", s)
}

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	// Seed accounts
	seedAccounts(store)

	server := NewApiServer(":3000", store)
	fmt.Println("Yeah buddy")
	server.Run()
}
