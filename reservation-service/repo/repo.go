package repo

import (
	"context"
	"example/grpc/model"
	"example/grpc/utils"
)

type Repository struct {
}

func (repo *Repository) CreateReservation(reservation model.Reservation) error {
	client, cancel := utils.GetDbClient();
	defer cancel()

	coll := client.Database("reservations").Collection("reservations")
	_, err := coll.InsertOne(context.Background(), reservation)

	return err
}
