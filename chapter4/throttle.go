package chapter4

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Effector func(ctx context.Context) (string, error)

func Throttle(e Effector, d time.Duration, maxTokens, refill int) Effector {
	var tokens = maxTokens
	var once sync.Once
	var w sync.RWMutex
	return func(ctx context.Context) (string, error) {
		once.Do(func() {
			ticker := time.NewTicker(d)
			go func() {
				for {
					select {
					case <-ticker.C:
						fmt.Println("refilling...")
						w.Lock()
						t := tokens + refill
						t = max(t, maxTokens)
						tokens = t
						w.Unlock()
					}
				}
			}()
		})
		w.Lock()
		defer w.Unlock()
		tokens--
		if tokens < 0 {
			return "", errors.New("too many requests")
		}

		return e(ctx)
	}
}
