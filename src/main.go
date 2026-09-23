package main

type character struct {
	name, class              string
	skill                    []string
	level, pvMax, pvAct, niv int
	argent                   int
	inventaire               map[string]int
	InventoryCapacity        int
	InventoryUpgradesCnt     int
	equipment                equipement
}
type equipement struct {
	torse  string
	bottes string
	gants  string
}

func main() {
	c := characterCreation()
	displayInfo(c)
	for menu(&c) {
		isDead(&c)
	}
}
