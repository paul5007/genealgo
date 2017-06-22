package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var (
	rng *rand.Rand
)

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func TestNewBox(t *testing.T) {
	b := NewBox(rng)
	if len(b.(Box)) != boxsize {
		t.FailNow()
	}
	t.Log(b)
}
func TestCrossover(t *testing.T) {
	p1 := NewBox(rng)
	p2 := NewBox(rng)
	t.Log(p1.(Box))
	t.Log(p2.(Box))
	o1, o2 := p1.Crossover(p2, rng)
	var found bool
	for i := range o1.(Box) {
		if o1.(Box)[i] != o2.(Box)[i] {
			t.Log(o1.(Box))
			t.Log(o2.(Box))
			found = true
			break
		}
	}
	if !found {
		t.Fail()
	}
}

func TestFitness(t *testing.T) {
	b := Box{1, 1, 1, 1, 1}
	fit := b.Fitness()
	fmt.Println(fit)
	if fit != 5.00 {
		t.Fail()
	}
}

func TestMutate(t *testing.T) {
	b := NewBox(rng)
	var sum1, sum2 int
	for _, each := range b.(Box) {
		sum1 += each
	}
	t.Log(b.(Box))
	b.Mutate(rng)
	t.Log(b.(Box))
	for _, each := range b.(Box) {
		sum2 += each
	}
	if sum1 == sum2 {
		t.Fail()
	}
}
