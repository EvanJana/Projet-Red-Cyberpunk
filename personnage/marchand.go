package main

import (
	"fmt"
	"sort"
)

var argent = 100000

var choix = map[int]string{
	1:  "Stimulant",
	2:  "Surchage",
	3:  "Crash",
	4:  "Suicide",
	5:  "IEM",
	6:  "Armure de combat",
	7:  "Bottes de soldat",
	8:  "Gants de précision",
	9: "Potion de poison",
}

var objects = map[string]int{
	"Stimulant":          30,
	"Surchage":           100,
	"Crash":              110,
	"Suicide":            400,
	"IEM":                50,
	"Armure de combat":   150,
	"Bottes de soldat":   80,
	"Gants de précision": 60,
	"Potion de poison":   40,
}

func marchand(inventaire map[string]int) {
	fmt.Printf("Argent disponible : $%d\n", argent)

	numeros := make([]int, 0, len(choix))
	for numero := range choix {
		numeros = append(numeros, numero)
	}
	sort.Ints(numeros)

	for _, numero := range numeros {
		objet := choix[numero]
		fmt.Printf("%d. %s - $%d\n", numero, objet, objects[objet])
	}
	fmt.Println("11. Retour")

	var choixAcheteur int
	fmt.Scan(&choixAcheteur)

	if choixAcheteur == 10 {
		return
	}

	objet, existe := choix[choixAcheteur]
	if !existe {
		fmt.Println("Choix invalide.")
		return
	}

	prix := objects[objet]
	if argent < prix {
		fmt.Println("Tu n'as pas assez d'argent.")
		return
	}
	if limiteinv(inventaire) {
		return
	}

	argent -= prix
	addInventory(inventaire, objet)
	fmt.Printf("Achat effectué : %s\n", objet)
	fmt.Printf("Il vous reste $%d\n", argent)
}

func addInventory(inventaire map[string]int, objet string) {
	if limiteinv(inventaire) == true {
		fmt.Println("Inventaire plein")
	} else {
		inventaire[objet]++
	}
}

func removeInventory(inventaire map[string]int, objet string) {
	quantite, existe := inventaire[objet]
	if !existe {
		return
	}

	if quantite <= 1 {
		delete(inventaire, objet)
		return
	}

	inventaire[objet] = quantite - 1
}
