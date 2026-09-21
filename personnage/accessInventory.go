package main

import "fmt"

func accessInventory(c *character) {
	if len(c.inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	for objet, quantite := range c.inventaire {
		fmt.Printf("Objet : %s, Quantité : %d\n", objet, quantite)
	}

	fmt.Println("1. Utiliser un objet")
	fmt.Println("2. Accéder au marchand")
	fmt.Println("3. Retour")

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1":
		fmt.Println("Quel objet voulez-vous utiliser ?")
		var objet string
		fmt.Scan(&objet)

		switch objet {
		case "1", "stimulant", "Stimulant":
			takePot(c)
		case "2", "IEM":
			fmt.Println("vous avez utilisez l'IEM")
		case "3", "Armure de combat":
			fmt.Println("vous avez utilisez l'Armure de combat")
		case "4", "Bottes de soldat":
			fmt.Println("Vous avez utilisez les Bottes de soldat")
		case "5", "Gants de précision":
			fmt.Println("Vous avez utilisez les Gants de précision")
		case "6", "Virus":
			fmt.Println("Vous avez utilisez les Virus")

		default:
			if _, existe := c.inventaire[objet]; existe {
				fmt.Printf("L'objet %q ne peut pas encore être utilisé.\n", objet)
			} else {
				fmt.Println("Objet introuvable dans l'inventaire.")
			}
		}
	case "2":
		accessMerchant(c)
	case "3":
		fmt.Println("Retour au menu principal.")
	default:
		fmt.Println("Choix invalide.")
	}
}
