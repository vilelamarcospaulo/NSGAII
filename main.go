package main

import (
	"NSGAII/nsgaii"
	"fmt"
)

func main() {
	optimal := []nsgaii.Individual{}
	for i := 0; i < 10; i++ {
		ag := nsgaii.NSGAII{}
		ag.Run(500, 1000, 200, .02, false)

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

		newOptimal := []nsgaii.Individual{}
		for _, ind := range optimal {
			isDominated := false
			for _, ind2 := range optimal {
				if ind2.Dominate(&ind) {
					isDominated = true
					break
				}
			}
			if !isDominated {
				newOptimal = append(newOptimal, ind)
			}
		}
		optimal = newOptimal

		fmt.Println("Execução: ", i, " //  ", len(optimal), " soluções nao dominadas")
	}

	ag := nsgaii.NSGAII{}
	ag.Run(500, 1000, 200, .02, false)

	ag.ParetoOptimal = optimal
	fmt.Println("Error rate: ", ag.CalcErrorRate())
	fmt.Println("Pareto subset: ", ag.CalcParetoSubset())
	fmt.Println("Generational distance: ", ag.CalcGenerationalDistance())
	fmt.Println("Spread (m3): ", ag.CalcSpread())
	fmt.Println("Maximum Spread (m3): ", ag.CalcMaximumSpread())

}
