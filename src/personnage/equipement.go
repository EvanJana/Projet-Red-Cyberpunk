package main

var slotParObjet = map[string]string{
    "Armure de combat":   "torse",
    "Bottes de soldat":   "bottes",
    "Gants de précision": "gants",
}

func equiper(inventaire map[string]int, equipement *equipement) bool {
    for objet, slot := range slotParObjet {
        if inventaire[objet] > 0 {
            switch slot {
			case "torse":
                equipement.torse = objet
            case "bottes":
                equipement.bottes = objet
            case "gants":
                equipement.gants = objet
            }
            removeInventory(inventaire, objet)
            return true
        }
    }
    return false
} 