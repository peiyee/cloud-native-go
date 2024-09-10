package chapter4

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var errCh = make(chan error, 1)

func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	var ticker *time.Ticker
	var result string
	var err error
	var once sync.Once
	var threshold time.Time

	return func(ctx context.Context) (string, error) {
		currentTime := time.Now()
		threshold = currentTime.Add(d)
		once.Do(func() {
			ticker = time.NewTicker(time.Millisecond * 100)
			go func() {
				defer func() {
					ticker.Stop()
					once = sync.Once{}
					close(errCh)
				}()
				for {
					select {
					case <-ticker.C:
						diff := time.Now().Sub(currentTime)
						fmt.Println(diff.Milliseconds())
						if time.Now().After(threshold) {
							result, err = circuit(ctx)
							fmt.Println(result, err)
							errCh <- err

							return
						}
					case <-ctx.Done():
						err = ctx.Err()
						return
					}
				}
			}()
		})

		return result, err
	}
}
