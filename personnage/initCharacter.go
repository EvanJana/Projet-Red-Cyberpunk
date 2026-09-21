package main

import "fmt"

func initCharacter() character {
	fmt.Println("=== CRÉATION DU PERSONNAGE ===")
	fmt.Println("Entrez le nom de votre personnage (uniquement des lettres) :")
	var saisename string
	fmt.Scan(&saisename)

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Netrunner (100 PV max)")
	fmt.Println("2. Assassin (75 PV max)")
	fmt.Println("3. Berserk (125 PV max)")
	fmt.Println("nom de la classe :")
	var saisieclass string
	fmt.Scan(&saisieclass)

	var pvMax int
	inventaire := make(map[string]int)
	switch saisieclass {
	case "2", "Assassin", "assassin":
		pvMax = 75
		saisieclass = "Assassin"
		inventaire["Stimulant"] = 3
	case "1", "Netrunner", "netrunner":
		pvMax = 100
		saisieclass = "Netrunner"
		inventaire["Stimulant"] = 2
	case "3", "Berserk", "berserk":
		pvMax = 125
		saisieclass = "Berserk"
		inventaire["Stimulant"] = 1
	default:
		fmt.Println("Classe inconnue, attribution de la classe Netrunner par défaut.")
		saisieclass = "Netrunner"
		pvMax = 100
		inventaire["Stimulant"] = 2
	}

	return character{
		name:       saisename,
		class:      saisieclass,
		level:      1,
		pvMax:      pvMax,
		pvAct:      pvMax,
		inventaire: inventaire,
	}
}
