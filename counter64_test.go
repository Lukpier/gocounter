package gocounter

import (
	"fmt"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c Counter64
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println(c.Inc(), c.Get()) // increment and get may return different values
			wg.Done()
		}()
	}
	wg.Wait()
}
