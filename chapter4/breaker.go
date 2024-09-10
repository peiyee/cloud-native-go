package chapter4

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Circuit func(ctx context.Context) (string, error)

func Breaker(c Circuit, failureThresHold uint) Circuit {
	// 1) set threshold
	// 2) count consecutive failures
	// 3) if consecutive failures > threshold, check current time > last attempt
	// 		3.1) return error
	//		3.2) reset consecutive failures and last attempt
	var consecutiveFailure = 0
	var lastAttempt = time.Now()
	var r sync.RWMutex

	return func(ctx context.Context) (string, error) {
		r.RLock()
		d := consecutiveFailure - int(failureThresHold)
		if d > 0 {
			// check current time is after 2 seconds from last attempt
			shouldRetryAt := lastAttempt.Add(time.Second * 2)
			current := time.Now()
			if current.Before(shouldRetryAt) {
				r.RUnlock()
				return "", errors.New("service Unavailable")
			}
		}
		r.RUnlock()
		response, err := c(ctx)
		r.Lock()
		defer r.Unlock()
		lastAttempt = time.Now()
		if err != nil {
			consecutiveFailure++
			fmt.Println(consecutiveFailure, "consecutiveFailure")
			fmt.Println(lastAttempt, "lastAttempt")
			return "", err
		}
		consecutiveFailure = 0
		return response, nil
	}
}
