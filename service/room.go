package service

import (
	"context"
	"room-reservation/ent"
	"room-reservation/repository"

	"go.uber.org/zap"
)

type RoomService interface {
	GetAllRooms(ctx context.Context) ([]*ent.Room, error)
	GetRoomByID(ctx context.Context, id int) (*ent.Room, error)
}

type roomService struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewRoomService(repoRegistry repository.Repository, logger *zap.Logger) RoomService {
	return &roomService{repoRegistry: repoRegistry, logger: logger}
}

func (s *roomService) GetAllRooms(ctx context.Context) ([]*ent.Room, error) {
	return s.repoRegistry.Room().GetAllRooms(ctx)
}

func (s *roomService) GetRoomByID(ctx context.Context, id int) (*ent.Room, error) {
	return s.repoRegistry.Room().GetRoomByID(ctx, id)
}
