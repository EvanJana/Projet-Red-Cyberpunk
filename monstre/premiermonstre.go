package main

import "fmt"

func monster() {
	nomMonstre := "Cyberpsycho"
	pvMonstre := 100
	pvActuel := pvMonstre
	DmgMonstre := 5
	_, _, _, _ = nomMonstre, pvMonstre, pvActuel, DmgMonstre
}
func initCyberpsycho() {
	nomMonstre := "Cyberpsycho d'entrainement"
	pvMonstre := 40
	pvActuel := pvMonstre
	DmgMonstre := 5
	_, _, _, _ = nomMonstre, pvMonstre, pvActuel, DmgMonstre
}

func cyberpsychoPattern(tour int, degats int) (int, string) {
	if tour > 0 && tour%3 == 0 {
		return degats * 2, "Le Cyberpsycho utilise son attaque spéciale !"
	}

	return degats, "Le Cyberpsycho attaque !"
}
func tour(nomAttaquant string, degats int, pvActuel *int, pvMax int, nomCible string) (int, string) {
	if *pvActuel <= 0 {
		return 0, nomCible + " est déjà KO !"
	}
	*pvActuel -= degats
	if *pvActuel < 0 {
		*pvActuel = 0
	}
	messageattaque := (nomAttaquant + " attaque " + nomCible + " et fait" + fmt.Sprint(degats) + " dégâts !")
	messagePV := nomCible + " a maintenant " + fmt.Sprint(*pvActuel) + "/" + fmt.Sprint(pvMax) + " PV"
	fmt.Println(messageattaque)
	fmt.Println(messagePV)
	return *pvActuel, messagePV
}

func main() {
	tourActuel := 1
	degatsBase := 5
	degats, messageAttaque := cyberpsychoPattern(tourActuel, degatsBase)
	fmt.Println(messageAttaque)

	pvMax := 100
	pvActuel := pvMax
	tour("Cyberpsycho", degats, &pvActuel, pvMax, "Personnage")
}
