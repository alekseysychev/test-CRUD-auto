package interfaces

import (
	"github.com/alekseysychev/test-CRUD-auto/internal/domain/brand"
	"github.com/alekseysychev/test-CRUD-auto/internal/domain/model"
)

type Vehicle struct {
	id      string
	brand   *brand.Brand
	model   *model.Model
	price   int
	status  int
	mileage int
}

type VehicleStorage interface {
	Create(v *Vehicle) (*Vehicle, error)
	Read(id string) (*Vehicle, error)
	Update(id string) (*Vehicle, error)
	Delete(id string) (*Vehicle, error)
}
