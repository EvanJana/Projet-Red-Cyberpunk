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
		fmt.Println("1. Humain (100 PV max)")
		fmt.Println("2. Elfe (80 PV max)")
		fmt.Println("3. Nain (120 PV max)")
		fmt.Scan(&class)

		switch strings.ToLower(class) {
		case "1", "humain":
			class = "Humain"
		case "2", "elfe":
			class = "Elfe"
		case "3", "nain":
			class = "Nain"
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
		"Humain": 100,
		"Elfe":   80,
		"Nain":   120,
	}[class]

	return character{
		name:       name,
		class:      class,
		skill:      "Coup de Poing",
		level:      1,
		pvMax:      pvMax,
		pvAct:      pvMax / 2,
		inventaire: make(map[string]int),
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
