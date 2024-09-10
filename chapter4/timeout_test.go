package chapter4

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeout(t *testing.T) {
	mySlowFunc := func(input string) (string, error) {
		time.Sleep(time.Second * 2)
		return input, nil
	}
	resp, err := Timeout(mySlowFunc)
	fmt.Println(resp, err)
}
