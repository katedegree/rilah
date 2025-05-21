package service

import (
	"back/domain/entity"
	"back/domain/repository"
)

type pointService struct {
	pointRepository repository.PointRepository
}

func NewPointService(
	pointRepository repository.PointRepository,
) *pointService {
	return &pointService{
		pointRepository: pointRepository,
	}
}

func (s *pointService) EnsurePoint(pe *entity.PointEntity) error {
	existing, err := s.pointRepository.FindByUserAndGroup(pe)
	if err != nil {
		return err
	}

	if existing == nil {
		_, err = s.pointRepository.Create(pe)
	} else {
		_, err = s.pointRepository.Restore(pe)
	}

	return err
}
