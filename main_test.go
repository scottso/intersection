package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func BenchmarkSortedIntersection(b *testing.B) {
	var x, y []int
	for i := 0; i < 16384; i++ {
		x = append(x, rand.Intn(100))
		y = append(x, rand.Intn(100))
	}

	b.ReportAllocs()
	b.ResetTimer()

	sort.Ints(x)
	sort.Ints(y)

	_ = SortedIntersection(x, y)
}

func BenchmarkUnsortedIntersection(b *testing.B) {
	var x, y []int
	for i := 0; i < 16384; i++ {
		x = append(x, rand.Intn(100))
		y = append(x, rand.Intn(100))
	}

	b.ReportAllocs()
	b.ResetTimer()
	_ = Intersection(x, y)
}
