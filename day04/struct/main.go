package main

import "fmt"

type person struct {
	name string
	city string
	age  int8
}

type student struct {
	name string
	age  int
}

func main() {
	person1 := person{
		name: "小明",
		city: "上海",
		age:  18,
	}

	person2 := &person{
		"小红",
		"深圳",
		45,
	}

	fmt.Printf("person 1 is %#v\n", person1)
	fmt.Printf("person 2 is %#v\n", person2)

	fmt.Println(person1.age, person2.age)

	m := make(map[string]*student)
	stus := [3]student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		fmt.Printf("%p ", &stu)
		fmt.Println(stu.name)
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

}
