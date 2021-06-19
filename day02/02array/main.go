package main

import "fmt"

// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool
	var a2 [4]bool

	fmt.Printf("%T,%T\n", a1, a2)

	// 数组的初始化
	// 1、最普通的初始化
	a1 = [3]bool{true, false, true}
	fmt.Println(a1)

	// 2、根据数组元素推导个数
	a3 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a3)

	// 3、根据索引赋值
	a4 := [5]int{0: 1, 2: 3}
	fmt.Println(a4)

	// 数组的遍历
	cities := []string{"北京", "上海", "成都"}
	for i, city := range cities {
		fmt.Println(i, city)
	}

	var matrix [3][2]int
	matrix = [3][2]int{
		{4, 5},
		{7, 8},
		{1, 2},
	}

	fmt.Println(matrix)

	b1 := [2]int{1, 2}
	b2 := b1
	b1[0] = 99
	fmt.Println(b2, b1)
}
