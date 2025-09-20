package main

import (
	"fmt"
	"runtime"
	"testing"
)

const testSize = 1000000 // 1 million entries for testing

// memory usage of boolean map
func BenchmarkBooleanMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolMap := createBooleanMap(testSize)
		_ = boolMap
	}
}

// memory usage of struct map
func BenchmarkStructMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		structMap := createStructMap(testSize)
		_ = structMap
	}
}

// memory usage of integer map
func BenchmarkIntMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intMap := createIntMap(testSize)
		_ = intMap
	}
}

// demonstrates memory usage differences
func TestMemoryComparison(t *testing.T) {
	fmt.Printf("\n=== Memory Usage Comparison ===\n")

	// Test boolean map
	initial := memoryUsage()
	boolMap := createBooleanMap(testSize)
	fmt.Printf("Bool map size: %d\n", len(boolMap))
	boolMemory := memoryUsage() - initial
	fmt.Printf("Boolean map (%d entries): %.2f MB\n",
		testSize, float64(boolMemory)/1024/1024)

	// Cleanup and GC
	boolMap = nil
	runtime.GC()
	runtime.GC() // Double GC to ensure cleanup

	// Test struct map
	initial = memoryUsage()
	structMap := createStructMap(testSize)
	fmt.Printf("Struct map size: %d\n", len(structMap))
	structMemory := memoryUsage() - initial
	fmt.Printf("Struct map (%d entries): %.2f MB\n",
		testSize, float64(structMemory)/1024/1024)

	// Cleanup and GC
	structMap = nil
	runtime.GC()
	runtime.GC()

	// Test int map for comparison
	initial = memoryUsage()
	intMap := createIntMap(testSize)
	fmt.Printf("Int map size: %d\n", len(intMap))
	intMemory := memoryUsage() - initial
	fmt.Printf("Integer map (%d entries): %.2f MB\n",
		testSize, float64(intMemory)/1024/1024)

	// Cleanup
	intMap = nil
	runtime.GC()

	// Calculate savings
	if boolMemory > structMemory {
		savings := float64(boolMemory-structMemory) / 1024 / 1024
		fmt.Printf("\nMemory saved using struct{}: %.2f MB\n", savings)
		fmt.Printf("Percentage improvement: %.1f%%\n",
			float64(boolMemory-structMemory)/float64(boolMemory)*100)
	}
}

// tests basic set operations with optimized approach
func TestSetOperations(t *testing.T) {
	fmt.Printf("\n=== Set Operations Test ===\n")

	colors := make(map[string]struct{})
	colors["red"] = struct{}{}
	colors["blue"] = struct{}{}
	colors["green"] = struct{}{}

	fmt.Printf("Initial set size: %d\n", len(colors))
	if _, exists := colors["red"]; exists {
		fmt.Println("✓ red exists in set")
	}

	if _, exists := colors["purple"]; !exists {
		fmt.Println("✓ purple does not exist in set")
		colors["purple"] = struct{}{}
		fmt.Println("✓ purple added to set")
	}

	delete(colors, "blue")
	fmt.Printf("Final set size after removing blue: %d\n", len(colors))
}

// tests with 10 million entries like in the video
func TestLargeScaleMemoryUsage(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping large scale test in short mode")
	}

	fmt.Printf("\n=== Large Scale Memory Test (10M entries) ===\n")

	const largeSize = 10000000 // 10 million entries

	// Test boolean map
	fmt.Println("Testing boolean map...")
	initial := memoryUsage()
	largeBoolMap := createBooleanMap(largeSize)
	fmt.Printf("Bool map size: %d\n", len(largeBoolMap))
	boolMemory := memoryUsage() - initial
	fmt.Printf("Boolean map (10M entries): %.2f MB\n",
		float64(boolMemory)/1024/1024)

	largeBoolMap = nil
	runtime.GC()
	runtime.GC()

	// Test struct map
	fmt.Println("Testing struct map...")
	initial = memoryUsage()
	largeStructMap := createStructMap(largeSize)
	fmt.Printf("Large Struct map size: %d\n", len(largeStructMap))
	structMemory := memoryUsage() - initial
	fmt.Printf("Struct map (10M entries): %.2f MB\n",
		float64(structMemory)/1024/1024)

	largeStructMap = nil
	runtime.GC()

	// Show savings
	if boolMemory > structMemory {
		savings := float64(boolMemory-structMemory) / 1024 / 1024
		fmt.Printf("Memory saved: %.2f MB\n", savings)
	}
}
