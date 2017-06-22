package genealgo

import "math/rand"

// Chromosome takes the form of an individual
type Chromosome interface {
	Crossover(c Chromosome, rng *rand.Rand) (Chromosome, Chromosome)
	Mutate(rng *rand.Rand)
	Fitness() float64
}

// ChromosomeModel holds metadata for a particular chromosome
type ChromosomeModel struct {
	Chromosome Chromosome
	FitScore   float64 // individual fitscore
}

// ChromosomeFactory is a method that generates a new Chromosome
type ChromosomeFactory func(rng *rand.Rand) Chromosome

// Population is a group of Chromosomes
type Population interface{}

// PopulationModel holds metadata for a particular population
type PopulationModel struct {
	Chromosomes []Chromosome
	FitScore    float64 // total fitscore
}
