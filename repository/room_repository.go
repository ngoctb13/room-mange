package repository

import (
	"context"
	"room-reservation/ent"
)

type RoomRepository interface {
	GetAllRooms(ctx context.Context) ([]*ent.Room, error)
	GetRoomByID(ctx context.Context, id int) (*ent.Room, error)
}

type roomRepository struct {
	client *ent.Client
}

func NewRoomRepository(client *ent.Client) RoomRepository {
	return &roomRepository{client: client}
}

func (r *roomRepository) GetAllRooms(ctx context.Context) ([]*ent.Room, error) {
	return r.client.Room.Query().AllX(ctx), nil
}

func (r *roomRepository) GetRoomByID(ctx context.Context, id int) (*ent.Room, error) {
	return r.client.Room.Get(ctx, id)
}
