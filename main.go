package main

import (
	"NSGAII/nsgaii"
)

func main() {
	optimal := []nsgaii.Individual{}
	for i := 0; i < 10; i++ {
		ag := nsgaii.NSGAII{}
		ag.Run(100, 1000, 1000, .02)

		for _, ind := range ag.Population {
			if ind.Rank != 0 {
				continue
			}

			appendInd := true

			for _, ind2 := range optimal {
				if ind2.Dominate(&ind) || ind.Equals(ind2) {
					appendInd = false
					break
				}
			}

			if appendInd {
				optimal = append(optimal, ind)
			}
		}
	}

}
