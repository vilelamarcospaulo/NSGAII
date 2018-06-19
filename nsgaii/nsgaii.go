package nsgaii

import (
	"fmt"
	"math"
	"math/rand"
	"sort"

	"github.com/Arafatk/glot"
)

//NSGAII :: Representacao da estrutura do AG do tipo NSGAII
type NSGAII struct {
	Population []Individual

	PopulationSize int
	ChildSize      int
	VectorSize     int

	Generation          int
	MutationProbability float64

	plot *glot.Plot
}

//Run :: inicializa a configuração e processa o ag
func (nsgaii *NSGAII) Run(Generations int, PopulationSize int, ChildSize int, MutationProbability float64) {
	nsgaii.Generation = 0
	nsgaii.PopulationSize = PopulationSize
	nsgaii.ChildSize = ChildSize
	nsgaii.MutationProbability = MutationProbability

	nsgaii.plot, _ = glot.NewPlot(2, true, true)

	nsgaii.newPopulation()
	for nsgaii.Generation = 0; nsgaii.Generation <= Generations; nsgaii.Generation++ {
		nsgaii.nextPopulation()
		nsgaii.doPlot()
	}
}

//Plot :: Plota a populacao atual
func (nsgaii *NSGAII) doPlot() {
	xaxis := make([]float64, 1)
	yaxis := make([]float64, 1)
	for i := 0; i < nsgaii.PopulationSize; i++ {
		xaxis = append(xaxis, nsgaii.Population[i].Goals[0])
		yaxis = append(yaxis, nsgaii.Population[i].Goals[1])
	}

	points := [][]float64{xaxis, yaxis}
	nsgaii.plot.AddPointGroup(" ", "points", points)

	title := fmt.Sprintf("%s%d", "Generation: ", nsgaii.Generation)
	nsgaii.plot.SetTitle(title)

	nsgaii.plot.SetXLabel("SUM(sin(Pi * N))")
	nsgaii.plot.SetYLabel("SUM(sin(Pi * N))")

	// nsgaii.plot.SetYrange(-5, 5)
	// nsgaii.plot.SetXrange(-5, 5)
	// if nsgaii.Generation%5 == 0 && nsgaii.Generation < 10 {
	// 	file := fmt.Sprintf("%s%d%s", "plots/gen", nsgaii.Generation, ".png")
	// 	err := nsgaii.plot.SavePlot(file)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// } else if nsgaii.Generation%10 == 0 && nsgaii.Generation < 10 {
	// 	file := fmt.Sprintf("%s%d%s", "plots/gen", nsgaii.Generation, ".png")
	// 	err := nsgaii.plot.SavePlot(file)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// } else if nsgaii.Generation%100 == 0 && nsgaii.Generation < 100 {
	// 	file := fmt.Sprintf("%s%d%s", "plots/gen", nsgaii.Generation, ".png")
	// 	err := nsgaii.plot.SavePlot(file)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	nsgaii.plot.ResetPlot()
}

//NewPopulation :: Cria uma população inicial aleatoria
func (nsgaii *NSGAII) newPopulation() {
	nsgaii.VectorSize = nsgaii.PopulationSize + nsgaii.ChildSize
	nsgaii.Population = make([]Individual, nsgaii.VectorSize)

	for i := 0; i < nsgaii.VectorSize; i++ {
		nsgaii.Population[i].NewRandom()
		nsgaii.Population[i].Eval()
	}
	nsgaii.rank()
}

func (nsgaii NSGAII) selectParentByTour() (int, Individual) {
	index := rand.Intn(nsgaii.PopulationSize)
	individual := nsgaii.Population[index]
	for i := 1; i < 3; i++ {
		if newIndex := rand.Intn(nsgaii.PopulationSize); nsgaii.Population[newIndex].Better(individual) {
			index = newIndex
			individual = nsgaii.Population[index]
		}
	}
	return index, individual
}

//NextPopulation :: Gera a população t + 1, com base na atual (t)
func (nsgaii *NSGAII) nextPopulation() {
	for i := nsgaii.PopulationSize; i < nsgaii.VectorSize; i += 2 {
		child1 := &nsgaii.Population[i]
		child2 := &nsgaii.Population[i+1]

		indexParent1, parent1 := nsgaii.selectParentByTour()
		indexParent2, parent2 := nsgaii.selectParentByTour()

		for indexParent1 == indexParent2 {
			indexParent2, parent2 = nsgaii.selectParentByTour()
		}

		Crossover(parent1, parent2, child1, child2)

		child1.Mutation(nsgaii.MutationProbability)
		child2.Mutation(nsgaii.MutationProbability)

		//Avalia os filhos gerados de acordo com o novo DNA
		child1.Eval()
		child2.Eval()
	}

	nsgaii.crowdingDistance(nsgaii.rank())
	nsgaii.reinsert()
}

//Rank :: Ranqueia os individuos da população de acordo com a nao dominancia
// ao qual os mesmos pertencem, no rank em que extrapolar o tamanho da populacao
// para de rankear, já que a populacao está completa e retorna esse rank para
//
func (nsgaii *NSGAII) rank() int {
	rankedIndividuals := 0

	var unranked [](*Individual)
	for i := 0; i < nsgaii.VectorSize; i++ {
		nsgaii.Population[i].Rank = 100000
		unranked = append(unranked, &nsgaii.Population[i])
	}

	rank := 0

	for rank = 0; rankedIndividuals < nsgaii.PopulationSize; rank++ {
		for cIdx := 0; cIdx < len(unranked); cIdx++ {
			isNotDominated := true
			for _, otherIndividual := range unranked {
				if otherIndividual.Dominate(unranked[cIdx]) {
					isNotDominated = false
					break
				}
			}

			if isNotDominated {
				unranked[cIdx].Rank = rank
				rankedIndividuals++
				//Remove o individo da listan de unranked
				unranked = append(unranked[:cIdx], unranked[cIdx+1:]...)
			}
		}
	}

	if rankedIndividuals > nsgaii.PopulationSize {
		return rank - 1
	}

	return -1
}

//Reinsert :: Reordena a populacao completa com base no rank e no crowding distance
func (nsgaii *NSGAII) reinsert() {
	sort.Sort(ByRankAndCrowdingDistance(nsgaii.Population))
}

//CrowdingDistance :: CalcUla a crowding distance para os elementos do rank em questao
func (nsgaii *NSGAII) crowdingDistance(rank int) {
	if rank == -1 {
		return
	}

	var onRank [](*Individual)
	for i := 0; i < nsgaii.VectorSize; i++ {
		if nsgaii.Population[i].Rank == rank {
			nsgaii.Population[i].CrowdingDistance = 0
			onRank = append(onRank, &nsgaii.Population[i])
		}
	}
	size := len(onRank)

	for goal := 0; goal < nsgaii.Population[0].GoalsSize; goal++ {
		for i := 0; i < size; i++ {
			onRank[i].CurrentGoal = goal
		}

		sort.Sort(ByGoal(onRank))

		onRank[0].CrowdingDistance += 10000
		onRank[size-1].CrowdingDistance += 10000

		for index := 1; index < size-1; index++ {
			goalAverage := math.Abs(onRank[index-1].Goals[goal] - onRank[index+1].Goals[goal])
			onRank[index].CrowdingDistance += goalAverage
		}
	}
}
