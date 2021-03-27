package database

import "github.com/alekseysychev/test-CRUD-auto/internal/domain/vehicle"

type MemoryStorage struct {
	objects map[string]vehicle.Vehicle
}

func New() {
	// vehicle.Vehicle
}
