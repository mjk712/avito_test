package db

import (
	"avitotest/internal/database/query"
	"avitotest/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	*sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) AddUserIntoSegment(user *models.UserSegments) error {
	segmentId := &models.Segment{}
	tx, err := r.Begin()
	if err != nil {
		return err
	}
	for _, segment := range user.SegmentsForDelete {
		err := r.Get(segmentId, query.GetSegmentIdByName, segment)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.Exec(query.DeleteUserFromSegment, user.UserId, segmentId.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for _, segment := range user.SegmentsForAdd {
		err := r.Get(segmentId, query.GetSegmentIdByName, segment)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.Exec(query.AddUserIntoSegment, user.UserId, segmentId.Id)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	//after time remove user from segment
	f := func() {
		for _, segment := range user.SegmentsForAdd {
			r.Get(segmentId, query.GetSegmentIdByName, segment)
			r.Exec(query.DeleteUserFromSegment, user.UserId, segmentId.Id)
		}
	}
	d := time.Until(user.RemoveTime)
	time.AfterFunc(d, f)
	return nil
}

func (r *UserRepo) GetUserSegments(userId string) ([]*models.UserSegment, error) {
	var userSegments = make([]*models.UserSegment, 0)
	rows, err := r.Queryx(query.GetUserSegments, userId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var a models.UserSegment
		err = rows.StructScan(&a)
		if err != nil {
			return nil, err
		}
		userSegments = append(userSegments, &a)
	}
	rows.Close()

	return userSegments, nil
}
