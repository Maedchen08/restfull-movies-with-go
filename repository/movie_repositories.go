package repository

import (
	"gorm.io/gorm"
	"homework-rakamin-go-sql/models"
)

type MoviesRepository struct {
	db *gorm.DB
}

func NewMoviesRepository(db *gorm.DB) *MoviesRepository {
	return &MoviesRepository{db: db}
}

type RepositoryInterface interface {
	CreateMovie(movie *models.Movies) (int, error)
	GetMovie(slug string) (models.Movies, error)
	UpdateMovie(movie *models.Movies, slug string) error
	DeleteMovie(slug string) error
}

func (mr *MoviesRepository) CreateMovie(movie *models.Movies) (int, error) {
	err := mr.db.Create(movie).Error
	if err != nil {
		return movie.ID, err
	}
	return movie.ID, nil
}

func (mr *MoviesRepository) GetMovie(slug string) (models.Movies, error) {
	var movie models.Movies
	query := `SELECT id,title,slug,description,duration,image FROM movies WHERE slug =?`

	err := mr.db.Raw(query, slug).Scan(&movie).Error
	if err != nil {
		return movie, err
	}

	if movie.ID == 0 {
		return movie, gorm.ErrRecordNotFound
	}
	return movie, nil
}

func (mr *MoviesRepository) UpdateMovie(movie *models.Movies, slug string) error {
	query := `UPDATE movies SET title = ?, description = ?, duration = ?, image = ? , slug = ? WHERE slug = ?`
	result := mr.db.Exec(query, movie.Title, movie.Description, movie.Duration, movie.Image,movie.Slug,slug)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (mr *MoviesRepository) DeleteMovie(slug string) error {
	var movie models.Movies
	result := mr.db.Where("slug = ?", slug).Delete(&movie)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
