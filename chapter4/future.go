package chapter4

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Future interface {
	Result() (string, error)
}

type innerFuture struct {
	once  sync.Once
	wg    sync.WaitGroup
	res   string
	err   error
	resCh <-chan string
	errCh <-chan error
}

func (r *innerFuture) Result() (string, error) {
	r.once.Do(func() {
		r.wg.Add(1)
		defer r.wg.Done()
		r.res = <-r.resCh
		r.err = <-r.errCh
	})
	r.wg.Wait()
	return r.res, r.err
}

func FutureSlowFunction(ctx context.Context) Future {
	resCh := make(chan string)
	slowErrCh := make(chan error)
	go func() {
		select {
		case <-time.After(time.Second * 2):
			resCh <- "I sleep for 2 seconds"
			slowErrCh <- nil
		case <-ctx.Done():
			resCh <- ""
			slowErrCh <- ctx.Err()
		}
	}()
	return &innerFuture{
		resCh: resCh,
		errCh: slowErrCh,
	}
}

func FutureRunMain() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	future := FutureSlowFunction(ctx)
	resp, err := future.Result()
	fmt.Println(resp, err)
}
