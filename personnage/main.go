package main

type character struct {
	name, class         string
	level, pvMax, pvAct int
	inventaire          map[string]int
}

func main() {
	c := initCharacter()
	displayInfo(c)

	for menu(&c) {
		isdead(&c)
	}
}
