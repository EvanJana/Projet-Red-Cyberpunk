package main

import (
  "fmt"
  "math/rand"
  "time"
)

var argent int

type Ennemi struct {
  Nom   string
  Style  string
  PVMax  int
  PVActuel int
  Degats int
  Boss  bool
}

func monster(nombreRencontres int) Ennemi {
  if nombreRencontres > 0 && nombreRencontres%10 == 0 {
    return Ennemi{
      Nom:   "Adam Smasher",
      PVMax:  200,
      PVActuel: 200,
      Degats: 10,
      Boss:  true,
    }
  }

  ennemisNormaux := []Ennemi{
    {
      Nom:   "Cyberpsycho",
      PVMax:  100,
      PVActuel: 100,
      Degats: 5,
    },
    {
      Nom:   "Punk brutal de quartier",
      Style:  "brutal",
      PVMax:  120,
      PVActuel: 120,
      Degats: 8,
    },
    {
      Nom:   "Punk tireur de quartier",
      Style:  "tireur",
      PVMax:  55,
      PVActuel: 55,
      Degats: 12,
    },
    {
      Nom:   "Punk junkie de quartier",
      Style:  "junkie",
      PVMax:  70,
      PVActuel: 70,
      Degats: 6,
    },
  }

  choixAleatoire := rand.Intn(len(ennemisNormaux))
  return ennemisNormaux[choixAleatoire]
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
    return degats * 2, "Le Cyberpsycho utilise son attaque spéciale !"
  }

  return degats, "Le Cyberpsycho attaque !"
}

func punkQuartierPattern(style string, tour int, degats int) (int, string, bool) {
  switch style {
  case "brutal":
    return degats, "Le Punk brutal attaque au corps à corps !", false
  case "tireur":
    return degats, "Le Punk tireur ouvre le feu à distance !", false
  case "junkie":
    if rand.Intn(4) == 0 {
      return 0, "Le Punk junkie fait une overdose et perd son tour !", true
    }
    if tour%2 == 0 {
      return degats * 2, "Le Punk junkie attaque frénétiquement !", false
    }
    return degats, "Le Punk junkie attaque de façon imprévisible !", false
  default:
    return degats, "Le Punk de quartier attaque !", false
  }
}

func adamSmasherPattern(tour int, degats int) (int, string) {
  if tour > 0 && tour%5 == 0 {
    return degats + 20, "Adam Smasher utilise son attaque spéciale Skullcrusher !"
  }

  return degats, "Adam Smasher t'écrase !"
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
  rand.Seed(time.Now().UnixNano())

  nombreRencontres := 0
  adversaire := monster(nombreRencontres)
  fmt.Println("Adversaire :", adversaire.Nom)

  tourActuel := 1
  var degats int
  var messageAttaque string
  if adversaire.Boss {
    degats, messageAttaque = adamSmasherPattern(tourActuel, adversaire.Degats)
  } else if adversaire.Style != "" {
    degats, messageAttaque, _ = punkQuartierPattern(adversaire.Style, tourActuel, adversaire.Degats)
  } else {
    degats, messageAttaque = cyberpsychoPattern(tourActuel, adversaire.Degats)
  }
  fmt.Println(messageAttaque)

  tour(adversaire.Nom, degats, &adversaire.PVActuel, adversaire.PVMax, "Personnage")
}

