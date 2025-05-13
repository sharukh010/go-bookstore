package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sharukh010/go-bookstore/pkg/models"
	"github.com/sharukh010/go-bookstore/pkg/utils"
)


func GetBook(w http.ResponseWriter,r *http.Request){
	newBooks := models.GetAllBooks()
	res,_ := json.Marshal(newBooks)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter,r *http.Request){
	var bookId string = mux.Vars(r)["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails,db:= models.GetBookById(Id)
	if db.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No book found with given id"))
		return
	}
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter,r *http.Request){
	newBook := &models.Book{}
	utils.ParseBody(r,newBook)
	b:=newBook.CreateBook()
	res,_ := json.Marshal(b)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter,r *http.Request){
	var bookId string = mux.Vars(r)["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book,db := models.DeleteBook(Id)
	if db.RowsAffected == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No book found with given id"))
		return
	}
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r,updateBook)
	var bookId string = mux.Vars(r)["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)

	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails,db:=models.GetBookById(Id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}