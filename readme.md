<h1 align="center">infux</h1>
<p align="center"><i>A high-performance, in-memory caching library for Go, designed for speed and concurrency.</i></p>

---

## 📖 Overview

`infux` is a lightweight yet powerful in-memory caching library for Go. It's built with performance in mind, utilizing a **sharded map** approach to significantly reduce lock contention and maximize concurrent access.

Whether you're building a web service, a data processing pipeline, or any application requiring fast access to frequently used data, `infux` provides a reliable and efficient solution.

### ✨ Key Features

- **High Performance:** Optimized for speed through minimized lock contention.  
- **Thread-Safe:** Safe for concurrent use across multiple goroutines.  
- **Sharded Architecture:** Employs a sharded map with FNV-1a hashing for even key distribution and reduced contention.  
- **Simple API:** Easy to integrate and use with a straightforward interface.  
- **In-Memory:** Blazing fast access to cached data.  

---

## 🚀 Getting Started

To start using `infux` in your Go project, follow these simple steps.

### Installation

Fetch the library using `go get`:

```bash
go get github.com/VectroLabs/infux
````

### Basic Usage

Here's a quick example demonstrating how to create a cache, set, get, and delete items.

```go
package main

import (
	"fmt"
	"github.com/VectroLabs/infux"
)

func main() {
	// Create a new cache instance
	cache := infux.New()

	// Set a key-value pair
	key := "myKey"
	value := []byte("myValue")
	cache.Set(key, value)
	fmt.Printf("Set: %s -> %s\n", key, string(value))

	// Get an item from the cache
	val, found := cache.Get(key)
	if found {
		fmt.Printf("Get: %s -> %s (found)\n", key, string(val))
	}

	// Check if a key exists
	if cache.Has(key) {
		fmt.Printf("Has: %s (true)\n", key)
	}

	// Get total item count
	fmt.Printf("Total items: %d\n", cache.Len())

	// Delete the key
	cache.Delete(key)
	fmt.Printf("Deleted: %s\n", key)

	// Confirm deletion
	_, found = cache.Get(key)
	if !found {
		fmt.Printf("Get: %s (not found after deletion)\n", key)
	}
}
```

---

## ⚙️ API Reference

### `infux.New()`

Creates a new `infux` cache instance.

### `cache.Set(key string, value []byte)`

Sets a key-value pair in the cache.

* `key`: The key to set.
* `value`: The value to associate with the key.

### `cache.Get(key string) ([]byte, bool)`

Retrieves a value from the cache based on the provided key.

* `key`: The key to retrieve.
* Returns:

  * `[]byte`: The value associated with the key, if found.
  * `bool`: A boolean indicating whether the key was found.

### `cache.Delete(key string)`

Removes a key-value pair from the cache.

* `key`: The key to delete.

### `cache.Has(key string) bool`

Checks whether the given key exists in the cache.

* `key`: The key to check.
* Returns: `true` if the key exists, `false` otherwise.

### `cache.Len() int`

Returns the total number of key-value pairs currently stored in the cache.

---

## 💡 How it Works

`infux` achieves its high performance and thread safety through a **sharded map** architecture.

* **Sharding:** The cache is divided into `shardCount` (256 by default) individual `cacheShard` instances. Each shard is an independent map with its own `sync.RWMutex`.

* **Hashing:** It uses the FNV-1a hash algorithm to determine which shard a key belongs to. This ensures an even distribution of keys across shards.

* **Reduced Contention:** Operations on different keys hit different shards, minimizing the need for global locks and improving throughput.

* **`sync.RWMutex`:** Each shard uses a `sync.RWMutex` to allow multiple readers and synchronized writes.

> The choice of `shardCount = 256` (a power of two) allows efficient bitwise operations for shard selection while maintaining high concurrency.

---

## 🤝 Contributing

Contributions are welcome! If you have suggestions, features, or bugs:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature-name`).
3. Make your changes.
4. Commit (`git commit -m 'feat: Add new feature'`).
5. Push (`git push origin feature/your-feature-name`).
6. Open a Pull Request.

---

## 📄 License

`infux` is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

<p align="center">Made with ❤️ by VectroLabs</p>
