package main

import "fmt"

// append() 函数为切片添加元素
func appendSlice() {
	s1 := []string{"北京", "上海", "广州"}
	fmt.Printf("s1 is: %v, length of s1: %d, capacity of s1: %d\n", s1, len(s1), cap(s1))

	s3 := s1[:2]
	fmt.Printf("s3 is: %v, length of s3: %d, capacity of s3: %d\n", s3, len(s3), cap(s3))

	s1 = append(s1, "深圳")
	fmt.Printf("s1 is: %v, length of s1: %d, capacity of s1: %d\n", s1, len(s1), cap(s1))

	s3[0] = "呼和浩特"
	fmt.Printf("s3 is: %v, length of s3: %d, capacity of s3: %d\n", s3, len(s3), cap(s3))

	s2 := s1[2:4]
	s2[0] = "南昌"
	fmt.Printf("s2 is: %v, length of s2: %d, capacity of s2: %d\n", s2, len(s2), cap(s2))

	s2 = append(s2, "杭州")
	s2[0] = "武汉"
	fmt.Printf("s2 is: %v, length of s2: %d, capacity of s2: %d\n", s2, len(s2), cap(s2))
	fmt.Printf("s1 is: %v, length of s1: %d, capacity of s1: %d\n", s1, len(s1), cap(s1))

	s1 = append(s1, s2...) // 必须要加三个句号，其作用是将切片分成多个元素
	fmt.Printf("s1 is: %v, length of s1: %d, capacity of s1: %d\n", s1, len(s1), cap(s1))

	fmt.Printf("s2 is: %v, length of s2: %d, capacity of s2: %d\n", s2, len(s2), cap(s2))

	fmt.Println("appendSlice()方法结束\n\n\n\n\n")
}

func copySlice() {
	//a1 := []int{1, 2, 3}
	//a2 := a1
	//a3 := make([]int, 3, 3)
	//copy(a3, a1)
	//
	//fmt.Println(a1, a2, a3)
	//
	//a1[0] = 1001
	//fmt.Println(a1, a2, a3)
	x1 := []int{1, 3, 5, 7, 9}
	s1 := x1[:]
	fmt.Printf("s1: %v, len : %d, cap : %d\n", s1, len(s1), cap(s1))
	// 1. 切片不保留具体的值
	// 2. 切片对应一个具体的数组
	// 3. 底层数组是一块连续的空间
	s1 = append(s1[0:1], s1[2:]...) // 删除切片中第二个元素

	fmt.Printf("s1: %v, len : %d, cap : %d\n", s1, len(s1), cap(s1))
	fmt.Println(x1)
}

func haha() {
	appendSlice()
	copySlice()
}
