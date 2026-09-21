package main

import "fmt"

func accessMerchant(c *character) {
	fmt.Println("\n=== MARCHAND ===")
	fmt.Println("Bienvenue chez le marchand !")
	marchand(c.inventaire)
}
