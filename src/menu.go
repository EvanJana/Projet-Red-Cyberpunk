package main

import "fmt"

func menu(c *character) bool {
	fmt.Println("==================================")
	fmt.Println("1. Afficher les informations du personnage")
	fmt.Println("2. Accéder au contenu de l’inventaire")
	fmt.Println("3. Cyberforgeron")
	fmt.Println("4. Combattre")
	fmt.Println("5. Quitter")
	fmt.Println("==================================")

	var choixMenu string
	var choixCombat string
	fmt.Scan(&choixMenu)

	switch choixMenu {
	case "1":
		displayInfo(*c)
		displaySkill(*c)
	case "2":
		accessInventory(c)
	case "3":
		ferailleur(c)
	case "4":
		fmt.Println("1. entrainement")
		fmt.Println("2. combat")
		fmt.Scan(&choixCombat)
		switch choixCombat {
		case "1":
			trainingFight(c)
			resetCombatHP(c)
		case "2":
			combat(c)
			resetCombatHP()
		}
		
	case "5":
		fmt.Println("Au revoir !")
		return false
	default:
		fmt.Println("Choix invalide.")
	}

	return true
}
