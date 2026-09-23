package main

import "fmt"

func accessMerchant(c *character) {
	fmt.Println("\n=== MARCHAND ===")
	fmt.Println("Bienvenue chez le marchand !")
	if !c.firstMerchantVisit {
		if addInventory(c, "Stimulant", 1) {
			fmt.Println("Le marchand vous offre un stimulant gratuit pour votre première visite.")
		} else {
			fmt.Println("Votre inventaire est plein, le stimulant gratuit n'a pas pu être ajouté.")
		}
		c.firstMerchantVisit = true
	}
	marchand(c)
}
