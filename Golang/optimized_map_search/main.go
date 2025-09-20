package main

import (
	"fmt"
	"runtime"
)

// Traditional Approach - Using a map with boolean values for set operations
// This works but we waste the memory by storing boolean values (1 byte each) when we only need to check for key existence
/* func main() {
    colors := map[string]bool{
        "red":    true,
        "green":  true,
        "blue":   false,
        "yellow": true,
    }

    if colors["blue"] {
        fmt.Println("Color exists")
    } else {
        fmt.Println("Color doesn't exist")
    }
} */

// Optimized Approach - Using empty struct as values (zero bytes in memory)
func main() {
	colors := map[string]struct{}{
		"red":    {},
		"green":  {},
		"blue":   {},
		"yellow": {},
	}

	// Checking existence using comma ok idiom
	if _, exists := colors["pink"]; exists {
		fmt.Println("Color exists")
	} else {
		fmt.Println("Color doesn't exist")
	}

	if _, exists := colors["blue"]; exists {
		fmt.Println("Color exists")
	} else {
		fmt.Println("Color doesn't exist")
	}

	// adding new colors to the set
	colors["purple"] = struct{}{}
	fmt.Printf("Total colors in set: %d\n", len(colors))

	// removing from set
	delete(colors, "red")
	fmt.Printf("Colors after removing red: %d\n", len(colors))
}

// memoryUsage returns current allocated memory in bytes
func memoryUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.Alloc
}

// creates a map with boolean values for performance testing
func createBooleanMap(size int) map[string]bool {
	boolMap := make(map[string]bool, size)
	for i := 0; i < size; i++ {
		boolMap[fmt.Sprintf("key%d", i)] = true
	}
	return boolMap
}

// creates a map with empty struct values for performance testing
func createStructMap(size int) map[string]struct{} {
	structMap := make(map[string]struct{}, size)
	for i := 0; i < size; i++ {
		structMap[fmt.Sprintf("key%d", i)] = struct{}{}
	}
	return structMap
}

// creates a map with integer values for comparison
func createIntMap(size int) map[string]int {
	intMap := make(map[string]int, size)
	for i := 0; i < size; i++ {
		intMap[fmt.Sprintf("key%d", i)] = 1
	}
	return intMap
}

// Example of how to use the optimized approach in practice
/* func ExampleOptimizedSet() {
	// Create a visited URLs tracker
	visitedURLs := make(map[string]struct{})

	urls := []string{
		"https://example.com",
		"https://golang.org",
		"https://github.com",
		"https://example.com", // Duplicate
	}

	for _, url := range urls {
		if _, visited := visitedURLs[url]; !visited {
			fmt.Printf("Processing new URL: %s\n", url)
			visitedURLs[url] = struct{}{}
		} else {
			fmt.Printf("Skipping already visited: %s\n", url)
		}
	}
	fmt.Printf("Total unique URLs visited: %d\n", len(visitedURLs))

	// Output:
	// Processing new URL: https://example.com
	// Processing new URL: https://golang.org
	// Processing new URL: https://github.com
	// Skipping already visited: https://example.com
	// Total unique URLs visited: 3
}
*/
