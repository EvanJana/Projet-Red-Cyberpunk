package main

type character struct {
	name, class                   string
	skill                         []string
	level, exp, pvMax, pvAct, niv int
	combatsAleatoires             int
	vientDeMourir                 bool
	argent                        int
	inventaire                    map[string]int
	InventoryCapacity             int
	InventoryUpgradesCnt          int
	firstMerchantVisit            bool
	equipment                     equipement
}
type equipement struct {
	torse  string
	bottes string
	gants  string
}

func main() {
	c := characterCreation()
	c.pvAct = c.pvMax
	displayInfo(c)
	for menu(&c) {
		isDead(&c)
	}
}
