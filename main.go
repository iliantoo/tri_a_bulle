package main

import(
	"fmt"
)

func main () {
	x := []int{20, 30, 10, 40}
	for i := 0; i < len(x); i++ {
		n := x[i]
		fmt.Printf("valeur tableau : %d\n", n)
	}
}


/*
[64,34,25,12,22,11,90]
[0 , 1]
*/