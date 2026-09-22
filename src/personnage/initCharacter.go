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
		case "1", "Netrunner":
			class = "Netrunner"
		case "2", "assassin":
			class = "Assassin"
		case "3", "Berserk":
			class = "Berserk"
		default:
			fmt.Println("Classe invalide.")
			continue
		}
		break
	}

	return initCharacter(name, class)
}

func initCharacter(name, class string) character {
	pvMax := map[string]int{
		"Netrunner": 100,
		"Assassin":  75,
		"Berserk":   125,
	}[class]

	return character{
		name:       name,
		class:      class,
		skill:      []string{"Coup de Poing"},
		level:      1,
		pvMax:      pvMax,
		pvAct:      pvMax / 2,
		inventaire: make(map[string]int),
		InventoryCapacity:   	10,
		InventoryUpgradesCnt:    1,
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
