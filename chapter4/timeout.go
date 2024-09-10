package chapter4

import (
	"context"
	"fmt"
	"time"
)

type slowFunc func(string) (string, error)

func Timeout(f slowFunc) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	respCh := make(chan string)
	//errCh := make(chan error)
	go func() {
		resp, _ := f("testing")
		respCh <- resp
	}()

	select {
	case result := <-respCh:
		fmt.Println(result)
		return result, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
