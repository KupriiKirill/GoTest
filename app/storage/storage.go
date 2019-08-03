package storage

import (
	"fmt"
	"sync"

	"github.com/kirillkuprii/gotest/app/contract"
)

// Storage contains methods to manipulate data
type Storage interface {
	Init()
	AddItems([]*contract.Coupon)
	GetAllItems() map[UID]*contract.Coupon
}

// UID is unique id of the coupon in storage
type UID = int

type inMemoryStorage struct {
	mutex   sync.Mutex
	items   map[UID]*contract.Coupon
	lastUID UID
}

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

func (t *inMemoryStorage) Lock() {
	t.mutex.Lock()
}

func (t *inMemoryStorage) Unlock() {
	t.mutex.Unlock()
}

func (t *inMemoryStorage) Init() {
	t.lastUID = 0
	fmt.Println("storage initialized")
}

func (t *inMemoryStorage) AddItems(coupons []*contract.Coupon) {
	go t.additemsInternal(coupons)
}

func (t *inMemoryStorage) GetAllItems() map[UID]*contract.Coupon {
	return t.items
}

func (t *inMemoryStorage) additemsInternal(coupons []*contract.Coupon) {
	t.Lock()
	defer t.Unlock()
	for _, coupon := range coupons {
		t.lastUID++
		t.items[t.lastUID] = coupon
	}
}
