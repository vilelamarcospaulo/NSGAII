package main

import (
	"NSGAII/nsgaii"
	"fmt"
)

func main() {
	ag := nsgaii.NSGAII{}
	ag.Run(500, 500, 500, .2)

	for _, ind := range ag.Population {
		fmt.Println(ind.Rank, "||", ind.DNA, "||", ind.Goals)
	}
}
