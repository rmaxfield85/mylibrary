package main

import (
	"github.com/rmaxfield85/mylibrary"
	"github.com/rmaxfield85/mylibrary/router/gorilla"
	"github.com/rmaxfield85/mylibrary/store/localstore"
	"os"
	"os/signal"
)

const (
	defaultPort = 12250
)

var books = []mylibrary.Book{
	{
		Info: mylibrary.BookInfo{
			Title:  "Two Fox Tales",
			Author: "Bob",
			Pages:  42,
		},
		Content: "Two foxes were walking through the woods when one began to giggle...",
	},
	{
		Info: mylibrary.BookInfo{
			Title:  "Star Warts",
			Author: "Betty",
			Pages:  452,
		},
		Content: "Long, long ago, on the side of a nose far, far away...",
	},
	{
		Info: mylibrary.BookInfo{
			Title:  "Election Woes",
			Author: "Stacy",
			Pages:  1000,
		},
		Content: "In a country not so far away...",
	},
}

func main() {
	store := localstore.New()

	// Create initial inventory
	for i := range books {
		_, _ = store.Create(&books[i])
	}

	router := gorilla.New(defaultPort, store)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until ^C
	<-c

	router.Stop()
	store.Stop()
	os.Exit(0)
}
