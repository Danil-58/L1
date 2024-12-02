package main

import (
	"fmt"
	"sync"
)

//Структура безопасной Map
type SafeMap struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (s *SafeMap) Set(key string, val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = val
}
func (s *SafeMap) Get(key string, val int) int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[key]
}

func main() {
	var wg sync.WaitGroup

	safemap := NewMap()
	stringelems := []string{"go", "hello", "world", "mu", "test"}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		for i, v := range stringelems {
			safemap.Set(v, i)
		}
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		for i, v := range stringelems {
			fmt.Println(safemap.Get(v, i))
		}
		wg.Done()
	}(&wg)
	
	wg.Wait()
}