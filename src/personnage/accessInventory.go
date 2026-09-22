package main

import (
	"fmt"
	"sort"
)

func accessInventory(c *character) {
	fmt.Println("\n=== INVENTAIRE ===")
	objets := inventoryItems(c.inventaire)
	if len(objets) == 0 {
		fmt.Println("Inventaire vide.")
	} else {
		for numero, objet := range objets {
			fmt.Printf("%d. %s (Quantite : %d)\n", numero+1, objet, c.inventaire[objet])
		}
	}

	fmt.Println("1. Utiliser un objet")
	fmt.Println("2. Acceder au marchand")
	fmt.Println("3. Acceder a l'ordinateur")
	fmt.Println("4. Augmenter la capacitée de l'inventaire")
	fmt.Println("5. Retour")

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1":
		useInventoryItem(c, objets)
	case "2":
		accessMerchant(c)
	case "3":
		accessComputer(c)
	case "4":
		upgradeInventorySlot(c)
	case "5":
		fmt.Println("Retour au menu principal.")
	default:
		fmt.Println("Choix invalide.")
	}
}

func inventoryItems(inventaire map[string]int) []string {
	objets := make([]string, 0, len(inventaire))
	for objet, quantite := range inventaire {
		if quantite > 0 {
			objets = append(objets, objet)
		}
	}
	sort.Strings(objets)
	return objets
}

func useInventoryItem(c *character, objets []string) {
	if len(objets) == 0 {
		fmt.Println("Aucun objet utilisable.")
		return
	}

	fmt.Println("Entrez le numero de l'objet a utiliser :")
	var numero int
	fmt.Scan(&numero)
	if numero < 1 || numero > len(objets) {
		fmt.Println("Objet introuvable dans l'inventaire.")
		return
	}

	objet := objets[numero-1]
	switch objet {
	case "Stimulant":
		takePot(c)
	case "Virus":
		removeInventory(c.inventaire, objet)
		poisonPot(100)
	default:
		fmt.Printf("L'objet %q ne peut pas encore etre utilise directement.\n", objet)
	}
}

func accessComputer(c *character) {
	programmes := make([]string, 0)
	for programme := range programmeSorts {
		if c.inventaire[programme] > 0 {
			programmes = append(programmes, programme)
		}
	}
	sort.Strings(programmes)

	if len(programmes) == 0 {
		fmt.Println("Vous n'avez aucun programme dans votre inventaire.")
		return
	}

	fmt.Println("=== ORDINATEUR ===")
	for numero, programme := range programmes {
		fmt.Printf("%d. %s\n", numero+1, programme)
	}
	fmt.Println("0. Retour")

	var choixProgramme int
	fmt.Scan(&choixProgramme)
	if choixProgramme < 1 || choixProgramme > len(programmes) {
		return
	}

	spellBook(c, programmes[choixProgramme-1])
}
