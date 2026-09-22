package main

var slotParObjet = map[string]string{
    "Armure de combat":   "torse",
    "Bottes de soldat":   "bottes",
    "Gants de précision": "gants",
}

var bonusPVParObjet = map[string]int{
    "Armure de combat":   25,
    "Bottes de soldat":   15,
    "Gants de précision": 10,
}

func equiper(c *character, objet string) bool {
    slot, existe := slotParObjet[objet]
    if !existe || c.inventaire[objet] <= 0 {
        return false
    }

    ancienObjet := equipementDansSlot(c.equipment, slot)
    removeInventory(c.inventaire, objet)
    if ancienObjet != "" {
        c.inventaire[ancienObjet]++
    }
    definirEquipementDansSlot(&c.equipment, slot, objet)

    c.pvMax += bonusPVParObjet[objet] - bonusPVParObjet[ancienObjet]
    if c.pvAct > c.pvMax {
        c.pvAct = c.pvMax
    }
    return true
}

func equipementDansSlot(equipement equipement, slot string) string {
    switch slot {
    case "torse":
        return equipement.torse
    case "bottes":
        return equipement.bottes
    case "gants":
        return equipement.gants
    default:
        return ""
    }
}

func definirEquipementDansSlot(equipement *equipement, slot, objet string) {
    switch slot {
    case "torse":
        equipement.torse = objet
    case "bottes":
        equipement.bottes = objet
    case "gants":
        equipement.gants = objet
    }
}