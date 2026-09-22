package main

import (
	"fmt"
	"sort"
)

var argent = 100000
var nbObjet int

var choix = map[int]string{
	1: "Stimulant",
	2: "Programme : Surcharge",
	3: "Programme : Crash",
	4: "Programme : Suicide",
	5: "IEM",
	6: "Armure de combat",
	7: "Bottes de soldat",
	8: "Gants de précision",
	9: "Virus",
}

var objects = map[string]int{
	"Stimulant":             30,
	"Programme : Surcharge": 100,
	"Programme : Crash":     110,
	"Programme : Suicide":   400,
	"IEM":                   50,
	"Armure de combat":      150,
	"Bottes de soldat":      80,
	"Gants de précision":    60,
	"Virus":                 40,
}

func marchand(c *character) {
	fmt.Printf("Argent disponible : $%d\n", argent)

	numeros := make([]int, 0, len(choix))
	for numero := range choix {
		numeros = append(numeros, numero)
	}
	sort.Ints(numeros)

	for _, numero := range numeros {
		objet := choix[numero]
		fmt.Printf("%d. %s - $%d\n", numero, objet, objects[objet])
	}
	fmt.Println("10. Retour")

	var choixAcheteur int
	fmt.Scan(&choixAcheteur)

	if choixAcheteur == 10 {
		return
	}

	objet, existe := choix[choixAcheteur]
	if !existe {
		fmt.Println("Choix invalide.")
		return
	}

	prix := objects[objet]
	if argent < prix {
		fmt.Println("Tu n'as pas assez d'argent.")
		return
	}

	fmt.Printf("Combien de %s veux-tu acheter ?\n", objet)
	var quantite int
	fmt.Scan(&quantite)
	if quantite <= 0 {
		fmt.Println("La quantité doit être supérieure à zéro.")
		return
	}

	prixTotal := prix * quantite
	if argent < prixTotal {
		fmt.Println("Tu n'as pas assez d'argent pour cette quantité.")
		return
	}

	if inventoryCount(c.inventaire)+quantite > c.InventoryCapacity {
		fmt.Printf("Achat annulé : il ne reste que %d place(s) dans l'inventaire.\n", c.InventoryCapacity-inventoryCount(c.inventaire))
		return
	}

	argent -= prixTotal
	c.inventaire[objet] += quantite
	fmt.Printf("Achat effectué : %d x %s\n", quantite, objet)
	fmt.Printf("Il vous reste $%d\n", argent)
}

func addInventory(inventaire map[string]int, objet string, capacite int) {
	if limiteinv(inventaire, capacite) == true {
		fmt.Println("Inventaire plein")
	} else {
		inventaire[objet]++
	}
}

func removeInventory(inventaire map[string]int, objet string) {
	quantite, existe := inventaire[objet]
	if !existe {
		return
	}

	if quantite <= 1 {
		delete(inventaire, objet)
		return
	}

	inventaire[objet] = quantite - 1
}
