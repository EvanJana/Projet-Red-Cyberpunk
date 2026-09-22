package main

import "fmt"

func limiteinv(inventaire map[string]int) bool {
	total := 0
	for _, quantite := range inventaire {
		total += quantite
	}

	if total >= characterCreation().InventoryCapacity {
		fmt.Println("Vous avez atteint la limite d'objets dans l'inventaire")
		return true
	}
	return false
}
