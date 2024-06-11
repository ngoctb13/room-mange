package repository

import (
	"context"
	"room-reservation/ent"
)

type UserRepository interface {
	PreFunction(ctx context.Context) (string, error)
}

type userRepoImpl struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepoImpl{
		client: client,
	}
}

func (rps userRepoImpl) PreFunction(ctx context.Context) (string, error) {
	return "Success - Repository", nil
}
