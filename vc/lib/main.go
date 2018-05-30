package main

import (
	"fmt"
	"logitech.com/vc/lib/repository"
)



func main() {
	item, err := repository.Query("123456789");
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(item.Camera)
	fmt.Println(item.MAC)
}
