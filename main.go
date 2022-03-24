package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"golang.org/x/exp/constraints"
)

func SortedIntersection[T constraints.Ordered](a, b []T) (c []T) {
	var a_ptr, b_ptr int

	for a_ptr < len(a) && b_ptr < len(b) {
		switch {
		case a[a_ptr] == b[b_ptr]:
			c = append(c, a[a_ptr])
			a_ptr += 1
			b_ptr += 1
		case a[a_ptr] < b[b_ptr]:
			a_ptr += 1
		default:
			b_ptr += 1
		}
	}

	return
}

// Doesn't care if they're sorted
func Intersection[T constraints.Ordered](a, b []T) (c []T) {
	m := make(map[T]bool)

	for _, elem := range a {
		m[elem] = true
	}

	for _, elem := range b {
		if _, ok := m[elem]; ok {
			c = append(c, elem)
		}
	}
	return
}

func main() {
	var a, b []int

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 40; i++ {
		a = append(a, rand.Intn(40))
		b = append(b, rand.Intn(40))
	}

	c := Intersection(a, b)
	sort.Ints(c)
	fmt.Println(c)

	sort.Ints(a)
	sort.Ints(b)
	fmt.Println(SortedIntersection(a, b))
}
