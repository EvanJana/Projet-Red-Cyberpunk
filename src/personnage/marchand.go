package main

import (
	"fmt"
	"sort"
)

var nbObjet int

var choix = map[int]string{
	1:  "Stimulant",
	2:  "Programme : Surcharge",
	3:  "Programme : Crash",
	4:  "Programme : Suicide",
	5:  "IEM",
	6:  "Armure de combat",
	7:  "Bottes de soldat",
	8:  "Gants de précision",
	9:  "Virus",
	10: "Augmentation d'inventaire",
}

var objects = map[string]int{
	"Stimulant":                 30,
	"Programme : Surcharge":     100,
	"Programme : Crash":         110,
	"Programme : Suicide":       400,
	"IEM":                       50,
	"Armure de combat":          150,
	"Bottes de soldat":          80,
	"Gants de précision":        60,
	"Virus":                     40,
	"Augmentation d'inventaire": 30,
}

var limitesObjets = map[string]int{
	"Programme : Surcharge":     1,
	"Programme : Crash":         1,
	"Programme : Suicide":       1,
	"Augmentation d'inventaire": 1,
}

func marchand(c *character) {
	fmt.Printf("Argent disponible : %d pièces\n", c.argent)

	numeros := make([]int, 0, len(choix))
	for numero := range choix {
		numeros = append(numeros, numero)
	}
	sort.Ints(numeros)

	for _, numero := range numeros {
		objet := choix[numero]
		fmt.Printf("%d. %s - $%d\n", numero, objet, objects[objet])
	}
	fmt.Println("11. Retour")

	var choixAcheteur int
	fmt.Scan(&choixAcheteur)

	if choixAcheteur == 11 {
		return
	}

	objet, existe := choix[choixAcheteur]
	if !existe {
		fmt.Println("Choix invalide.")
		return
	}

	prix := objects[objet]
	if c.argent < prix {
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

	if objet == "Augmentation d'inventaire" {
		if quantite != 1 {
			fmt.Println("Une seule amélioration peut être achetée à la fois.")
			return
		}
		if c.InventoryUpgradesCnt >= 3 {
			fmt.Println("Tu as déjà acheté le nombre maximum d'améliorations.")
			return
		}
		c.argent -= prix
		c.InventoryUpgradesCnt++
		c.InventoryCapacity += 10
		fmt.Printf("Capacité d'inventaire améliorée : %d places.\n", c.InventoryCapacity)
		fmt.Printf("Il vous reste %d pièces.\n", c.argent)
		return
	}

	if limite, limiteDefinie := limitesObjets[objet]; limiteDefinie && c.inventaire[objet]+quantite > limite {
		fmt.Printf("Limite atteinte : vous ne pouvez posséder qu'un seul exemplaire de %s.\n", objet)
		return
	}

	prixTotal := prix * quantite
	if c.argent < prixTotal {
		fmt.Println("Tu n'as pas assez d'argent pour cette quantité.")
		return
	}

	if !addInventory(c, objet, quantite) {
		return
	}

	c.argent -= prixTotal
	fmt.Printf("Achat effectué : %d x %s\n", quantite, objet)
	fmt.Printf("Il vous reste %d pièces\n", c.argent)
}

func addInventory(c *character, objet string, quantite int) bool {
	if quantite <= 0 {
		return false
	}
	if inventoryCount(c.inventaire)+quantite > c.InventoryCapacity {
		fmt.Printf("Achat annulé : il ne reste que %d place(s) dans l'inventaire.\n", c.InventoryCapacity-inventoryCount(c.inventaire))
		return false
	}
	c.inventaire[objet] += quantite
	return true
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
