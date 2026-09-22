package main

import "fmt"

func displaySkill(c character) {
	fmt.Println("=== SORTS ===")
	if len(c.skill) == 0 {
		fmt.Println("Aucun sort appris.")
		return
	}

	for numero, sort := range c.skill {
		fmt.Printf("%d. %s\n", numero+1, sort)
		fmt.Println("=============")
	}
}
