package main

import (
	"fmt"
)

func main() {
	s := make([]byte, 2, 4)
	s[0] = 100
	s1 := [1]byte(s[1:])
	s2 := [2]byte(s)
	// s1数组里元素的地址和s指向的数组的元素地址不一样，s2同理
	fmt.Printf("%T, %v, %p, %p\n", s1, s1[0], &s1[0], &s[1])
	fmt.Printf("%T, %v, %v, %p\n", s2, s2[0], &s2[0], s)
}
