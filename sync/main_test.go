package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

const times = 1000

func BenchmarkUseIncrementOperator(b *testing.B) {
	var cnt uint32
	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			cnt++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}

func BenchmarkUseAtomicAddUint32(b *testing.B) {
	var cnt uint32
	var wg sync.WaitGroup

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			atomic.AddUint32(&cnt, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}

func BenchmarkUseSyncMutexLock(b *testing.B) {
	var cnt uint32
	var wg sync.WaitGroup
	mu := new(sync.Mutex)

	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			cnt++
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(cnt)
}
