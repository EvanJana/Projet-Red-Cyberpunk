package main

import "fmt"

func upgradeInventorySlot(c *character) bool {
	const maxUpgrades = 3
	const prixUpgrade = 30
	var choix string

	if c.InventoryUpgradesCnt >= maxUpgrades {
		fmt.Printf("Limite maximale d'améliorations d'inventaire atteinte (%d/3).\n", c.InventoryUpgradesCnt)
		return false
	} else {
		fmt.Println("1. Voulez vous augmenter l'inventaire pour 30 pieces")
		fmt.Println("2. Retour")
		fmt.Scan(&choix)
		if choix == "1" && argent >= prixUpgrade {
			c.InventoryUpgradesCnt++
			c.InventoryCapacity += 10
			argent -= prixUpgrade
			fmt.Printf("Capacite inventaire = %d\n",c.InventoryCapacity)
			fmt.Println("argent =",argent)
			return true
	}
	}
	return false
}
