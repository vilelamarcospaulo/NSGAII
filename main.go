package main

import (
	"NSGAII/nsgaii"
	"fmt"
	"time"
)

func main() {
	// agAux := nsgaii.NSGAII{}
	// agAux.Population = optimal
	// agAux.DoPlot()

	ag := nsgaii.NSGAII{}
	start := time.Now()
	ag.Run(1000, 2000, 400, .02, true)
	elapsed := time.Since(start)

	fmt.Println("Time: ", elapsed)
	fmt.Println("Pareto size: ", ag.PopulationSize)
	fmt.Println("Error rate: ", ag.CalcErrorRate())
	fmt.Println("Pareto subset: ", ag.CalcParetoSubset())
	fmt.Println("Maximum Spread (m3): ", ag.CalcMaximumSpread())
}
