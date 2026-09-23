package main

import (
	"fmt"
	"os"
)

func isDead(c *character) {
	if c.pvAct > 0 {
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
		c.vientDeMourir = true
	case "2", "Quitter", "quitter":
		fmt.Println("Au revoir !")
		os.Exit(0)
	}
}
