package chapter4

import (
	"sync"
)

func Funnel(sources ...<-chan string) <-chan string {
	dest := make(chan string)

	wg := sync.WaitGroup{}
	wg.Add(len(sources))

	for _, ch := range sources {
		go func(c <-chan string) {
			defer wg.Done()
			for n := range c {
				dest <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(dest)
	}()

	return dest
}
