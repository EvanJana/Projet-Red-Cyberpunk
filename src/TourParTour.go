package main

import "fmt"

func resetCombatHP(c *character) {
	if c.pvAct > 0 {
		c.pvAct = c.pvMax
	}
}

func combat(c *character) {
	resetCombatHP(c)
	adversaire := monstre(0)
	tourCombat := 1

	fmt.Println("=== COMBAT ===")
	fmt.Println("Vous affrontez", adversaire.Nom, "avec", adversaire.PVActuel, "PV.")

	for c.pvAct > 0 && adversaire.PVActuel > 0 {
		fmt.Println()
		fmt.Println("================================")
		fmt.Println("           TOUR", tourCombat)
		fmt.Println("================================")
		if characterTurn(c, &adversaire) {
			fmt.Println("Vous avez fui le combat.")
			break
		}
		if c.pvAct <= 0 || adversaire.PVActuel <= 0 {
			break
		}
		monstrePattern(c, &adversaire, tourCombat)
		tourCombat++
	}

	if adversaire.PVActuel == 0 {
		fmt.Println("Victoire !", adversaire.Nom, "est vaincu.")
		if adversaire.Drop != "" {
			if adversaire.Drop == "dollars" {
				c.argent += 10
				fmt.Println("Vous récupérez 25 dollars.")
			} else {
				fmt.Println("Vous récupérez :", adversaire.Drop)
				addInventory(c, adversaire.Drop, 1)
			}
		}
	} else if c.pvAct <= 0 {
		fmt.Println("Defaite ! Votre personnage est K.O.")
	} else {
		fmt.Println("Le combat est interrompu.")
	}
}

func trainingFight(c *character) {
	resetCombatHP(c)
	adversaire := robotEntrainement()
	tourCombat := 1

	fmt.Println("=== COMBAT ===")
	fmt.Println("Vous affrontez", adversaire.Nom, "avec", adversaire.PVActuel, "PV.")

	for c.pvAct > 0 && adversaire.PVActuel > 0 {
		fmt.Println()
		fmt.Println("================================")
		fmt.Println("           TOUR", tourCombat)
		fmt.Println("================================")
		if characterTurn(c, &adversaire) {
			fmt.Println("Vous avez fui le combat.")
			break
		}
		if c.pvAct <= 0 || adversaire.PVActuel <= 0 {
			break
		}
		monstrePattern(c, &adversaire, tourCombat)
		tourCombat++
	}

	if adversaire.PVActuel == 0 {
		fmt.Println("Victoire !", adversaire.Nom, "est vaincu.")
		if adversaire.Drop != "" {
			if adversaire.Drop == "dollars" {
				c.argent += 10
				fmt.Println("Vous récupérez 25 dollars.")
			} else {
				fmt.Println("Vous récupérez :", adversaire.Drop)
				addInventory(c, adversaire.Drop, 1)
			}
		}
	} else if c.pvAct <= 0 {
		fmt.Println("Defaite ! Votre personnage est K.O.")
	} else {
		fmt.Println("Le combat est interrompu.")
	}
}

func characterTurn(c *character, adversaire *Ennemi) bool {
	fmt.Println("=== MENU DE COMBAT ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Fuir")

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1", "Attaquer", "attaquer":
		attackChoice(c, adversaire)
		return false
	case "2", "Inventaire", "inventaire":
		accessInventory(c)
		return false
	case "3", "Fuir", "fuir":
		return true
	default:
		fmt.Println("Choix invalide.")
		return false
	}
}

func attackChoice(c *character, adversaire *Ennemi) bool {
	if len(c.skill) == 0 {
		fmt.Println("Aucune attaque disponible.")
		return false
	}

	fmt.Println("=== ATTAQUES ===")
	for numero, skill := range c.skill {
		fmt.Println(numero+1, ".", skill)
	}
	fmt.Println("0. Retour")

	var choix int
	fmt.Scan(&choix)
	if choix == 0 {
		return false
	}
	if choix < 1 || choix > len(c.skill) {
		fmt.Println("Attaque invalide.")
		return false
	}

	skill := c.skill[choix-1]
	degats := skillDamage(skill)
	fmt.Println("Attaque utilisee :", skill)
	fmt.Println("Degats infliges :", degats)
	applyDamage(&adversaire.PVActuel, degats)
	fmt.Println("PV restants de", adversaire.Nom, ":", adversaire.PVActuel, "/", adversaire.PVMax)
	return true
}

func skillDamage(skill string) int {
	switch skill {
	case "Coup de Poing":
		return 5
	case "Surcharge":
		return 10
	case "Crash":
		return 15
	case "Suicide":
		return 20
	default:
		return 5
	}
}

func applyDamage(pvActuel *int, degats int) {
	*pvActuel -= degats
	if *pvActuel < 0 {
		*pvActuel = 0
	}
}

func monstrePattern(c *character, adversaire *Ennemi, tour int) {
	degats := adversaire.Degats
	message := adversaire.Nom + " attaque !"

	switch {
	case adversaire.Boss:
		degats, message = adamSmasherPattern(tour, adversaire.Degats)
	case adversaire.Nom == "Cyberpsycho":
		degats, message = cyberpsychoPattern(tour, adversaire.Degats)
	case adversaire.Style != "":
		var perdSonTour bool
		degats, message, perdSonTour = punkQuartierPattern(adversaire.Style, tour, adversaire.Degats)
		if perdSonTour {
			fmt.Println(message)
			return
		}
	}

	monstreTurn(c, degats, message)
}

func monsterPattern(c *character, adversaire *Ennemi, tour int) {
	monstrePattern(c, adversaire, tour)
}

func monstreTurn(c *character, degats int, message string) {
	fmt.Println(message, "Il inflige", degats, "degats.")
	applyDamage(&c.pvAct, degats)
	fmt.Println("PV restants de", c.name, ":", c.pvAct, "/", c.pvMax)
}

func monsterTurn(c *character, degats int, message string) {
	monstreTurn(c, degats, message)
}