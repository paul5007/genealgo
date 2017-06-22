package main

import (
	"math/rand"
	"time"

	"fmt"

	g "github.com/paul5007/genealgo"
)

// box size is length of chromosome
const boxsize = 10

// values is number of values a gene has
const values = 2

// Box is chromosome
type Box []int

// NewBox creates a new box with random contents
func NewBox(rng *rand.Rand) g.Chromosome {
	b := Box{}
	for i := 0; i < boxsize; i++ {
		b = append(b, rng.Intn(values))
	}
	return b
}

// Crossover helps implement Chromosome interface
func (b Box) Crossover(c g.Chromosome, rng *rand.Rand) (g.Chromosome, g.Chromosome) {
	x := c.(Box)
	index := rng.Intn(len(b))
	b1 := Box{}
	b2 := Box{}
	for i := 0; i < index; i++ {
		b1 = append(b1, b[i])
		b2 = append(b2, x[i])
	}
	for i := index; i < len(b); i++ {
		b1 = append(b1, x[i])
		b2 = append(b2, b[i])
	}
	return b1, b2
}

// Mutate helps implement Chromosome interface
func (b Box) Mutate(rng *rand.Rand) {
	i := rng.Intn(len(b))
	if b[i] == 0 {
		b[i] = 1
	} else {
		b[i] = 0
	}
}

// Fitness helps implement Chromosome interface
func (b Box) Fitness() float64 {
	fitness := 0
	for _, each := range b {
		fitness += each
	}
	return float64(fitness)
}

func main() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	popSize := 10
	// create population 1
	pop := g.PopulationModel{}
	for i := 0; i < popSize; i++ {
		b := NewBox(rng)
		pop.Chromosomes = append(pop.Chromosomes, b)
		pop.FitScore += b.Fitness()
	}

	var trial int
	for pop.FitScore < float64(boxsize*popSize) {
		// create population 2
		pop2 := g.PopulationModel{}
		for len(pop2.Chromosomes) < len(pop.Chromosomes) {
			//select parents
			var found1, found2 bool
			b1 := Box{}
			for !found1 {
				for _, each := range pop.Chromosomes {
					wheel := rng.Intn(int(pop.FitScore))
					if float64(wheel) < each.Fitness() {
						b1 = each.(Box)
						found1 = true
						break
					}
				}
			}
			b2 := Box{}
			for !found2 {
				for _, each := range pop.Chromosomes {
					wheel := rng.Intn(int(pop.FitScore))
					if float64(wheel) < each.Fitness() {
						b2 = each.(Box)
						found2 = true
						break
					}
				}
			}
			// crossover parents
			x1, x2 := b1.Crossover(b2, rng)

			// try to mutate offspring
			const mut = 100
			mutChance := rng.Intn(mut)
			if mutChance == 1 {
				x1.Mutate(rng)
			}
			mutChance = rng.Intn(mut)
			if mutChance == 1 {
				x2.Mutate(rng)
			}

			pop2.Chromosomes = append(pop2.Chromosomes, x1)
			pop2.Chromosomes = append(pop2.Chromosomes, x2)
			pop2.FitScore += x1.Fitness()
			pop2.FitScore += x2.Fitness()
		}
		trial++
		fmt.Println("Trial:", trial)
		fmt.Println(pop)
		fmt.Println(pop2)

		pop = pop2
	}
}
