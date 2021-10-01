package repository

import (
	"booking/internal/models"
	"time"
)

type DatabaseRepo interface {
	InsertReservation(res models.Reservation) (int, error)
	InsetRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDates(start, end time.Time, roomID int) (bool, error)
}
