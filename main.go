/*package main

import (
    "fmt"
    "time"
)

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

/*Début du programme
    Démarrer le chronomètre

    Définir le tableau x <-- [10, 98, 54, 65, 52, 69, 96, 1]

    Pour i = 0 à longueur(x) - 1 incrémenter i faire
        Pour j = 0 à longueur(x) - i - 1 incrémenter j faire
            Si x[j] > x[j + 1] alors
                Échanger x[j] et x[j + 1]
            Fin Si
        Fin Pour
    Fin Pour

    Pour chaque élément dans x faire
        Afficher "valeur tableau : [élément]"
    Fin Pour

    Arrêter le chronomètre
    Afficher "Temps d'exécution : [durée]"
Fin du programme
*/