package services

import (
	"homework-rakamin-go-sql/models"
	"homework-rakamin-go-sql/repository"
)

type MoviesService struct {
	movieRepository repository.RepositoryInterface
}

func NewMoviesService(movieRepository repository.RepositoryInterface) *MoviesService {
	return &MoviesService{
		movieRepository: movieRepository,
	}
}

type ServiceInterface interface {
	CreateMovie(movie *models.Movies) (*models.Movies, error)
	GetMovie(slug string) (models.Movies, error)
	UpdateMovie(movie *models.Movies, slug string) (models.Movies, error)
	DeleteMovie(slug string) error
}

func (ms *MoviesService) CreateMovie(movie *models.Movies) (*models.Movies, error) {
	id, err := ms.movieRepository.CreateMovie(movie)
	if err != nil {
		return nil, err
	}
	movie.ID = id
	return movie, nil
}

func (ms *MoviesService) GetMovie(slug string) (models.Movies, error) {
	movies, err := ms.movieRepository.GetMovie(slug)
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (ms *MoviesService) UpdateMovie(movie *models.Movies, slug string) (movies models.Movies, err error) {
	err = ms.movieRepository.UpdateMovie(movie, slug)
	if err != nil {
		return movies, err
	}

	movies, err = ms.movieRepository.GetMovie(movie.Slug)
	if err != nil {
		return movies, err
	}
	return movies, nil
}

func (ms *MoviesService) DeleteMovie(slug string) error {
	err := ms.movieRepository.DeleteMovie(slug)
	if err != nil {
		return err
	}
	return nil
}
