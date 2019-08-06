package storage

import (
	"fmt"
	"sync"

	"github.com/kirillkuprii/gotest/app/contract"
)

type inMemoryStorage struct {
	mutex   sync.Mutex
	items   map[UID]*contract.Coupon
	lastUID UID
}

func (t *inMemoryStorage) Lock() {
	t.mutex.Lock()
}

func (t *inMemoryStorage) Unlock() {
	t.mutex.Unlock()
}

func (t *inMemoryStorage) Init() {
	t.lastUID = 0
	t.items = make(map[UID]*contract.Coupon)
	fmt.Println("storage initialized")
}

func (t *inMemoryStorage) AddItems(coupons []*contract.Coupon) {
	go t.additemsInternal(coupons)
}

func (t *inMemoryStorage) DeleteItems(ids []UID) {
	go t.deleteItemsInternal(ids)
}

func (t *inMemoryStorage) UpdateItems(items *map[UID]contract.Coupon) {
	go t.updateItemsInternal(items)
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

func (t *inMemoryStorage) deleteItemsInternal(ids []UID) {
	t.Lock()
	defer t.Unlock()
	for _, uid := range ids {
		delete(t.items, uid)
	}
}

func (t *inMemoryStorage) updateItemsInternal(items *map[UID]contract.Coupon) {
	t.Lock()
	defer t.Unlock()
	for uid, coupon := range *items {
		t.items[uid] = &coupon
	}
}
