package nap

import (
	"sync"
	"sync/atomic"
)

var nextId atomic.Uint64

var elemMap = make(map[uint64]*elem)
var elemMu = new(sync.RWMutex)

func AddElem(e *elem) uint64 {
	elemMu.Lock()
	nextId.Add(1)
	id := nextId.Load()
	e.Set("data-nap-id", id)
	elemMap[id] = e
	elemMu.Unlock()
	//fmt.Println("added", id)
	return id
}

func ReleaseElem(id uint64) {
	//fmt.Println("releaseElem", id)
	elemMu.Lock()
	if e, ok := elemMap[id]; ok {
		delete(elemMap, id)
		for _, event := range e.events {
			event.Release()
		}
	}
	elemMu.Unlock()
}
