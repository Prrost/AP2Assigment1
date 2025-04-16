package mapping

import (
	inventorypb "github.com/Prrost/assignment1proto/proto/inventory"
	"inventory-service/domain"
)

func ToProtoProduct(object domain.Object) *inventorypb.Product {
	return &inventorypb.Product{
		Id:        int64(object.ID),
		Name:      object.Name,
		Amount:    object.Amount,
		Available: object.Available,
	}
}
