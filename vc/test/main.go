package main

import "fmt"

func main() {
	 f()
	 fmt.Println("complete")
}

func f()  {
	defer func() {
		if r :=recover(); r != nil {
			fmt.Println("recover")
		}
		fmt.Println("calling g")
		fmt.Println("Returned normally from g.")
	}()

	g(0)
}

func g(i int)  {
	if i>3 {
		fmt.Println("Panicking")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("defer")
	fmt.Println("printing int g", i)
	g(i+1)
}
