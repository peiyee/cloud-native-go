package chapter4

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestDebounceLast(t *testing.T) {
	var input string
	var err error
	testCircuit := func(ctx context.Context) (string, error) {
		return input, err
	}
	dc := DebounceLast(testCircuit, time.Millisecond*500)
	ctx := context.Background()
	input = "abc1"
	dc(ctx)
	input = "abc2"
	err = errors.New("123")
	dc(ctx)
	input = "ab3"
	dc(ctx)
	//err = nil
	input = "abc4"
	dc(ctx)
	myErr := <-errCh
	fmt.Println(myErr)
}
