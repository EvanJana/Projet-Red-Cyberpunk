package main

import "fmt"

func menu(c *character) bool {
	fmt.Println("==================================")
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder au contenu de l’inventaire")
	fmt.Println("3. Cyberforgeron")
	fmt.Println("4. Quitter")
	fmt.Println("==================================")

	var choixMenu string
	fmt.Scan(&choixMenu)

	switch choixMenu {
	case "1":
		displayInfo(*c)
		displaySkill(*c)
	case "2":
		accessInventory(c)
	case "3":
		ferailleur(c.inventaire)
	case "4":
		fmt.Println("Au revoir !")
		return false
	default:
		fmt.Println("Choix invalide.")
	}

	return true
}
