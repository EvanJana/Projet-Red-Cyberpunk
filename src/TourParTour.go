package main

import "fmt"

func resetCombatHP(c *character) {
	if c.vientDeMourir {
		c.vientDeMourir = false
		return
	}

	c.pvAct = c.pvMax
}

func affichageCombatDebut(c *character, adversaire Ennemi) {
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Println("             COMBAT EN COURS            ")
	fmt.Println("========================================")
	fmt.Printf("Joueur : %s  |  PV : %d/%d\n", c.name, c.pvAct, c.pvMax)
	fmt.Printf("Adversaire : %s  |  PV : %d/%d\n", adversaire.Nom, adversaire.PVActuel, adversaire.PVMax)
	fmt.Println("========================================")
}

func combat(c *character) {
	resetCombatHP(c)
	c.combatsAleatoires++
	adversaire := monstre(c.combatsAleatoires)
	tourCombat := 1

	affichageCombatDebut(c, adversaire)

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
		gainExperience(c, adversaire.Experience)
		c.argent += 20
		fmt.Println("Vous gagnez 20 dollars pour cette victoire.")
		if adversaire.Drop != "" {
			if adversaire.Drop == "dollars" {
				c.argent += 40
				fmt.Println("Vous récupérez 40 dollars en bonus.")
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

	affichageCombatDebut(c, adversaire)

	for c.pvAct > 0 && adversaire.PVActuel > 0 {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Printf("                 TOUR %d\n", tourCombat)
		fmt.Println("========================================")
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

	if adversaire.PVActuel == 20 {
		fmt.Println("Victoire !", adversaire.Nom, "est vaincu.")
		c.argent += 15
		fmt.Println("Vous gagnez 15 dollars pour cette victoire.")
		if adversaire.Drop != "" {
			if adversaire.Drop == "dollars" {
				c.argent += 10
				fmt.Println("Vous récupérez 10 dollars en bonus.")
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
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Println("              ACTIONS")
	fmt.Println("========================================")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Inventaire")
	fmt.Println("3. Fuir")
	fmt.Println("========================================")

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

	fmt.Println("")
	fmt.Println("========================================")
	fmt.Println("               ATTAQUES")
	fmt.Println("========================================")
	for numero, skill := range c.skill {
		fmt.Printf("%d. %s\n", numero+1, skill)
	}
	fmt.Println("0. Retour")
	fmt.Println("========================================")

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
	degats := skillDamage(skill) + (c.level-1)*5
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Printf("%s utilise %s\n", c.name, skill)
	fmt.Printf("Dégâts infligés : %d\n", degats)
	fmt.Println("========================================")
	applyDamage(&adversaire.PVActuel, degats)
	fmt.Printf("PV restants de %s : %d/%d\n", adversaire.Nom, adversaire.PVActuel, adversaire.PVMax)
	return true
}

func skillDamage(skill string) int {
	switch skill {
	case "Coup de Poing":
		return 10
	case "Surcharge":
		return 15
	case "Crash":
		return 20
	case "Suicide":
		return 250
	default:
		return 5
	}
}

func gainExperience(c *character, experience int) {
	ancienNiveau := c.level
	c.exp += experience
	c.level = c.exp/100 + 1

	if c.level == ancienNiveau {
		fmt.Printf("Vous gagnez %d XP. Total : %d/%d XP.\n", experience, c.exp, c.level*100)
		return
	}

	nouveauxNiveaux := c.level - ancienNiveau
	c.pvMax += nouveauxNiveaux * 10
	c.pvAct += nouveauxNiveaux * 10
	fmt.Printf("Vous gagnez %d XP. Total : %d/%d XP.\n", experience, c.exp, c.level*100)
	fmt.Printf("Niveau supérieur ! Vous êtes maintenant niveau %d.\n", c.level)
	fmt.Printf("Vous gagnez %d PV max et vos dégâts augmentent de %d.\n", nouveauxNiveaux*10, nouveauxNiveaux*5)
}

func applyDamage(pvActuel *int, degats int) {
	*pvActuel -= degats
	if *pvActuel < 0 {
		*pvActuel = 0
	}
}

func monstrePattern(c *character, adversaire *Ennemi, tour int) {
	degats := adversaire.Degats

	switch {
	case adversaire.Nom == "Robot d'entraînement":
		monstreTurn(c, adversaire.Nom, degats, "Frappe de maintenance")
		return
	case adversaire.Boss:
		degats, nomAttaque := adamSmasherPattern(tour, adversaire.Degats)
		monstreTurn(c, adversaire.Nom, degats, nomAttaque)
		return
	case adversaire.Nom == "Cyberpsycho":
		degats, nomAttaque := cyberpsychoPattern(tour, adversaire.Degats)
		monstreTurn(c, adversaire.Nom, degats, nomAttaque)
		return
	case adversaire.Style != "":
		var perdSonTour bool
		var nomAttaque string
		degats, nomAttaque, perdSonTour = punkQuartierPattern(adversaire.Style, tour, adversaire.Degats)
		if perdSonTour {
			fmt.Println("Le monstre passe son tour.")
			return
		}
		monstreTurn(c, adversaire.Nom, degats, nomAttaque)
		return
	}

	monstreTurn(c, adversaire.Nom, degats, "Attaque basique")
}

func monsterPattern(c *character, adversaire *Ennemi, tour int) {
	monstrePattern(c, adversaire, tour)
}

func monstreTurn(c *character, nomAdversaire string, degats int, message string) {
	fmt.Println("")
	fmt.Println("========================================")
	fmt.Printf("%s utilise %s\n", nomAdversaire, message)
	fmt.Printf("Dégâts infligés : %d\n", degats)
	fmt.Println("========================================")
	applyDamage(&c.pvAct, degats)
	fmt.Printf("PV restants de %s : %d/%d\n", c.name, c.pvAct, c.pvMax)
}

func monsterTurn(c *character, nomAdversaire string, degats int, message string) {
	monstreTurn(c, nomAdversaire, degats, message)
}
