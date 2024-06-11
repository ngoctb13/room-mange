package service

import (
	"context"
	"room-reservation/ent"
	"room-reservation/repository"
	"time"

	"go.uber.org/zap"
)

type BookingService struct {
	repo   repository.BookingRepository
	logger *zap.Logger
}

func NewBookingService(repo repository.BookingRepository, logger *zap.Logger) *BookingService {
	return &BookingService{repo: repo, logger: logger}
}

func (s *BookingService) GetBookingsByDate(ctx context.Context, date time.Time) ([]*ent.Booking, error) {
	return s.repo.GetBookingsByDate(ctx, date)
}

func (s *BookingService) GetBookingsByRoomAndDate(ctx context.Context, roomId int, date time.Time) ([]*ent.Booking, error) {
	return s.repo.GetBookingsByRoomAndDate(ctx, roomId, date)
}
