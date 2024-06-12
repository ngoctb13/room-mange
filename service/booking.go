package service

import (
	"context"
	"room-reservation/ent"
	"room-reservation/repository"
	"time"

	"go.uber.org/zap"
)

type BookingService interface {
	GetBookingsByDate(ctx context.Context, date time.Time) ([]*ent.Booking, error)
	GetBookingsByRoomAndDate(ctx context.Context, roomId int, date time.Time) ([]*ent.Booking, error)
	CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.Booking, error)
	UpdateBooking(ctx context.Context, id int, input ent.UpdateBookingInput) (*ent.Booking, error)
	DeleteBooking(ctx context.Context, id int) (*ent.Booking, error)
}

type bookingService struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewBookingService(repoRegistry repository.Repository, logger *zap.Logger) BookingService {
	return &bookingService{repoRegistry: repoRegistry, logger: logger}
}

func (s *bookingService) GetBookingsByDate(ctx context.Context, date time.Time) ([]*ent.Booking, error) {
	return s.repoRegistry.Booking().GetBookingsByDate(ctx, date)
}

func (s *bookingService) GetBookingsByRoomAndDate(ctx context.Context, roomId int, date time.Time) ([]*ent.Booking, error) {
	return s.repoRegistry.Booking().GetBookingsByRoomAndDate(ctx, roomId, date)
}

func (s *bookingService) CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.Booking, error) {
	return s.repoRegistry.Booking().CreateBooking(ctx, input)
}

func (s *bookingService) DeleteBooking(ctx context.Context, id int) (*ent.Booking, error) {
	return s.repoRegistry.Booking().DeleteBooking(ctx, id)
}

func (s *bookingService) UpdateBooking(ctx context.Context, id int, input ent.UpdateBookingInput) (*ent.Booking, error) {
	return s.repoRegistry.Booking().UpdateBooking(ctx, id, input)
}
