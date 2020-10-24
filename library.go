package mylibrary

const (
	V1Prefix = "/library/v1"
	AllBooks = "/books"
	OneBook  = "/book"
)

// BookID is a book's unique identifier
type BookID string

// BookInfo contains book information, excluding its content
type BookInfo struct {
	BID    BookID `json:"bookId"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

// Book is the full definition of a book object
type Book struct {
	Info    BookInfo `json:"info"`
	Content string   `json:"content"`
}

// Store defines the required behavior of a library data store
type Store interface {
	Create(book *Book) (BookID, error)
	Read(bid BookID) (*Book, error)
	Update(book *Book) error
	Delete(bid BookID) error
	List() []BookInfo
	Stop()
}

// Router defines the necessary behavior of a router
type Router interface {
	Stop()
}
