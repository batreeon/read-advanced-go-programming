package kvstore

import (
	"fmt"
	"math/rand"
	"time"
)

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered
	p.mu.Lock()
	p.filter[id] = func(key string) { 
		ch <- key 
	}
	p.mu.Unlock()
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}
