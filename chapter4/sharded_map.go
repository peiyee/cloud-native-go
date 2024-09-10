package chapter4

import (
	"crypto/sha1"
	"sync"
)

type Shard struct {
	rw   sync.RWMutex
	data map[string]string
}

type ShardedMap []*Shard

func NewShardedMap(n int) ShardedMap {
	shards := make([]*Shard, n)
	for i := 0; i < n; i++ {
		shards[i] = &Shard{
			data: make(map[string]string),
		}
	}
	return shards
}

func (s ShardedMap) getShardKey(key string) int {
	checksum := sha1.Sum([]byte(key))
	hash := int(checksum[17])
	return hash % len(s)
}

func (s ShardedMap) getShard(key string) *Shard {
	shardKey := s.getShardKey(key)
	return s[shardKey]
}

func (s ShardedMap) GetKeys() []string {
	keys := make([]string, 0)
	m := sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(len(s))

	for _, shard := range s {
		ss := shard
		go func() {
			ss.rw.RLock()
			defer func() {
				wg.Done()
				ss.rw.RUnlock()
			}()
			for k := range ss.data {
				m.Lock()
				keys = append(keys, k)
				m.Unlock()
			}
		}()
	}
	wg.Wait()
	return keys
}

func (s ShardedMap) Get(input string) string {
	shard := s.getShard(input)
	shard.rw.RLock()
	defer shard.rw.RUnlock()
	return shard.data[input]
}

func (s ShardedMap) Set(key, value string) {
	shard := s.getShard(key)
	shard.rw.Lock()
	defer shard.rw.Unlock()

	shard.data[key] = value
}
