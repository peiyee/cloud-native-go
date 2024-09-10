package chapter4

import (
	"fmt"
	"testing"
)

func TestShardedMap(t *testing.T) {
	shardedMap := NewShardedMap(5)
	fmt.Println(shardedMap.Get("abc"))
	shardedMap.Set("abc", "1")
	shardedMap.Set("beta", "2")
	shardedMap.Set("gama", "3")
	shardedMap.Set("theta", "4")
	fmt.Println(shardedMap.Get("abc"))
	fmt.Println(shardedMap.Get("beta"))
	fmt.Println(shardedMap.Get("gama"))
	fmt.Println(shardedMap.Get("theta"))
	keys := shardedMap.GetKeys()
	fmt.Println(keys)
}
