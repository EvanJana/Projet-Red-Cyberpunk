package main

import  "fmt"
func accessInventory(c *character) {
	if len(c.inventaire) == 0 {
		fmt.Println("Inventaire vide.")
	} else if len(c.inventaire)>10{
		limiteinv(c.inventaire)
	}else{
		for objet, quantite := range c.inventaire {
			fmt.Printf("Objet : %s, Quantité : %d\n", objet, quantite)
		}
	}

	fmt.Println("1. Utiliser un objet")
	fmt.Println("2. Marchand")
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
		case "2", "Lames Mantis":
			fmt.Println("Vous avez obtenu la capacité Lames Mantis")
		case "3", "Gorilla Arms":
			fmt.Println("Vous avez obtenu la capacité Gorilla Arms")
		case "4", "Lanceur de projectiles":
			fmt.Println("Vous avez obtenu la capacité Lanceur de projectiles")
		case "6", "virus":
			poisonPot(10)
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
