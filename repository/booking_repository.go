package repository

import (
	"context"
	"room-reservation/ent"
	"room-reservation/ent/booking"
	"room-reservation/ent/room"
	"strconv"
	"time"
)

type BookingRepository interface {
	GetBookingsByDate(ctx context.Context, date time.Time) ([]*ent.Booking, error)
	GetBookingsByRoomAndDate(ctx context.Context, roomId int, date time.Time) ([]*ent.Booking, error)
	CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.Booking, error)
	UpdateBooking(ctx context.Context, id int, input ent.UpdateBookingInput) (*ent.Booking, error)
	DeleteBooking(ctx context.Context, id int) (*ent.Booking, error)
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

func (r *bookingRepository) CreateBooking(ctx context.Context, input ent.CreateBookingInput) (*ent.Booking, error) {
	userID, err := strconv.Atoi(input.UserID)
	if err != nil {
		return nil, err
	}
	roomID, err := strconv.Atoi(input.RoomID)
	if err != nil {
		return nil, err
	}

	bookingCreate := r.client.Booking.
		Create().
		SetInformation(input.Information).
		SetStartHour(input.StartHour).
		SetEndHour(input.EndHour).
		SetUserID(userID).
		SetRoomID(roomID)

	if input.BookingDate != nil {
		bookingCreate.SetBookingDate(*input.BookingDate)
	} else {
		bookingCreate.SetBookingDate(time.Now())
	}

	return bookingCreate.Save(ctx)
}

func (r *bookingRepository) UpdateBooking(ctx context.Context, id int, input ent.UpdateBookingInput) (*ent.Booking, error) {
	upd := r.client.Booking.UpdateOneID(id)
	if input.Information != nil {
		upd.SetInformation(*input.Information)
	}
	if input.StartHour != nil {
		upd.SetStartHour(*input.StartHour)
	}
	if input.EndHour != nil {
		upd.SetEndHour(*input.EndHour)
	}
	if input.BookingDate != nil {
		upd.SetBookingDate(*input.BookingDate)
	}
	if input.UserID != nil {
		userID, err := strconv.Atoi(*input.UserID)
		if err != nil {
			return nil, err
		}
		upd.SetUserID(userID)
	}
	if input.RoomID != nil {
		roomID, err := strconv.Atoi(*input.RoomID)
		if err != nil {
			return nil, err
		}
		upd.SetRoomID(roomID)
	}

	return upd.Save(ctx)
}

func (r *bookingRepository) DeleteBooking(ctx context.Context, id int) (*ent.Booking, error) {
	booking, err := r.client.Booking.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := r.client.Booking.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return booking, nil
}
