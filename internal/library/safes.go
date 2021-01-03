package library

import (
	"sync"

	"Caesar/internal/relation"
)

/*
实现一个线程安全的slice
*/

type ThreadsSlice struct {
	locker *sync.RWMutex
	slices []relation.StorePath
}

func (ts *ThreadsSlice) Add(element relation.StorePath) {
	ts.locker.Lock()
	ts.slices = append(ts.slices, element)
	ts.locker.Unlock()
}

func (ts *ThreadsSlice) Get() []relation.StorePath {
	return ts.slices
}

func NewSlice() *ThreadsSlice {
	return &ThreadsSlice{
		locker: &sync.RWMutex{},
		slices: []relation.StorePath{},
	}

}
