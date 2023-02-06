package main

import "fmt"

// TODO --> 2/9 Tests case failed !!

func hour_glass_2d_add(arr [][]int32) int32 {

	var total int32
	var max int32

	fmt.Printf("len --> %d", len(arr))

	for i := 0; i < len(arr); i += 1 {

		if i+2 == len(arr) {
			return max
		}

		fmt.Printf("i --> %d\n", i)

		// for j := 0; j < len(arr); j += 1 {
		// 	if j+2 <= len(arr)-1 {
		// 		total = total + arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
		// 		fmt.Printf("j --> %d\n", j)
		// 		fmt.Printf("total --> %d\n", total)
		// 	}
		// }

		for j := 0; j < len(arr); j += 1 {
			if j+2 <= len(arr)-1 {
				total = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

				// fmt.Printf("j --> %d\n", j)
				// fmt.Printf("total --> %d\n", total)
			}
			if max < total {
				max = total
				fmt.Printf("MAX --> %d\n", max)
			}
		}
	}
	return max
}

func main() {
	var arr = [][]int32{
		{1, 1, 1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0},
		{1, 1, 1, 0, 0, 0},
		{0, 0, 2, 4, 4, 0},
		{0, 0, 0, 2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	hour_glass_2d_add(arr)
}
