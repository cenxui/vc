package util

import (
	"testing"
	"fmt"
)

func Test(t *testing.T) {
	fmt.Println("\a")
	fmt.Println(string(0x4eac))
}

func TestUntype(t *testing.T) {
	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000');

	fmt.Printf("%T\n", 0) // "int"
	fmt.Printf("%T\n", 0.0) // "float64"
	fmt.Printf("%T\n", 0i) // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}

func TestMap(t *testing.T) {
	v :=  make(map[string]int)
	fmt.Println(v)
	s := &v
	fmt.Println(s)
}