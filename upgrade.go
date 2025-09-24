package main

import (
    "fmt"
    "math/rand"
    "time"
)

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
	start := time.Now()
    var taille int
    fmt.Print("Combien de nombres dans le tableau ? ")
    fmt.Scan(&taille)

    var choix string
    fmt.Print("Souhaitez-vous entrer les nombres vous-même ? (oui/non) : ")
    fmt.Scan(&choix)

    x := make([]int, taille)

    if choix == "oui" || choix == "OUI" || choix == "o" || choix == "O"{
        for i := 0; i < taille; i++ {
            fmt.Printf("Entrez le nombre %d : ", i+1)
            fmt.Scan(&x[i])
        }
    } else {
        rand.Seed(time.Now().UnixNano())
        for i := 0; i < taille; i++ {
            x[i] = rand.Intn(100)
        }

        fmt.Println("Nombres générés automatiquement :")
        for i := 0; i < taille; i++ {
            fmt.Printf("%d ", x[i])
        }
        fmt.Println()
    }

	fmt.Println("Tableau pas trié :")
    for i := 0; i < taille; i++ {
        fmt.Printf("%d ", x[i])
    }
    fmt.Println()

    bubbleSort(x)

    fmt.Println("Tableau trié :")
    for i := 0; i < taille; i++ {
        fmt.Printf("%d ", x[i])
    }
    fmt.Println()
	
	elapsed := time.Since(start)
    fmt.Printf("Temps d'exécution : %s\n", elapsed)
}
