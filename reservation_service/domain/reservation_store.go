package domain

type ReservationStore interface {
	GetAll() ([]Reservation, error)
	CreateReservation(reservation Reservation) error
	GetFirstByDates(accommodation int32, dateFrom, dateTo string) (Reservation, error)
}
