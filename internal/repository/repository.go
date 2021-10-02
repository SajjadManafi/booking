package repository

import (
	"booking/internal/models"
	"time"
)

type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
	InsetRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}
