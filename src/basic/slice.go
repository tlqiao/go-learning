package main

import "fmt"

func main() {
	sliceOne()
	var array = [5]int{1, 2, 3, 4, 5}
	sums(array[:])
	sumsTwo(array[:])
	printSlice()
	//sliceTwo()
	//sliceThree()
}
func sliceOne() {
	var myArray [6]int
	var mySlice = myArray[2:4]
	for i := 0; i < len(myArray); i++ {
		myArray[i] = i + 5
	}
	for i := 0; i < len(mySlice); i++ {
		fmt.Printf("slice value is %d \n ", mySlice[i])
	}
}

func sliceTwo() {
	var slice1 []int = make([]int, 10)
	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * 5
	}
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("the value is %d\n", slice1[i])
	}
}

func sliceThree(){
	var slice2 []int = make([]int, 5)
	slice2[0]=2
	slice2[1]=4
	slice2[2]=6
	slice2[3]=8
	slice2[4]=10
	for index,value := range slice2 {
		fmt.Printf("the index is %d and the value is %d \n", index, value)
	}
}

func sums(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s = s + a[i]
	}
	return s
}

func sumsTwo(a [] int) int {
	s :=0
	for key,value :=range a {
		s=s+value
		fmt.Printf("the key is %d" , key)
	}
	fmt.Printf("the value is %d",s)
	return s
}

func printSlice(){
	x := make([]int, 3, 5)
	y :=[] int {1,3,6}
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
	fmt.Printf("the value of y is %d",y)
}
