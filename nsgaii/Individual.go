package nsgaii

import (
	"math"
	"math/rand"
)

//Individual :: Representacao do individuo do AG
type Individual struct {
	DNA     *([]float64)
	DNASize int

	CurrentGoal int
	Goals       []float64
	GoalsSize   int
	gX          float64

	Rank             int
	CrowdingDistance float64

	//Usado para calcular GenerationalDistance
	Distance float64
}

func randomValidValue() float64 {
	return random(0, 1)
}

//NewRandom :: Gera um individuo
func (individual *Individual) NewRandom() {
	individual.DNASize, individual.GoalsSize = 30, 2 //5, 2

	dna := make([]float64, individual.DNASize)

	individual.DNA = &dna
	individual.Goals = make([]float64, individual.GoalsSize)

	for i := 0; i < individual.GoalsSize; i++ {
		individual.Goals[i] = 0.0
	}

	for i := 0; i < individual.DNASize; i++ {
		(*individual.DNA)[i] = randomValidValue()
	}

	individual.Rank = -1
	individual.CrowdingDistance = -1
}

//Eval :: Calcula os valores para cada objetivo com base no DNA
//deve mudar para cada problema.
func (individual *Individual) Eval() {
	for i := 0; i < individual.GoalsSize; i++ {
		individual.Goals[i] = 0.0
	}

	individual.Goals[0] = (*individual.DNA)[0]
	individual.gX = 0.0

	for i := 1; i < individual.DNASize; i++ {
		individual.gX += (*individual.DNA)[i]
		// individual.Goals[0] += math.Sin(math.Pi * (*individual.DNA)[i])
		// individual.Goals[1] += math.Cos(math.Pi * (*individual.DNA)[i])
	}
	individual.gX *= 9.0 / (float64(individual.DNASize) - 1)
	individual.gX += 1

	individual.Goals[1] =
		individual.gX *
			(1 -
				math.Sqrt(individual.Goals[0]/individual.gX) -
				((individual.Goals[0] / individual.gX) * math.Sin(10*math.Pi*individual.Goals[0])))

	individual.Rank = -1
	individual.CrowdingDistance = -1
}

func random(min float64, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

//Mutation :: Realiza uma mutação no individuo
//se o individuo sofrer mutação, sorteia uma posição e altera o valor dela
//por um outro valor no dominio valido
func (individual *Individual) Mutation(probability float64) {
	for i := 0; i < individual.DNASize; i++ {
		if rand.Float64() > probability {
			continue
		}
		sum := random(-0.5, 0.5)
		(*individual.DNA)[i] += sum
		if (*individual.DNA)[i] > 1 {
			(*individual.DNA)[i] = 1
		} else if (*individual.DNA)[i] < 0 {
			(*individual.DNA)[i] = 0
		}
	}
}

//Dominate :: Checa se um individuo é dominado por outro
// se for menor em ao menos um objeto é considerado falso
// se for maior em ao menos um e indiferente aos outros é true
func (individual *Individual) Dominate(other *Individual) bool {
	hasOneLessThan := false
	for i := 0; i < individual.GoalsSize; i++ {
		if individual.Goals[i] > other.Goals[i] {
			return false
		}
		if individual.Goals[i] < other.Goals[i] {
			hasOneLessThan = true
		}
	}
	return hasOneLessThan
}

//Better :: Checa se um individuo é melhor que o outro baseado no rank e crowding distance
func (individual *Individual) Better(other Individual) bool {
	return individual.Rank < other.Rank ||
		(individual.Rank == other.Rank && individual.CrowdingDistance > other.CrowdingDistance)
}

//Crossover :: Realiza o crosoover entre dois pais "parent1" e "parent2"
//e coloca o resultado em "child1" e "child2"
func Crossover(parent1 Individual, parent2 Individual, child1 *Individual, child2 *Individual) {
	for i := 0; i < parent1.DNASize; i++ {
		if rand.Intn(2) == 0 {
			(*child1.DNA)[i], (*child2.DNA)[i] = (*parent1.DNA)[i], (*parent2.DNA)[i]
		} else {
			(*child1.DNA)[i], (*child2.DNA)[i] = (*parent2.DNA)[i], (*parent1.DNA)[i]
		}
	}

	// middle := int(math.Ceil(float64(parent1.DNASize) / 2.0))
	// for i := 0; i < middle; i++ {
	// 	(*child1.DNA)[i], (*child2.DNA)[i] = (*parent1.DNA)[i], (*parent2.DNA)[i]
	// }

	// for i := middle; i < parent1.DNASize; i++ {
	// 	(*child1.DNA)[i], (*child2.DNA)[i] = (*parent2.DNA)[i], (*parent1.DNA)[i]
	// }

	child1.Rank = -1
	child2.Rank = -1

	child1.CrowdingDistance = -1
	child2.CrowdingDistance = -1

	for i := 0; i < parent1.GoalsSize; i++ {
		child1.Goals[i], child2.Goals[i] = 0, 0
	}
}

//GoalsDistance :: Calcula a soma da distancia euclidiana de cada objetivo, entre os dois individuos
func (individual *Individual) GoalsDistance(other Individual) float64 {
	individual.Distance = 0

	for i := 0; i < individual.GoalsSize; i++ {
		individual.Distance += math.Pow(individual.Goals[i]-other.Goals[i], 2.0)
	}

	return individual.Distance
}

//Equals :: verifica se tme o mesmo DNA
func (individual Individual) Equals(other Individual) bool {
	for i := 0; i < individual.DNASize; i++ {
		if (*individual.DNA)[i] != (*other.DNA)[i] {
			return false
		}
	}
	return true
}
