package localstore

import (
	"errors"
	"github.com/rmaxfield85/mylibrary"
	"strconv"
	"sync/atomic"
)

var id int64

type library struct {
	books map[mylibrary.BookID]*mylibrary.Book
}

func New() *library {
	return &library{
		books: make(map[mylibrary.BookID]*mylibrary.Book, 100),
	}
}

func (s *library) Create(book *mylibrary.Book) (mylibrary.BookID, error) {
	book.Info.BID = mylibrary.BookID(strconv.Itoa(int(atomic.AddInt64(&id, 1))))

	s.books[book.Info.BID] = book

	return book.Info.BID, nil
}

func (s *library) Read(bid mylibrary.BookID) (*mylibrary.Book, error) {
	if book := s.books[bid]; book == nil {
		return nil, errors.New("missing")
	}

	return s.books[bid], nil
}

func (s *library) Update(delta *mylibrary.Book) error {
	book := s.books[delta.Info.BID]

	if book == nil {
		return errors.New("not found")
	}

	if delta.Info.Author != "" {
		book.Info.Author = delta.Info.Author
	}

	if delta.Info.Title != "" {
		book.Info.Title = delta.Info.Title
	}

	if delta.Info.Pages != 0 {
		book.Info.Pages = delta.Info.Pages
	}

	if delta.Content != "" {
		book.Content = delta.Content
	}

	return nil
}

func (s *library) Delete(bid mylibrary.BookID) error {
	delete(s.books, bid)
	return nil
}

func (s *library) List() []mylibrary.BookInfo {
	var books []mylibrary.BookInfo

	for i := range s.books {
		books = append(books, s.books[i].Info)
	}

	return books
}

func (s *library) Stop() {}
