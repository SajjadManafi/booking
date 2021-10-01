package repository

import "booking/internal/models"

type DatabaseRepo interface {

	InsertReservation(res models.Reservation) error
}
