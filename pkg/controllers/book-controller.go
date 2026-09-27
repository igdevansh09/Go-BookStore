package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/igdevansh09/go-bookstore/pkg/models"
	"github.com/igdevansh09/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newbooks := models.GetAllBooks()
	res, _ := json.Marshal(newbooks)
	w.Header().Set("Content-type", "pkglications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing the bookId in GetBookById controller")
	}
	bookdetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("Content-type", "pkglications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing the bookId in DeleteBook controller")
	}
	book := models.DeleteBook(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-type", "pkglications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatebook := &models.Book{}
	utils.ParseBody(r,updatebook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId,0,0)
	if err!=nil {
		fmt.Println("Error while parsing the bookId in UpdateBook controller")
	}
	bookdetails, db:= models.GetBookById(id)
	if updatebook.Name != "" {
		bookdetails.Name = updatebook.Name
	}
	if updatebook.Author != "" {
		bookdetails.Author = updatebook.Author
	}
	if updatebook.Publication != "" {
		bookdetails.Publication = updatebook.Publication
	}
	db.Save(&bookdetails)
	res, _:= json.Marshal(bookdetails)
	w.Header().Set("Content-type", "pkglications/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
