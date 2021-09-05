package handlers

import (
	"errors"
	"homework-rakamin-go-sql/models"
	"homework-rakamin-go-sql/services"
	"homework-rakamin-go-sql/utils"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MoviesHandler struct {
	movieService services.ServiceInterface
}

func NewMoviesHandler(movieService services.ServiceInterface) *MoviesHandler {
	return &MoviesHandler{
		movieService: movieService,
	}
}

type MovieHandlerInterface interface {
	AddNewMovies(c *fiber.Ctx) error
	GetMovieBySlug(c *fiber.Ctx) error
	PutMovie(c *fiber.Ctx) error
	DeleteMovie(c *fiber.Ctx) error
}

func (mh *MoviesHandler) AddNewMovies(c *fiber.Ctx) error {
	movie := &models.Movies{}

	err := c.BodyParser(movie)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	validate := utils.NewValidator()
	
	if err := validate.Struct(movie); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}

	response, err := mh.movieService.CreateMovie(movie)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":   false,
		"message": response,
	})
}

func (mh *MoviesHandler) GetMovieBySlug(c *fiber.Ctx) error {
	slug := c.Params("slug")
	response, err := mh.movieService.GetMovie(slug)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "data not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "data not found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "success retrieve data ",
		"result":  response,
	})
}

func (mh *MoviesHandler) PutMovie(c *fiber.Ctx) error {
	movie := &models.Movies{}
	slug := c.Params("slug")
	var mysqlErr *mysql.MySQLError

	err1 := c.BodyParser(movie)
	if err1 != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err1.Error(),
		})
	}

	validate := utils.NewValidator()
	err := validate.Struct(movie)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": utils.ValidatorErrors(err),
		})
	}

	response, err := mh.movieService.UpdateMovie(movie, slug)
	if err != nil {
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error":   true,
				"message": err.Error(),
			})
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":   true,
				"message": "database not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":   false,
		"message": "data has been update",
		"result":  response,
	})

}

func (mh *MoviesHandler) DeleteMovie(c *fiber.Ctx) error {
	slug := c.Params("slug")
	err := mh.movieService.DeleteMovie(slug)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":true,
				"message":"data not found",

			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":true,
			"message":err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error":false,
		"message":"success",
	})
}
