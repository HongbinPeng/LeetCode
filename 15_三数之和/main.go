package main

import (
	"fmt"
	"sort"
)

type Person struct {
	age  int
	name string
}

func threeSum() {
	a1 := []Person{{age: 12, name: "tom"}, {age: 9, name: "tom"}, {age: 15, name: "tom"}}
	sort.Slice(a1, func(i, j int) bool {
		return a1[i].age < a1[j].age
	})
	for idx, p := range a1 {
		fmt.Printf("第%d个元素，年龄：%d,姓名：%s", idx, p.age, p.name)
	}
	// print(a1)
}
func main() {
	threeSum()
}
