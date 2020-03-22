package slice

import "sync"

// SyncSlice type that can be safely shared between goroutines
type SyncSlice struct {
	sync.RWMutex
	items []interface{}
}

// SyncSliceItem slice item
type SyncSliceItem struct {
	Index int
	Value interface{}
}

// Values an item to the concurrent slice
func (cs *SyncSlice) Values() []interface{} {
	cs.Lock()
	defer cs.Unlock()
	return cs.items
}

// Clear an item to the concurrent slice
func (cs *SyncSlice) Clear() {
	cs.Lock()
	defer cs.Unlock()
	cs.items = []interface{}{}
}

// Append pends an item to the concurrent slice
func (cs *SyncSlice) Append(item interface{}) {
	cs.Lock()
	defer cs.Unlock()
	cs.items = append(cs.items, item)
}

// Iter over the items in the concurrent slice
// Each item is sent over a channel, so that
// we can iterate over the slice using the builin range keyword
func (cs *SyncSlice) Iter() <-chan SyncSliceItem {
	c := make(chan SyncSliceItem)
	f := func() {
		cs.Lock()
		defer cs.Unlock()
		for index, value := range cs.items {
			c <- SyncSliceItem{index, value}
		}
		close(c)
	}
	go f()
	return c
}
