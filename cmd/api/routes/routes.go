package routes

import (
	"bookstore/cmd/api/controllers"
	"bookstore/cmd/api/middleware"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func InitRoutes() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/new", controllers.CreateBook)
	router.HandlerFunc(http.MethodGet, "/books", controllers.GetBooks)
	router.HandlerFunc(http.MethodGet, "/books/:id", controllers.GetBookById)
	router.HandlerFunc(http.MethodGet, "/genres", controllers.GetGenres)
	router.HandlerFunc(http.MethodGet, "/genre", controllers.GetBookByGenre)
	router.HandlerFunc(http.MethodGet, "/authors", controllers.GetAuthors)
	router.HandlerFunc(http.MethodGet, "/author", controllers.GetBooksByAuthor)
	router.HandlerFunc(http.MethodGet, "/years", controllers.GetYears)
	router.HandlerFunc(http.MethodGet, "/year", controllers.GetBooksByYear)
	router.HandlerFunc(http.MethodPut, "/edit/:id", controllers.UpdateBook)
	router.HandlerFunc(http.MethodDelete, "/delete/:id", controllers.DeleteBook)

	return middleware.EnableCORS(router)
}
