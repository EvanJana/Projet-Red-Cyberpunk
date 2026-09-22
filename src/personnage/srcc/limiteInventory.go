package main

func limiteinv(inventaire map[string]int, capacite int) bool {
	return inventoryCount(inventaire) >= capacite
}

func inventoryCount(inventaire map[string]int) int {
	total := 0
	for _, quantite := range inventaire {
		total += quantite
	}
	return total
}
