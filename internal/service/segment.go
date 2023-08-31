package service

import "avitotest/internal/repo"

type SegmentService struct {
	segmentRepo repo.Segment
}

func NewSegmentService(segmentRepo repo.Segment) *SegmentService {
	return &SegmentService{
		segmentRepo: segmentRepo,
	}
}

func (s *SegmentService) CreateSegment(name string) error {
	return s.segmentRepo.CreateSegment(name)
}

func (s *SegmentService) DeleteSegment(name string) error {
	return s.segmentRepo.DeleteSegment(name)
}
