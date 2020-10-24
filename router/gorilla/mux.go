package gorilla

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rmaxfield85/mylibrary"
	"io/ioutil"
	"net/http"
)

type gorillaObj struct {
}

var store mylibrary.Store
var router *mux.Router

func New(port int, s mylibrary.Store) mylibrary.Router {
	store = s

	r := &gorillaObj{}

	router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc(mylibrary.V1Prefix+mylibrary.AllBooks, allBooks).Methods(http.MethodGet)

	router.HandleFunc(mylibrary.V1Prefix+mylibrary.OneBook+"/{bid}", oneBook).Methods(
		http.MethodGet, http.MethodPatch, http.MethodDelete)

	router.HandleFunc(mylibrary.V1Prefix+mylibrary.OneBook, newBook).Methods(http.MethodPost)

	go http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	return r
}

func (r *gorillaObj) Stop() {}

func encode(wtr http.ResponseWriter, content interface{}) {
	wtr.Header().Add("Content-Type", "application/json; charset=utf-8")

	enc := json.NewEncoder(wtr)
	enc.SetIndent("", "  ")

	if err := enc.Encode(content); err != nil {
		wtr.WriteHeader(http.StatusInternalServerError)
	}
}

func allBooks(wtr http.ResponseWriter, _ *http.Request) {
	encode(wtr, store.List())
}

func oneBook(wtr http.ResponseWriter, req *http.Request) {
	bid := mylibrary.BookID(mux.Vars(req)["bid"])

	switch req.Method {
	case http.MethodGet:
		if book, err := store.Read(bid); err != nil {
			wtr.WriteHeader(http.StatusNotFound)
		} else {
			wtr.Header().Add("BID", string(book.Info.BID))
			encode(wtr, book)
		}

	case http.MethodPatch:
		book := new(mylibrary.Book)

		if doc, err := ioutil.ReadAll(req.Body); err != nil {
			wtr.WriteHeader(http.StatusInternalServerError)
		} else {
			if json.Unmarshal(doc, book) != nil {
				wtr.WriteHeader(http.StatusBadRequest)
			} else {
				book.Info.BID = bid
				if err := store.Update(book); err != nil {
					wtr.WriteHeader(http.StatusNotFound)
				}
			}
		}

	case http.MethodDelete:
		if err := store.Delete(bid); err != nil {
			wtr.WriteHeader(http.StatusNotFound)
		} else {
			wtr.WriteHeader(http.StatusOK)
		}
	}
}

func newBook(wtr http.ResponseWriter, req *http.Request) {
	if doc, err := ioutil.ReadAll(req.Body); err != nil {
		wtr.WriteHeader(http.StatusInternalServerError)
	} else {
		book := new(mylibrary.Book)

		if json.Unmarshal(doc, book) != nil {
			wtr.WriteHeader(http.StatusBadRequest)
		} else {
			if bid, err := store.Create(book); err != nil {
				wtr.WriteHeader(http.StatusInternalServerError)
			} else {
				wtr.WriteHeader(http.StatusOK)
				wtr.Header().Add("BID", string(bid))
			}
		}
	}
}
