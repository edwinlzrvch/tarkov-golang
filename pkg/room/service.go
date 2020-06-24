package room

import "../entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetAllRooms() []*entity.Room {
	return s.repo.FindAll()
}
