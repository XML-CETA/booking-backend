package model

type Status int32

const (
	Waiting Status = iota
	Reserved
	Expired
)

type Reservation struct{
	id int32
	accomodation int32
	offer int32
	dateFrom string
	dateTo string
	status Status
}

