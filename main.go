/*package main

import "fmt"

func main() {
	start := time.Now()
    x := []int{10, 98, 54, 65, 52, 69, 96, 1}

    for i := 0; i < len(x)-1; i++ {
        for j := 0; j < len(x)-i-1; j++ {
            if x[j] > x[j+1] {
                x[j], x[j+1] = x[j+1], x[j]
            }
        }
    }

    for i := 0; i < len(x); i++ {
        fmt.Printf("valeur tableau : %d\n", x[i])
    }
	elapsed := time.Since(start)
    fmt.Printf("Temps d'exécution : %s\n", elapsed)
}
