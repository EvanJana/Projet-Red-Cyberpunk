package main

import (
	"fmt"
	"math/rand"
)

type Ennemi struct {
	Nom      string
	Style    string
	Drop     string
	PVMax    int
	PVActuel int
	Degats   int
	Boss     bool
}

func dropAleatoire() string {
	choix := rand.Intn(100)
	switch {
	case choix < 30:
		return "Acier"
	case choix < 60:
		return "Kevlar"
	case choix < 80:
		return "Cuir"
	default:
		return "Puce neuronale"
	}
}

func robotEntrainement() Ennemi {
	return Ennemi{
		Nom:      "Robot d'entraînement",
		Drop:     "dollars",
		PVMax:    40,
		PVActuel: 40,
		Degats:   5,
	}
}

func monstre(nombreRencontres int) Ennemi {
	if nombreRencontres > 0 && nombreRencontres%10 == 0 {
		return Ennemi{
			Nom:      "Adam Smasher",
			Drop:     dropAleatoire(),
			PVMax:    200,
			PVActuel: 200,
			Degats:   10,
			Boss:     true,
		}
	}

	ennemisNormaux := []Ennemi{
		{
			Nom:      "Cyberpsycho",
			PVMax:    100,
			PVActuel: 100,
			Degats:   5,
		},
		{
			Nom:      "Punk brutal de quartier",
			Style:    "brutal",
			PVMax:    120,
			PVActuel: 120,
			Degats:   8,
		},
		{
			Nom:      "Punk tireur de quartier",
			Style:    "tireur",
			PVMax:    55,
			PVActuel: 55,
			Degats:   12,
		},
		{
			Nom:      "Punk junkie de quartier",
			Style:    "junkie",
			PVMax:    70,
			PVActuel: 70,
			Degats:   6,
		},
	}

	choixAleatoire := rand.Intn(len(ennemisNormaux))
	monstreAleatoire := ennemisNormaux[choixAleatoire]
	monstreAleatoire.Drop = dropAleatoire()
	return monstreAleatoire
}

func monster(nombreRencontres int) Ennemi {
	return monstre(nombreRencontres)
}
func initAdamSmasher() {
	nomMonstre := "Adam Smasher"
	pvMonstre := 200
	pvActuel := pvMonstre
	DmgMonstre := 10
	_, _, _, _ = nomMonstre, pvMonstre, pvActuel, DmgMonstre
}
func initCyberpsycho() {
	nomMonstre := "Cyberpsycho "
	pvMonstre := 100
	pvActuel := pvMonstre
	DmgMonstre := 5
	_, _, _, _ = nomMonstre, pvMonstre, pvActuel, DmgMonstre
}
func cyberpsychoPattern(tour int, degats int) (int, string) {
	if tour > 0 && tour%3 == 0 {
		return degats * 2, "Attaque spéciale"
	}

	return degats, "Attaque basique"
}
func punkQuartierPattern(style string, tour int, degats int) (int, string, bool) {
	switch style {
	case "brutal":
		return degats, "Attaque au corps à corps", false
	case "tireur":
		return degats, "Tir à distance", false
	case "junkie":
		if rand.Intn(4) == 0 {
			return 0, "Overdose", true
		}
		if tour%2 == 0 {
			return degats * 2, "Attaque frénétique", false
		}
		return degats, "Attaque imprévisible", false
	default:
		return degats, "Attaque basique", false
	}
}
func adamSmasherPattern(tour int, degats int) (int, string) {
	if tour > 0 && tour%5 == 0 {
		return degats + 20, "Skullcrusher"
	}

	return degats, "Écrasement"
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
