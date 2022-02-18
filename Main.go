package main

import "fmt"
import "time"

func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
	
}
func sayHello(){
	fmt.Println("Hello Satti G")
}
