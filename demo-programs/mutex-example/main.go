package main

import (
	"fmt"
	"sync"
)

func main() {
	s := []int{}
	var wg sync.WaitGroup
	var m sync.Mutex

	const iterations = 1000
	wg.Add(iterations)

	for i := 1; i <= iterations; i++ {
		go func() {
			m.Lock()
			defer m.Unlock()
			// if we did not acquire lock in this goroutine,
			// occasionally multiple goroutines could run simultaneously and
			// therefore access the memory of 's' at the same time, when that
			// happens, only one of the operation would take effect, since both
			// goroutines would access the same value of 's' and overwrite each
			// other.
			s = append(s, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(len(s))
}
