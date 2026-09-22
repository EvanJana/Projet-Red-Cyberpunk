package main

import "fmt"

var choixForge = map[int]string{
	1: "Armure de combat",
	2: "Bottes de soldat",
	3: "Gants de précision",
}

var recettesForge = map[string]map[string]int{
	"Armure de combat":   {"Kevlar": 1, "Acier": 2},
	"Bottes de soldat":   {"Cuir": 2, "Acier": 1},
	"Gants de précision": {"Cuir": 1, "Puce neuronale": 2},
}

func ferailleur(inventaire map[string]int) bool {
	fmt.Println("================================")
	fmt.Println("Bienvenue dans ma forge ! \nQue voulez-vous fabriquer ?")
	fmt.Println(choixForge)
	var nb int
	fmt.Scan(&nb)
	if nb < 1 || nb > 3 {
		fmt.Println("Je ne sais pas fabriquer ça, désolé.")
		return false
	}
	objet := choixForge[nb]
	recette, existe := recettesForge[objet]
	if existe == false {
		fmt.Println("Cette recette n'est pas disponible.")
		return false
	}

	materiauxManquants := false
	for materiau, quantiteRequise := range recette {
		if inventaire[materiau] < quantiteRequise {
			materiauxManquants = true
		}
	}

	if materiauxManquants == true {
		fmt.Println("Il te manque des matériaux pour fabriquer cet objet")
		for materiau, quantiteRequise := range recette {
			fmt.Printf(" - %s : %d requis, %d possédé(s)\n", materiau, quantiteRequise, inventaire[materiau])
		}
		return false
	} else if limiteinv(inventaire) {
		fmt.Println("Tu ne peux pas fabriquer plus d'objets, ton inventaire est plein.")
		return false
	} else {
		for materiau, quantiteRequise := range recette {
			inventaire[materiau] -= quantiteRequise
		}
		inventaire[objet] += 1
		fmt.Printf("Fabrication réussie : %s\n", objet)
		fmt.Println("================================")
		return true
	}
}
