package service

import (
	"room-reservation/ent"
	"room-reservation/internal/azuread"
	"room-reservation/repository"

	"go.uber.org/zap"
)

type Service interface {
	Auth() AuthService
	Booking() BookingService
	Room() RoomService
	User() UserService
}

type serviceImpl struct {
	authService    AuthService
	UserService    UserService
	RoomService    RoomService
	BookingService BookingService
}

func NewService(azureADOAuthClient azuread.AzureADOAuth, entClient *ent.Client, logger *zap.Logger) Service {
	repoRegistry := repository.NewRepository(entClient)

	return &serviceImpl{
		authService:    NewAuthService(azureADOAuthClient, logger),
		UserService:    NewUserService(repoRegistry, logger),
		BookingService: NewBookingService(repoRegistry, logger),
		RoomService:    NewRoomService(repoRegistry, logger),
	}
}

// hihih
// Auth returns the AuthService.
func (i serviceImpl) Auth() AuthService {
	return i.authService
}

// Storage returns the StorageService.
func (i serviceImpl) Booking() BookingService {
	return i.BookingService
}

// User returns the UserService.
func (i serviceImpl) User() UserService {
	return i.UserService
}

// News returns the NewsService.
func (i serviceImpl) Room() RoomService {
	return i.RoomService
}
