package main

import (
	"fmt"
	"sync"
)

func main() {
	var compte uint32 = 0
	var wg sync.WaitGroup
	var m sync.RWMutex // semaphore
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(compte *uint32, wg *sync.WaitGroup, m *sync.RWMutex) {
			defer wg.Done()
			m.Lock()
			*compte++
			m.Unlock()
		}(&compte, &wg, &m)
	}
	wg.Wait()
	fmt.Println(compte)
}
