package routes

import (
	"homework-rakamin-go-sql/handlers"

	"github.com/gofiber/fiber/v2"
)

type Routes struct{
	 movieHandler handlers.MovieHandlerInterface
}

func NewRoutes(movieHandler handlers.MovieHandlerInterface) *Routes {
	return &Routes{
		movieHandler: movieHandler,
	}
}

func (r *Routes) InitializeRoutes(app *fiber.App){
		app.Post("/movie",r.movieHandler.AddNewMovies)
		app.Get("/movie/:slug",r.movieHandler.GetMovieBySlug)
		app.Put("/movie/:slug",r.movieHandler.PutMovie)
		app.Delete("/movie/:slug",r.movieHandler.DeleteMovie)
	}