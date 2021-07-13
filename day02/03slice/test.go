package main

import (
	"fmt"
	"sort"
	"strings"
)

func test() {
	var a = make([]int, 5, 10)
	fmt.Println("a is ", a)
	fmt.Printf("%p", a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println()
	fmt.Println("a is ", a)
	fmt.Printf("%p", a)

}

func countWord(sentence string) {
	words := strings.Split(sentence, " ")
	wordMap := make(map[string]int, 15)
	for _, word := range words {
		value, ok := wordMap[word]
		if !ok {
			wordMap[word] = 1
		} else {
			wordMap[word] = value + 1
		}
		for word, count := range wordMap {
			fmt.Println(word, count)
		}
	}
}

func main() {
	sentence := "Thank you for your time you are the best"
	countWord(sentence)

	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	numbers := []int{5, 8, 5, 2, 8, 2, 4, 6, 3, 7, 9, 1, 3}
	sort.Ints(numbers)
	for _, num := range numbers {
		fmt.Println(num)
	}
}
