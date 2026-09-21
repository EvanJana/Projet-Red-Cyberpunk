package main

import "fmt"

func takePot(c *character) {
	quantite, disponible := c.inventaire["stimulant"]
	if !disponible || quantite <= 0 {
		fmt.Println("Vous n'avez plus de stimulant.")
		return
	}

	c.inventaire["stimulant"] = quantite - 1
	c.pvAct += 50
	if c.pvAct > c.pvMax {
		c.pvAct = c.pvMax
	}

	fmt.Printf("Stimulant utilisé. Points de vie : %d / %d\n", c.pvAct, c.pvMax)
}
