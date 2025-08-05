// Package infux provides a high-performance, in-memory caching library.
// It uses a sharded map to minimize lock contention and is optimized for speed.
package infux

import (
	"hash/fnv"
	"sync"
)

// The number of shards to use for the cache.
// This is set to 256 to provide a good balance of concurrency
// and memory overhead. A power of two is generally a good choice.
const shardCount = 256

// Cache is a thread-safe, high-performance in-memory cache.
type Cache struct {
	shards [shardCount]*cacheShard
}

// cacheShard is a single shard of the cache. It contains a map of key-value
// pairs and a read-write mutex to protect access to the map.
type cacheShard struct {
	items map[string][]byte
	mu    sync.RWMutex
}

// New creates and returns a new Cache instance.
func New() *Cache {
	c := &Cache{}
	for i := 0; i < shardCount; i++ {
		c.shards[i] = &cacheShard{
			items: make(map[string][]byte),
		}
	}
	return c
}

// getShard returns the cache shard for a given key.
// It uses the FNV-1a hash algorithm to distribute keys evenly across shards.
func (c *Cache) getShard(key string) *cacheShard {
	hasher := fnv.New32a()
	hasher.Write([]byte(key))
	return c.shards[hasher.Sum32()&(shardCount-1)]
}

// Set adds an item to the cache, replacing any existing item.
// The key must be a string and the value is a byte slice.
func (c *Cache) Set(key string, value []byte) {
	shard := c.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	shard.items[key] = value
}

// Get retrieves an item from the cache.
// It returns the value as a byte slice and a boolean indicating
// whether the key was found.
func (c *Cache) Get(key string) ([]byte, bool) {
	shard := c.getShard(key)
	shard.mu.RLock()
	defer shard.mu.RUnlock()
	val, found := shard.items[key]
	return val, found
}

// Delete removes an item from the cache.
func (c *Cache) Delete(key string) {
	shard := c.getShard(key)
	shard.mu.Lock()
	defer shard.mu.Unlock()
	delete(shard.items, key)
}

// Len returns the total number of items in the cache.
func (c *Cache) Len() int {
	total := 0
	for _, shard := range c.shards {
		shard.mu.RLock()
		total += len(shard.items)
		shard.mu.RUnlock()
	}
	return total
}

// Has checks if a key exists in the cache.
func (c *Cache) Has(key string) bool {
	_, found := c.Get(key)
	return found
}

