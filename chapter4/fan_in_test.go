package chapter4

import (
	"fmt"
	"testing"
)

func TestFunnel(t *testing.T) {
	sources := make([]<-chan string, 0)
	for i := 0; i < 3; i++ {
		outer := i
		ch := make(chan string)
		sources = append(sources, ch)
		go func() {
			for j := 0; j < 5; j++ {
				ch <- fmt.Sprint(outer, ":", j)
			}
			close(ch)
		}()

	}

	dest := Funnel(sources...)
	for d := range dest {
		fmt.Println(d)
	}
}
