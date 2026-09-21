package main

import "fmt"

var programmeSorts = map[string]string{
	"Programme : Surcharge": "Surcharge",
	"Programme : Crash":     "Crash",
	"Programme : Suicide":   "Suicide",
}

func spellBook(c *character, programme string) {
	sortNom, existe := programmeSorts[programme]
	if !existe {
		fmt.Println("Ce programme ne correspond à aucun sort.")
		return
	}

	for _, sortAppris := range c.skill {
		if sortAppris == sortNom {
			fmt.Printf("Vous connaissez déjà le sort %s.\n", sortNom)
			return
		}
	}

	c.skill = append(c.skill, sortNom)
	removeInventory(c.inventaire, programme)
	fmt.Printf("Le sort %s a été appris.\n", sortNom)
}
