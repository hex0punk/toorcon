package visitors

import (
	"fmt"
	"sync"
)

type count struct {
	num  	int
	lock    sync.Mutex
}

var visitorCount *count = nil
var once sync.Once

func New() *count {
	once.Do(func() {
		visitorCount = &count{}
	})
	return visitorCount
}

func (cn *count) Add() {
	cn.lock.Lock()
	cn.num++
	cn.lock.Unlock()
}

func (cn *count) Subtract() error {
	cn.lock.Lock()

	if cn.num == 0  {
		return fmt.Errorf("visitorCount is 0")
	}

	cn.num--
	cn.lock.Unlock()

	return nil
}

func (cn *count) GetCount() int {
	return cn.num
}


