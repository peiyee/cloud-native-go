package chapter4

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestThrottle(t *testing.T) {
	e := func(ctx context.Context) (string, error) {
		return "throttling function", nil
	}

	te := Throttle(e, time.Millisecond+500, 3, 3)
	ctx := context.Background()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		idx := i
		go func() {
			resp, err := te(ctx)
			fmt.Println("group1", idx, resp, err)
			wg.Done()
		}()
	}
	wg.Wait()
}
