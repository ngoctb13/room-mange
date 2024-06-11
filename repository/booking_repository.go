package repository

import (
	"context"
	"room-reservation/ent"
	"room-reservation/ent/booking"
	"room-reservation/ent/room"
	"time"
)

type BookingRepository interface {
	GetBookingsByDate(ctx context.Context, date time.Time) ([]*ent.Booking, error)
	GetBookingsByRoomAndDate(ctx context.Context, roomId int, date time.Time) ([]*ent.Booking, error)
}

type bookingRepository struct {
	client *ent.Client
}

func NewBookingRepository(client *ent.Client) BookingRepository {
	return &bookingRepository{client: client}
}

func (r *bookingRepository) GetBookingsByDate(ctx context.Context, date time.Time) ([]*ent.Booking, error) {
	start := date.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	return r.client.Booking.Query().Where(
		booking.And(
			booking.BookingDateGTE(start),
			booking.BookingDateLT(end),
		),
	).All(ctx)
}

func (r *bookingRepository) GetBookingsByRoomAndDate(ctx context.Context, roomId int, date time.Time) ([]*ent.Booking, error) {
	start := date.Truncate(24 * time.Hour)
	end := start.Add(24 * time.Hour)

	return r.client.Booking.Query().Where(
		booking.And(
			booking.HasRoomWith(room.ID(roomId)),
			booking.BookingDateGTE(start),
			booking.BookingDateLT(end),
		),
	).All(ctx)
}
