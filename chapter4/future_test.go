package chapter4

import (
	"fmt"
	"testing"
)

func TestTemp(t *testing.T) {
	resCh := make(chan string)
	inner := &innerFuture{resCh: resCh}

	inner.once.Do(func() {
		inner.res = <-inner.resCh
	})
	go func() {
		resCh <- "testing"
	}()

	fmt.Println(inner.res)
}

func TestFuture(t *testing.T) {
	FutureRunMain()

}
