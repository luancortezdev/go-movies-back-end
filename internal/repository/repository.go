package repository

import (
	"database/sql"

	"github.com/luancortezdev/go-movies-back-end/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
}
