package main

type character struct {
	name, class              string
	skill                    []string
	level, pvMax, pvAct, niv int
	inventaire               map[string]int
}

func main() {
	c := characterCreation()
	displayInfo(c)

	for menu(&c) {
		isDead(&c)
	}
}
