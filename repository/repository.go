package repository

import (
	"context"
	"fmt"
	"room-reservation/ent"

	"github.com/pkg/errors"
)

type Repository interface {
	User() UserRepository
	Booking() BookingRepository
	Room() RoomRepository

	// DoInTx executes the given function in a transaction.
	DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error
}

// RepoImpl is implementation of Repository
type RepoImpl struct {
	// BankAccountRepository
	entClient *ent.Client
	entTx     *ent.Tx
	user      UserRepository
	booking   BookingRepository
	room      RoomRepository
}

// NewRepository creates new repository registry
func NewRepository(entClient *ent.Client) Repository {
	return &RepoImpl{
		entClient: entClient,
		user:      NewUserRepository(entClient),
		booking:   NewBookingRepository(entClient),
		room:      NewRoomRepository(entClient),
	}
}

func (r *RepoImpl) User() UserRepository {
	return r.user
}

func (r *RepoImpl) Booking() BookingRepository {
	return r.booking
}

func (r *RepoImpl) Room() RoomRepository {
	return r.room
}

// DoInTx executes the given function in a transaction.
func (r *RepoImpl) DoInTx(ctx context.Context, txFunc func(ctx context.Context, repoRegistry Repository) error) error {
	if r.entTx != nil {
		return errors.WithStack(errors.New("invalid tx state, no nested tx allowed"))
	}

	tx, err := r.entClient.Tx(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	commited := false

	defer func() {
		if commited {
			return
		}
		// rollback if not commited
		_ = tx.Rollback()
	}()

	impl := &RepoImpl{
		entTx: tx,
		user:  NewUserRepository(tx.Client()),
	}

	if err := txFunc(ctx, impl); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return errors.WithStack(fmt.Errorf("failed to commit tx: %s", err.Error()))
	}

	commited = true
	return nil
}
