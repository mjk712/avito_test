package models

import "time"

type UserSegments struct {
	UserId            int       `json:"userid"`
	SegmentsForAdd    []string  `json:"add_segments"`
	SegmentsForDelete []string  `json:"delete_segments"`
	RemoveTime        time.Time `json:"remove_time`
}

type UserSegment struct {
	UserId      int    `json:"user_id"`
	SegmentName string `json:"segment_name"`
}
