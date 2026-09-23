package main

import (
	"fmt"
	"strings"
	"unicode"
)

func characterCreation() character {
	var name string
	for {
		fmt.Println("Entrez le nom de votre personnage (uniquement des lettres) :")
		fmt.Scan(&name)
		if isLettersOnly(name) {
			name = formatName(name)
			break
		}
		fmt.Println("Le nom doit contenir uniquement des lettres.")
	}

	var class string
	for {
		fmt.Println("Choisissez votre classe :")
		fmt.Println("1. Netrunner (100 PV max)")
		fmt.Println("2. Assassin (75 PV max)")
		fmt.Println("3. Berserk (125 PV max)")
		fmt.Scan(&class)

		switch strings.ToLower(class) {
		case "1", "netrunner":
			class = "Netrunner"
		case "2", "assassin":
			class = "Assassin"
		case "3", "berserk":
			class = "Berserk"
		default:
			fmt.Println("Classe invalide.")
			continue
		}
		break
	}

	c := initCharacter(name, class)
	switch class {
	case "Netrunner":
		c.initiative = 2
	case "Assassin":
		c.initiative = 2
	case "Berserk":
		c.initiative = 2
	}
	return c
}

func initCharacter(name, class string) character {
	pvMax := map[string]int{
		"Netrunner": 100,
		"Assassin":  75,
		"Berserk":   125,
	}[class]
	ramMax := map[string]int{
		"Netrunner": 125,
		"Assassin":  75,
		"Berserk":   50,
	}[class]

	initiative := 1
	switch class {
	case "Netrunner":
		initiative = 2
	case "Assassin":
		initiative = 2
	case "Berserk":
		initiative = 2
	}

	return character{
		name:                 name,
		class:                class,
		skill:                []string{"Coup de Poing"},
		level:                1,
		exp:                  0,
		pvMax:                pvMax,
		pvAct:                pvMax / 2,
		ramMax:               ramMax,
		ramAct:               ramMax,
		argent:               100,
		inventaire:           make(map[string]int),
		InventoryCapacity:    10,
		InventoryUpgradesCnt: 0,
		firstMerchantVisit:   false,
		initiative:           initiative,
	}
}

func isLettersOnly(value string) bool {
	if value == "" {
		return false
	}
	for _, character := range value {
		if !unicode.IsLetter(character) {
			return false
		}
	}
	return true
}

func formatName(name string) string {
	name = strings.ToLower(name)
	runes := []rune(name)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
