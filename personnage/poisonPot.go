package main

import (
	"fmt"
	"time"
)

func poisonPot(pvMonstre int) int {
	degat := 10
	pvMonstre -= degat
	fmt.Printf("L'ennemie a subi %d dégats\n", degat)
	fmt.Println("Le monstre a été empoisonné !")
	time.Sleep(1 * time.Second)
	pvMonstre -= degat
    fmt.Printf("L'ennemie a subi %d dégats\n", degat)
	time.Sleep(1 * time.Second)
	pvMonstre -= degat
    fmt.Printf("L'ennemie a subi %d dégats\n", degat)
	return pvMonstre
}
