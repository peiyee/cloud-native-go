package main

import (
	"fmt"

	"github.com/hashicorp/golang-lru/v2"
	"github.com/tjarratt/babble"
)

func main() {
	l, _ := lru.New[int, any](128)
	babbler := babble.NewBabbler()
	for i := 0; i < 256; i++ {
		word := babbler.Babble()
		fmt.Println(i, word)
		l.Add(i, word)
	}
	keys := l.Keys()
	for _, k := range keys {
		v, ok := l.Get(k)
		fmt.Println(k, ok, v)
	}
}
