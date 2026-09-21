package main

import "fmt"

func isdead(c *character) {
	if c.pvAct != 0 {
		return
	}

	fmt.Println("VOUS ÊTES MORT")
	fmt.Println("1. Ressusciter")
	fmt.Println("2. Quitter le jeu")

	var choixAfterEnd string
	fmt.Scan(&choixAfterEnd)

	switch choixAfterEnd {
	case "1", "Ressusciter", "ressusciter":
		c.pvAct = c.pvMax / 2
	case "2", "Quitter", "quitter":
		fmt.Println("Au revoir !")
	}
}