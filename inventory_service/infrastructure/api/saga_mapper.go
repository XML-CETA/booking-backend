package api

import (
	events "booking-backend/common/saga/create_order"
	"booking-backend/inventory_service/domain"
)

func mapUpdateProducts(command *events.CreateOrderCommand) map[*domain.Product]int64 {
	products := make(map[*domain.Product]int64)
	for _, item := range command.Order.Items {
		product := &domain.Product{
			ProductId: item.Product.Id,
			ColorCode: item.Product.Color.Code,
		}
		products[product] = -int64(item.Quantity)
	}
	return products
}

func mapRollbackProducts(command *events.CreateOrderCommand) map[*domain.Product]int64 {
	products := make(map[*domain.Product]int64)
	for _, item := range command.Order.Items {
		product := &domain.Product{
			ProductId: item.Product.Id,
			ColorCode: item.Product.Color.Code,
		}
		products[product] = int64(item.Quantity)
	}
	return products
}
