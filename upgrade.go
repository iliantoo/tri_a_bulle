package main

import (
    "fmt"
    "math/rand"
    "time"
	"unsafe"
)

func tri(arr []int) {
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

	// Affichage des adresses mémoire
    fmt.Printf("\nAdresse mémoire de la slice x : %p\n", unsafe.Pointer(&x))
    if taille > 0 {
        fmt.Printf("Adresse du premier élément : %p\n", unsafe.Pointer(&x[0]))
    }

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

    tri(x)

    fmt.Println("Tableau trié :")
    for i := 0; i < taille; i++ {
        fmt.Printf("%d ", x[i])
    }
    fmt.Println()

	elapsed := time.Since(start)
    fmt.Printf("Temps d'exécution : %s\n", elapsed)
}

/*
Début du programme
    Démarrer le chronomètre

    Afficher "Combien de nombres dans le tableau ?"
    Lire taille

    Afficher "Souhaitez-vous entrer les nombres vous-même ? (oui/non) :"
    Lire choix

    Créer un tableau x de taille "taille"

    Si choix est "oui" alors
        Pour i = 0 à taille - 1 incrémenter i, faire 
            Afficher "Entrez le nombre i+1 :"
            Lire x[i]
        Fin Pour
    Sinon
        Initialiser le générateur de nombres aléatoires
        Pour i = 0 à taille - 1 incrémenter i, faire
           nombre aléatoire entre 0 et 99 --> x[i]  
        Fin Pour

        Afficher "Nombres générés automatiquement :"
        Pour chaque élément dans x
            Afficher l'élément
        Fin Pour
    Fin Si

    Afficher "Tableau pas trié :"
    Pour chaque élément dans x
        Afficher l'élément
    Fin Pour

    Appeler la fonction tri(x)

    Afficher "Tableau trié :"
    Pour chaque élément dans x
        Afficher l'élément
    Fin Pour

    Arrêter le chronomètre
    Afficher "Temps d'exécution : [durée]"
Fin du programme


Fonction tri(tableau)
    longueur du tableau --> n  
    Pour i = 0 à n - 2 incrémenter i faire
        Pour j = 0 à n - i - 2 incrémenter j faire
            Si tableau[j] > tableau[j + 1] alors
                Échanger tableau[j] et tableau[j + 1]
            Fin Si
        Fin Pour
    Fin Pour
Fin Fonction
*/