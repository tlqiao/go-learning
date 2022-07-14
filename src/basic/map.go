package main
import "fmt"
func main() {
mapDemo()
}

func mapDemo(){
	var map1 = make(map[int]string)
	map1[1]="name"
	map1[2]="age"
	fmt.Printf("the map value is %s %s",map1[1],map1[2])
	map2 := map[string]string{}
	map2["address"]="chengdu"
	fmt.Printf("the map2 values is %s",map2["address"])
}
