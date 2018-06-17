package main

import (
	"NSGAII/nsgaii"
	"fmt"
)

func main() {
	ag := nsgaii.NSGAII{}
	ag.Run(500, 1000, 50, 0.5)

	for _, ind := range ag.Population {
		fmt.Println(ind.Rank, "||", ind.DNA, "||", ind.Goals)
	}
}
