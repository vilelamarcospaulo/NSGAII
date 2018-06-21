package main

import (
	"NSGAII/nsgaii"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findOptimal() []nsgaii.Individual {
	optimal := []nsgaii.Individual{}
	for i := 0; i < 10; i++ {
		ag := nsgaii.NSGAII{}
		ag.Run(500, 1000, 500, .02, false)

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

	f, _ := os.Create("pareto.dat")
	buffer := ""
	for _, ind := range optimal {
		buffer += fmt.Sprintf("%f %f\n", ind.Goals[0], ind.Goals[1])
	}
	f.Write([]byte(buffer))

	return optimal
}

func readOptimal() []nsgaii.Individual {
	optimal := []nsgaii.Individual{}
	file, _ := os.Open("pareto.dat")

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		buffer := scanner.Text()
		values := strings.Split(buffer, " ")

		ind := nsgaii.Individual{}
		ind.NewRandom()

		for i := 0; i < ind.GoalsSize; i++ {
			ind.Goals[i], _ = strconv.ParseFloat(values[i], 64)
		}

		optimal = append(optimal, ind)
	}

	return optimal
}

func main() {
	// optimal := readOptimal()

	// agAux := nsgaii.NSGAII{}
	// agAux.Population = optimal
	// agAux.DoPlot()

	ag := nsgaii.NSGAII{}
	ag.Run(500, 1000, 200, 0.02, true)

	// ag.ParetoOptimal = optimal
	// fmt.Println("Error rate: ", ag.CalcErrorRate())
	// fmt.Println("Pareto subset: ", ag.CalcParetoSubset())
	// fmt.Println("Generational distance: ", ag.CalcGenerationalDistance())
	// fmt.Println("Spread : ", ag.CalcSpread())
	// fmt.Println("Maximum Spread (m3): ", ag.CalcMaximumSpread())

}
