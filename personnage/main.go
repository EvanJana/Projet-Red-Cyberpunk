package main

import "fmt"

type character struct {
	name, class         string
	level, pvMax, pvAct int
	inventaire          map[string]int
}

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
		inventaire["stimulant"] = 3
	case "1", "Netrunner", "netrunner":
		pvMax = 100
		saisieclass = "Netrunner"
		inventaire["stimulant"] = 2
	case "3","Berserk", "berserk":
		pvMax = 125
		saisieclass = "Berserk"
		inventaire["stimulant"] = 1
	default:
		fmt.Println("Classe inconnue, attribution de la classe Netrunner par défaut.")
		saisieclass = "Netrunner"
		pvMax = 100
		inventaire["stimulant"] = 2
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

func displayInfo(c character) {
	fmt.Println("\n=== INFORMATIONS DU PERSONNAGE ===")
	fmt.Printf("Nom         : %s\n", c.name)
	fmt.Printf("Classe      : %s\n", c.class)
	fmt.Printf("Niveau      : %d\n", c.level)
	fmt.Printf("Points de vie : %d / %d\n", c.pvAct, c.pvMax)
	fmt.Printf("Inventaire  : stimulant x%d\n", c.inventaire["stimulant"])
	fmt.Println("==================================")
}

func accessInventory(c *character) {
	if len(c.inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}

	for objet, quantite := range c.inventaire {
		fmt.Printf("Objet : %s, Quantité : %d\n", objet, quantite)
    }
	fmt.Println("1. Mes objets")
	fmt.printLn("2. Vendeur")
	fmt.scan(&objet)
	switch	objet{
	case potion
		takePot()
	}
}

func menu(c *character) {
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
	default:
		fmt.Println("Choix invalide.")
	}
}

func takePot(c *character) {
	quantite, disponible := c.inventaire["stimulant"]
	if !disponible || quantite <= 0 {
		fmt.Println("Vous n'avez plus de stimulant.")
		return
	}

	c.inventaire["stimulant"] = quantite - 1
	c.pvAct += 50
	if c.pvAct > c.pvMax {
		c.pvAct = c.pvMax
	}

	fmt.Printf("Stimulant utilisé. Points de vie : %d / %d\n", c.pvAct, c.pvMax)
}

func main() {
	c := initCharacter()
	displayInfo(c)
	menu(&c)
	
}
