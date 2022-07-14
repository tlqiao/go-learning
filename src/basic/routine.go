package main
import (
	"fmt"
	"time"
)
func main(){
	go say("hello")
	say("world")
	s := [] int {1,4,6,8}
	c :=make(chan int)
	go sumData(s,c)
	y := <-c
	fmt.Println("the chanle value is x%",y)
	showChanelValue()
}
func say(s string ) {
	for i:=0; i<3;i++	{
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func sumData(s []int, c chan int){
	sum :=0
	for _,value :=range s{
		sum=sum+value
	}
	c <-sum
}

func showChanelValue(){
	ch :=make(chan int, 3)
	ch <- 1
	ch <-3
	ch <-5
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
