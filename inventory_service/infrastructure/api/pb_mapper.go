package api

import (
	pb "booking-backend/common/proto/inventory_service"
	"booking-backend/inventory_service/domain"
)

func mapProduct(product *domain.Product) *pb.Product {
	productPb := &pb.Product{
		Id:        product.ProductId,
		ColorCode: product.ColorCode,
		Quantity:  int64(product.Quantity),
	}
	return productPb
}
