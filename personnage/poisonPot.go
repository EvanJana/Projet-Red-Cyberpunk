package main

import (
    "fmt"
    "time"
)

func poisonPot(pvMonstre int) int {
    pvMonstre -= 10
    fmt.Println("Le monstre a été empoisonné !")
    time.Sleep(1 * time.Second)
    pvMonstre -= 10
    time.Sleep(1 * time.Second)
    pvMonstre -= 10
    return pvMonstre
}