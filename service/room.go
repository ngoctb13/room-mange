package service

import (
	"context"
	"room-reservation/ent"
	"room-reservation/repository"

	"go.uber.org/zap"
)

type RoomService struct {
	repo   repository.RoomRepository
	logger *zap.Logger
}

func NewRoomService(repo repository.RoomRepository, logger *zap.Logger) *RoomService {
	return &RoomService{repo: repo, logger: logger}
}

func (s *RoomService) GetAllRooms(ctx context.Context) ([]*ent.Room, error) {
	return s.repo.GetAllRooms(ctx)
}

func (s *RoomService) GetRoomByID(ctx context.Context, id int) (*ent.Room, error) {
	return s.repo.GetRoomByID(ctx, id)
}
