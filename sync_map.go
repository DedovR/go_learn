package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	value any
	date  time.Time
}

type SyncMap struct {
	mu   sync.Mutex
	data map[string]*Value
	ttl  time.Duration
}

func main() {
	sm := &SyncMap{data: map[string]*Value{}, ttl: 2 * time.Second}
	sm.Set("a", 1)
	fmt.Println(sm.Get("a"))
	time.Sleep(3 * time.Second)
	fmt.Println(sm.Get("a"))
}

func (sm *SyncMap) Set(key string, value any) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	v := &Value{
		value: value,
		date:  time.Now(),
	}
	sm.data[key] = v
}

func (sm *SyncMap) Get(key string) (any, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	v, ok := sm.data[key]
	if !ok {
		return nil, false
	}

	now := time.Now()
	if now.After(v.date.Add(sm.ttl)) {
		// delete from map
		return nil, false
	}

	return v.value, true
}


