package dbrepo

import (
	"booking/internal/models"
	"errors"
	"time"
)

// InsertReservation insert a reservation into the database
func (m *testDBRepo) InsertReservation(res models.Reservation) (int, error) {
	return 1, nil
}

// InsetRoomRestriction insert a Room Restriction into the database
func (m *testDBRepo) InsetRoomRestriction(r models.RoomRestriction) error {
	return nil
}

// SearchAvailabilityByRoomID returns true if availability exists for roomID, and false if not available
func (m *testDBRepo) SearchAvailabilityByRoomID(start, end time.Time, roomID int) (bool, error) {
	return false, nil
}

// SearchAvailabilityForAllRooms returns slice of available rooms, if any, for given date range
func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	var rooms []models.Room
	return rooms, nil
}

// GetRoomByID gets room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {
	var room models.Room

	if id > 2 {
		return room, errors.New("error")
	}
	return room, nil
}
