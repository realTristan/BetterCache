package main

import (
	"fmt"
	"time"
)

// Main function
func main() {
	// Init the cache
	var Cache *Cache = Init(-1) // -1 -> Unlimited Size

	// Test the full text search
	Cache.Set("key2", "my name is \"tristan\"")

	// Track speed
	startTime := time.Now()

	// Get the text search result
	res := Cache.FullTextSearch(TextSearch{
		Limit:      -1,
		Query:      []byte("tristan"),
		StrictMode: false,
	})
	// Print result
	fmt.Printf("Full Text Search -> (%v): %v\n\n", time.Since(startTime), res)

	startTime = time.Now()
	Cache.Set("key1", "my name is \"daniel\"")
	fmt.Printf("Set (key1) -> (%v)\n\n", time.Since(startTime))

	startTime = time.Now()
	k := Cache.Get("key1")
	fmt.Printf("Get Key -> (%v): %s\n\n", time.Since(startTime), k)

	startTime = time.Now()
	_cache := Cache.Show()
	fmt.Printf("Show Cache -> (%v): %s\n\n", time.Since(startTime), _cache)
}
