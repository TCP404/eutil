package set

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Remove(1)
	fmt.Println(s)

	s.Add(1, 2, 3, 4, 5, 6, 7)

	s2 := New[int]()
	s2.Add(3, 4, 5, 6, 7, 8, 9)
	s3 := New[int]()
	s2.Add(7, 8, 9, 10, 11, 12)

	// fmt.Println(s.Intersection(s2))

	// fmt.Println(s.Diff(s2))

	fmt.Println(Intersection(New[int]().Add(1, 2, 3), New[int]().Add(2, 3)))
	fmt.Println(Diff(New[int]().Add(1, 2, 3, 4), New[int]().Add(3, 4)))
	fmt.Println(Union(New[int]().Add(1, 2, 3), New[int]().Add(4, 5, 6)))

	fmt.Println(s2.One())
	fmt.Println(s2.One())
	fmt.Println(s2.One())

	fmt.Println(s.Members())
	fmt.Println(s2.Members())

	fmt.Println(s.Members())
	s.Union(s2).Union(s3)
	fmt.Println(s.Members())
}
