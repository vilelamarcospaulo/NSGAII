package main

import (
	"NSGAII/nsgaii"
	"fmt"
)

func main() {
	ag := nsgaii.NSGAII{}
	ag.Run(500, 1000, 800, .2)

	for _, ind := range ag.Population {
		fmt.Println(ind.Rank, "||", ind.DNA, "||", ind.Goals)
	}
}
