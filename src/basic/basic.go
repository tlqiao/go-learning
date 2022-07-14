package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var name string = "My Name is taoli"
	fmt.Printf("%t\n", strings.HasPrefix(name, "My"))
	fmt.Printf("%t\n", strings.Contains(name, "taoli"))
	fmt.Printf("%t\n", strings.Index(name, "tao"))
	t := time.Now()
	fmt.Println(t)
	fmt.Println(add(1, 2, 3))
	fmt.Println(returnMultipleValue(2, 2))
	sum()
	calculate(5)
	readString()
	callback(100, show)
	g := func(a, b int) int {return a-b}
	fmt.Printf("the value is %d", g(5,3))
}

func add(a int, b int, c int) int {
	return a + b + c
}
func returnMultipleValue(a int, b int) (int, int) {
	return 2 * a, 3 * b
}

func sum() {
	var sum int
	for i := 0; i < 5; i++ {
		sum = sum + i
	}
}

func calculate(i int) {
	for i >= 0 {
		i--
		fmt.Printf("the value is %d\n", i)
	}
}

func readString(){
	str  := "hello world"
	for pos, char := range str {
		fmt.Printf("Charactor on position %d is : %c \n", pos, char )
	}
}

func show( a int ,   b int )  {
	fmt.Printf("the value is %d %d \n" , a, b)
}

func callback(y int, f func(int , int)){
	f(y,3)
}

func readArray(){
	 firstArray:= [3]int{1,2,3}
	 for index, value := range firstArray{
		 fmt.Printf("the position is %d and value is %d", index, value)
	 }
}


