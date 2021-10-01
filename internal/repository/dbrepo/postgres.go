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
