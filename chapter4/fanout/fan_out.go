package fanout

import (
	"fmt"
	"sync"
	"time"
)

func Split(source <-chan string, n int) []<-chan string {
	dest := make([]<-chan string, 0)
	for i := 0; i < n; i++ {
		id := i
		d := make(chan string)
		dest = append(dest, d)
		go func() {
			defer close(d)
			for c := range source {
				d <- fmt.Sprintf("original: %s, modified: %s", c, fmt.Sprint("from ID: ", id))
			}
		}()
	}
	return dest
}

func RunMain() {
	source := make(chan string)
	go func() {
		defer close(source)
		for i := 0; i < 2; i++ {
			source <- fmt.Sprint("from main, source: ", i)
		}
	}()

	wg := sync.WaitGroup{}
	dests := Split(source, 2)
	wg.Add(len(dests))
	startTime := time.Now()
	for _, dest := range dests {
		go func(ch <-chan string) {
			defer wg.Done()
			for d := range ch {
				ans := Fibonacci(45)
				fmt.Println(d, ans)
			}
		}(dest)
	}
	wg.Wait()
	finishTime := time.Now()
	diff := finishTime.Sub(startTime).Seconds()
	fmt.Println(diff)
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
