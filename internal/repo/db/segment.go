package db

import (
	"avitotest/internal/database/query"
	"avitotest/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type SegmentRepo struct {
	*sqlx.DB
}

func NewSegmentRepo(db *sqlx.DB) *SegmentRepo {
	return &SegmentRepo{db}
}

func (r *SegmentRepo) CreateSegment(name string) error {
	_, err := r.Query(query.AddSegment, name)
	if err != nil {
		return err
	}
	return nil
}

func (r *SegmentRepo) DeleteSegment(name string) error {
	segmentId := &models.Segment{}
	tx, err := r.Begin()
	if err != nil {
		return err
	}
	err = r.Get(segmentId, query.GetSegmentIdByName, name)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec(query.DeleteSegmentFromUserSegment, segmentId.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(query.DeleteSegment, name)
	if err != nil {
		fmt.Println("aaaa")
		fmt.Println(name)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
