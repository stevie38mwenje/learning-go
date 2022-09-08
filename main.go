package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("hello world")
	myFunction()
	calculateArea(7)
	contiguous(1, 2, 3)
	p1 := Person{"steve", 4}
	var p2 = Person{Name: "steve", Age: 5}

	fmt.Println(p1)
	fmt.Println(p2)

	arr := [2][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}}
	//for i:=0;i< len(arr);i++{
	//	fmt.Println(arr[i])
	//}
	for _, e := range arr {
		fmt.Printf("%d", e)
	}

}

func myFunction() {
	fmt.Println("hello steve")
}

func calculateArea(radius int) float64 {
	result := math.Pi * math.Pow(float64(radius), 2)
	fmt.Println(result)
	return result
}

func contiguous(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	fmt.Println(sum)
	return sum
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
