package main

import "fmt"

type Ennemi struct {
	Nom      string
	PVMax    int
	PVActuel int
	Degats   int
}

func trainingRobot() Ennemi {
	return Ennemi{
		Nom:      "Robot d'entraînement",
		PVMax:    40,
		PVActuel: 40,
		Degats:   5,
	}
}

func trainingFight(c *character) {
	adversaire := trainingRobot()
	tourCombat := 1

	fmt.Println("=== COMBAT D'ENTRAINEMENT ===")
	fmt.Println("Vous affrontez", adversaire.Nom, "avec", adversaire.PVActuel, "PV.")

	for c.pvAct > 0 && adversaire.PVActuel > 0 {
		fmt.Println()
		fmt.Println("================================")
		fmt.Println("           TOUR", tourCombat)
		fmt.Println("================================")
		characterTurn(c, &adversaire)
		if c.pvAct <= 0 || adversaire.PVActuel <= 0 {
			break
		}
		monsterPattern(c, &adversaire)
		tourCombat++
	}

	if adversaire.PVActuel == 0 {
		fmt.Println("Victoire ! Le robot d'entraînement est vaincu.")
	} else {
		fmt.Println("Defaite ! Votre personnage est K.O.")
	}
}

func characterTurn(c *character, adversaire *Ennemi) {
	fmt.Println("=== MENU DE COMBAT ===")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")

	var choix string
	fmt.Scan(&choix)

	switch choix {
	case "1", "Attaquer", "attaquer":
		attackChoice(c, adversaire)
	case "2", "Inventaire", "inventaire":
		accessInventory(c)
	default:
		fmt.Println("Choix invalide.")
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

func monsterPattern(c *character, adversaire *Ennemi) {
	monsterTurn(c, adversaire)
}

func monsterTurn(c *character, adversaire *Ennemi) {
	fmt.Println(adversaire.Nom, "attaque et inflige", adversaire.Degats, "degats.")
	applyDamage(&c.pvAct, adversaire.Degats)
	fmt.Println("PV restants de", c.name, ":", c.pvAct, "/", c.pvMax)
}