package main

import "fmt"

func takePot(c *character) {
	quantite, disponible := c.inventaire["Stimulant"]
	if !disponible || quantite <= 0 {
		fmt.Println("Vous n'avez plus de stimulant.")
		return
	}

	removeInventory(c.inventaire, "Stimulant")
	c.pvAct += 50
	if c.pvAct > c.pvMax {
		c.pvAct = c.pvMax
	}

	fmt.Printf("Stimulant utilisé. Points de vie : %d / %d\n", c.pvAct, c.pvMax)
}
