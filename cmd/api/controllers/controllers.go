package controllers

import (
	"bookstore/cmd/api/models"
	"bookstore/cmd/api/utilities"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	//check for duplicate
	isDuplicate, err := models.CheckDuplicate(book.BookName)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}
	if isDuplicate {
		utilities.ErrorJSON(w, fmt.Errorf("Duplicate Book"))
		return
	}

	_, err = models.CreateNewBook(book)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, book, "book")
	// json.NewEncoder(w).Encode(bookID)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, books, "books")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	bookID := params.ByName("id")
	book, err := models.GetBookByID(bookID)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, book, "book")
}

func GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := models.GetGenres()
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, genres, "genres")
}

func GetBookByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	genre := params.ByName("genre")
	books, err := models.GetBooksByGenre(genre)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, books, "books")
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := models.GetAuthors()
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, authors, "authors")
}

func GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	author := params.ByName("author")
	books, err := models.GetBooksByAuthor(author)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, books, "books")
}

func GetYears(w http.ResponseWriter, r *http.Request) {
	years, err := models.GetYears()
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, years, "years")
}

func GetBooksByYear(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	year, err := strconv.Atoi(params.ByName("release_date"))
	if err != nil {
		utilities.ErrorJSON(w, err)
	}
	books, err := models.GetBooksByYear(year)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, books, "books")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	bookID := params.ByName("id")

	var updatedBook models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	err = models.UpdateBook(bookID, updatedBook)
	if err != nil {
		utilities.ErrorJSON(w, err)
	}

	utilities.WriteJSON(w, http.StatusOK, "Book Updated Successfully", "Success")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	bookID := params.ByName("id")

	err := models.DeleteBook(bookID)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, "Book deleted Successfully", "Success")
}

func Review(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	bookID := params.ByName("id")
	var newReview models.Review
	err := json.NewDecoder(r.Body).Decode(&newReview)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	err = models.MakeReview(bookID, newReview)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	utilities.WriteJSON(w, http.StatusOK, "Review Added Successfully", "review")

}

func GetAverage(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	bookID := params.ByName("id")

	averageRating, err := models.AverageRating(bookID)
	if err != nil {
		utilities.ErrorJSON(w, err)
		return
	}

	response := struct {
		AverageRating float64 `json:"averageRating"`
	}{
		AverageRating: averageRating,
	}

	utilities.WriteJSON(w, http.StatusOK, response, "data")
}
