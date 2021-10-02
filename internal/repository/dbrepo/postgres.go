package dbrepo

import (
	"booking/internal/models"
	"context"
	"time"
)

// InsertReservation insert a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var newID int

	stmt := `INSERT INTO public.reservations (first_name,last_name,email,phone,start_date,end_date
			,room_id,created_at,updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`

	err := m.DB.QueryRowContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// InsetRoomRestriction insert a Room Restriction into the database
func (m *postgresDBRepo) InsetRoomRestriction(r models.RoomRestriction) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt := `INSERT INTO public.room_restrictions (start_date,end_date,room_id,reservation_id,restriction_id,
            created_at,updated_at) values ($1, $2, $3, $4, $5, $6, $7)`

	_, err := m.DB.ExecContext(ctx, stmt,
		r.StartDate,
		r.EndDate,
		r.RoomID,
		r.ReservationID,
		r.RestrictionID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}

// SearchAvailabilityByRoomID returns true if availability exists for roomID, and false if not available
func (m *postgresDBRepo) SearchAvailabilityByRoomID(start, end time.Time, roomID int) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var numRows int
	query := `select count(id) from public.room_restrictions where room_id = $1 and $2 < end_date and $3 > start_date;`

	row := m.DB.QueryRowContext(ctx, query, roomID, start, end)
	err := row.Scan(&numRows)

	if err != nil {
		return false, err
	}

	if numRows != 0 {
		return false, err
	}
	return true, nil
}

// SearchAvailabilityForAllRooms returns slice of available rooms, if any, for given date range
func (m *postgresDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rooms []models.Room
	query := `select r.id, r.room_name from rooms r where r.id not in 
			(select room_id from room_restrictions rr where $1 < rr.end_date and $2 > rr.start_date)`

	rows, err := m.DB.QueryContext(ctx, query, start, end)
	if err != nil {
		return rooms, err
	}

	for rows.Next() {
		var room models.Room
		err := rows.Scan(
			&room.ID,
			&room.RoomName,
		)
		if err != nil {
			return rooms, err
		}

		rooms = append(rooms, room)

	}

	if err = rows.Err(); err != nil {
		return rooms, err
	}
	return rooms, nil
}
