package main

import (
	"fmt"
	"sort"
)

func displayInfo(c character) {
	fmt.Println("\n=== INFORMATIONS DU PERSONNAGE ===")
	fmt.Printf("Nom         : %s\n", c.name)
	fmt.Printf("Classe      : %s\n", c.class)
	fmt.Printf("Niveau      : %d\n", c.level)
	fmt.Printf("Points de vie : %d / %d\n", c.pvAct, c.pvMax)
	fmt.Printf("Argent      : %d pièces\n", c.argent)
	fmt.Println("Inventaire :")
	objets := make([]string, 0, len(c.inventaire))
	for nomObjet, quantite := range c.inventaire {
		if quantite > 0 {
			objets = append(objets, nomObjet)
		}
	}
	sort.Strings(objets)
	for numero, nomObjet := range objets {
		fmt.Printf("%d. %s x%d\n", numero+1, nomObjet, c.inventaire[nomObjet])
	}
}
