package main

import "fmt"

func displaySkill(c character) {
	fmt.Println("=== Compétences ===")
	if len(c.skill) == 0 {
		fmt.Println("Aucune compétences apprise.")
		return
	}

	for numero, sort := range c.skill {
		fmt.Printf("%d. %s\n", numero+1, sort)
		fmt.Println("=============")
	}
}
