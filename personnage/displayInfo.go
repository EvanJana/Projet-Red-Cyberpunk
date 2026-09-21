package main

import "fmt"

func displayInfo(c character) {
	fmt.Println("\n=== INFORMATIONS DU PERSONNAGE ===")
	fmt.Printf("Nom         : %s\n", c.name)
	fmt.Printf("Classe      : %s\n", c.class)
	fmt.Printf("Niveau      : %d\n", c.level)
	fmt.Printf("Points de vie : %d / %d\n", c.pvAct, c.pvMax)
	fmt.Printf("Inventaire  : stimulant x%d\n", c.inventaire["stimulant"])
	fmt.Println("==================================")
}
