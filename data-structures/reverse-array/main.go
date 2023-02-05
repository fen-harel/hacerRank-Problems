package main

import (
	"fmt"
)

func r_array(a []int) {

	var len int = len(a)
	fmt.Println(len)

	var r []int
	count := len - 1

	for i := 0; i < len; i += 1 {
		// r[i] = a[count]
		r = append(r, a[count])

		count -= 1
	}
	fmt.Println(a)
	fmt.Println(r)
}

func main() {
	var a = []int{2, 3, 4, 2, 1, 1, 1, 1, 2, 3, 4, 5, 6}
	r_array(a)
}
