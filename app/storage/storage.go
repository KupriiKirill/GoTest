package storage

import (
	"github.com/kirillkuprii/gotest/app/contract"
)

// Storage contains methods to manipulate data
type Storage interface {
	Init()
	AddItems([]*contract.Coupon)
	GetAllItems() map[UID]*contract.Coupon
	DeleteItems(ids []UID)
	UpdateItems(items *map[UID]contract.Coupon)
}

// UID is unique id of the coupon in storage
type UID = int

var storages = map[string]Storage{}

// OpenStorage a storage if exist or creates a new one
func OpenStorage(name string) (Storage, error) {
	if storage, ok := storages[name]; ok {
		return storage, nil
	}
	return createStorage(name)
}

func createStorage(name string) (Storage, error) {
	newStorage := new(inMemoryStorage)
	newStorage.Init()
	storages[name] = newStorage
	return newStorage, nil
}
