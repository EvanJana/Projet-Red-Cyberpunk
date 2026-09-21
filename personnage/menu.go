package main

import "fmt"

func menu(c *character) bool {
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder au contenu de l’inventaire")
	fmt.Println("3. Quitter")

	var choixMenu string
	fmt.Scan(&choixMenu)

	switch choixMenu {
	case "1":
		displayInfo(*c)
	case "2":
		accessInventory(c)
	case "3":
		fmt.Println("Au revoir !")
		return false
	default:
		fmt.Println("Choix invalide.")
	}

	return true
}
