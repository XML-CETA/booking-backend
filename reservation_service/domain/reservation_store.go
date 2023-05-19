package domain

type ReservationStore interface {
	GetAll() ([]Reservation, error)
	CreateReservation(reservation Reservation) error
	GetFirstActive(accommodation string, dateFrom, dateTo string) (Reservation, error)
}
