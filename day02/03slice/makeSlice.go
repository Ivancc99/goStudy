package main

import "fmt"

func gohaha() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))

	s3 := []int{1, 2, 3}
	s4 := s3
	fmt.Println(s3, s4)
	s3[0] = 100
	fmt.Println(s3, s4)

	b1 := [2]int{1, 2}
	b2 := b1
	b1[0] = 99
	fmt.Println(b2, b1)
}
