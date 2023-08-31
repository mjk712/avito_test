package repo

import (
	"avitotest/internal/models"
	"avitotest/internal/repo/db"

	"github.com/jmoiron/sqlx"
)

type User interface {
	AddUserIntoSegment(*models.UserSegments) error
	GetUserSegments(userId string) ([]*models.UserSegment, error)
}

type Segment interface {
	CreateSegment(name string) error
	DeleteSegment(name string) error
}

type Repositories struct {
	User
	Segment
}

func NewRepositories(pgdb *sqlx.DB) *Repositories {
	return &Repositories{
		User:    db.NewUserRepo(pgdb),
		Segment: db.NewSegmentRepo(pgdb),
	}
}
