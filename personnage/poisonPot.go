package main

import (
	"fmt"
	"time"
)

func poisonPot(pvMonstre int) int {
	const degat = 10

	fmt.Println("Le monstre a été empoisonné !")
	for seconde := 1; seconde <= 3; seconde++ {
		pvMonstre -= degat
		if pvMonstre < 0 {
			pvMonstre = 0
		}
		fmt.Printf("Seconde %d : le monstre a subi %d dégâts (%d PV restants)\n", seconde, degat, pvMonstre)
		if seconde < 3 {
			time.Sleep(1 * time.Second)
		}
	}

	return pvMonstre
}
