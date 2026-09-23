package main

import (
	"fmt"
	"sort"
	"strings"
)

func barreVie(actuel, max int) string {
	if max <= 0 {
		return "[ ]"
	}
	remplissage := int(float64(actuel) / float64(max) * 20)
	if remplissage < 0 {
		remplissage = 0
	}
	if remplissage > 20 {
		remplissage = 20
	}
	return "[" + strings.Repeat("#", remplissage) + strings.Repeat("-", 20-remplissage) + "]"
}

func displayInfo(c character) {
	fmt.Println("\n========================================")
	fmt.Println("           FICHE DU PERSONNAGE")
	fmt.Println("========================================")
	fmt.Printf("Nom         : %s\n", c.name)
	fmt.Printf("Classe      : %s\n", c.class)
	fmt.Printf("Niveau      : %d\n", c.level)
	fmt.Printf("Expérience  : %d/%d XP\n", c.exp, experienceRequise(c.level+1))
	fmt.Printf("Argent      : %d pièces\n", c.argent)
	fmt.Println("PV          :", c.pvAct, "/", c.pvMax, barreVie(c.pvAct, c.pvMax))
	fmt.Println("RAM         :", c.ramAct, "/", c.ramMax, barreVie(c.ramAct, c.ramMax))
	fmt.Println("========================================")
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
	fmt.Println("========================================")
}
