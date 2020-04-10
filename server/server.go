package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque interdum rutrum sodales. Nullam mattis fermentum libero, non volutpat. "
	s2 := "2"
	fmt.Printf("Size of s is: %d \n", unsafe.Sizeof(s))
	fmt.Printf("Size of s is: %d \n", unsafe.Sizeof(s2))
}
