package chapter4

import (
	"context"
	"errors"
	"fmt"
	"testing"
)

func TestBreaker(t *testing.T) {
	count := 0
	c := func(ctx context.Context) (string, error) {
		fmt.Println(count)
		if count >= 3 {
			return "success", nil
		}
		count++
		return "", errors.New("unavailable")
	}

	cb := Breaker(c, 2)
	resp, err := cb(context.Background())
	fmt.Println(resp, err)
	resp, err = cb(context.Background())
	fmt.Println(resp, err)
	resp, err = cb(context.Background())
	fmt.Println(resp, err)
	resp, err = cb(context.Background())
	fmt.Println(resp, err)
}
