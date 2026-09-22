package main

import (
	"fmt"
	"sort"
)

var choixForge = map[int]string{
	1: "Armure de combat",
	2: "Bottes de soldat",
	3: "Gants de précision",
}

var recettesForge = map[string]map[string]int{
	"Armure de combat":   {"Kevlar": 1, "Acier": 2},
	"Bottes de soldat":   {"Cuir": 2, "Acier": 1},
	"Gants de précision": {"Cuir": 1, "Puce neuronale": 2},
}

func ferailleur(c *character) bool {
	fmt.Println("================================")
	fmt.Println("Bienvenue dans ma forge ! \nQue voulez-vous fabriquer ?")
	numeros := make([]int, 0, len(choixForge))
	for numero := range choixForge {
		numeros = append(numeros, numero)
	}
	sort.Ints(numeros)
	for _, numero := range numeros {
		fmt.Printf("%d : %s\n", numero, choixForge[numero])
	}
	fmt.Println("4 : Retour")
	fmt.Println("================================")
	var nb int
	fmt.Scan(&nb)

	if nb == 4 {
		fmt.Println("Retour au menu précédent.")
		return false
	}

	if nb < 1 || nb > 3 {
		fmt.Println("Je ne sais pas fabriquer ça, désolé.")
		return false
	}
	objet := choixForge[nb]
	recette, existe := recettesForge[objet]
	if existe == false {
		fmt.Println("Cette recette n'est pas disponible.")
		return false
	}

	materiauxManquants := false
	for materiau, quantiteRequise := range recette {
		if c.inventaire[materiau] < quantiteRequise {
			materiauxManquants = true
		}
	}
	fmt.Println("================================")
	if materiauxManquants == true {
		fmt.Println("Il te manque des matériaux pour fabriquer cet objet")
		for materiau, quantiteRequise := range recette {
			fmt.Printf(" - %s : %d requis, %d possédé(s)\n", materiau, quantiteRequise, c.inventaire[materiau])
			fmt.Println("================================")
		}
		return false
	} else {
		if !addInventory(c, objet, 1) {
			fmt.Println("Tu ne peux pas fabriquer cet objet, ton inventaire est plein.")
			return false
		}
		for materiau, quantiteRequise := range recette {
			c.inventaire[materiau] -= quantiteRequise
		}
		fmt.Printf("Fabrication réussie : %s\n", objet)
		fmt.Println("================================")
		return true
	}
}
