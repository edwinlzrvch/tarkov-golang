package room

import "github.com/edwinlzrvch/tarkov-golang/pkg/entity"

// Service contains repository
type Service struct {
	repo Repository
}

// NewService creates a new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// GetAllRooms get all available rooms from DB
func (s *Service) GetAllRooms() []*entity.Room {
	return s.repo.FindAll()
}
