package main
import "fmt"

func main() {
var phone1 Phone
phone1 = new(NokiaPhone)
phone1.call()
phone1 = new(IPhone)
phone1.call()

var getInfo GetInfo
getInfo = new(GetApple)
fmt.Printf("the value is %s" ,getInfo.getName())
}

type Phone interface {
	call()
}
type NokiaPhone struct {
}
func (nokiaPhone NokiaPhone) call(){
	fmt.Printf("this is Nokia")
}
type IPhone struct{}
func (iphone IPhone) call() {
	fmt.Printf("this is Iphone")
}

type GetInfo interface {
	getName() string
}

type GetApple struct{}
func (getApple GetApple) getName() string {
return "apple"
}

