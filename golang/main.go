package main

import (
	"fmt"

	// "github.com/s4rvessh04/golang/structs"
)

func main() {
	fmt.Println("Hello Friend!")
	// structs.Test()
	// var slice1 = []uint32{1,2}
	// fmt.Printf("Length is %d and capacity is %d\n", len(slice1), cap(slice1))
	// slice1 = append(slice1, 3)
	// fmt.Printf("Length is %d and capacity is %d\n", len(slice1), cap(slice1))

	// Pointers
	var val int32 = 4	
	// var p *int32 = &val
	var ans = square(&val)
	fmt.Println(ans)
}

func square(num *int32) int32 {
	fmt.Printf("\nNum is: %v\n", *num)
	*num *= *num
	return *num
}
