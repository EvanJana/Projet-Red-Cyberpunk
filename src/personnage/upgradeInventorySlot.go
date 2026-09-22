package main

import "fmt"

func upgradeInventorySlot(c *character) bool {
	const maxUpgrades = 3

	if c.InventoryUpgradesCnt >= maxUpgrades {
		fmt.Printf("Limite maximale d'améliorations d'inventaire atteinte (%d/3).\n", c.InventoryUpgradesCnt)
		return false
	}else {
	c.InventoryUpgradesCnt++
	c.InventoryCapacity+=10
	fmt.Println("VOtre inventaire a été amelioré")
	}
	return true
}
