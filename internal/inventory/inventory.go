package inventory

import "clicker/internal/generator"

type Inventory struct {
	Liquid     float64
	Generators generator.Generators
}

func (i *Inventory) Evaluate() {
	i.Liquid += i.Generators.Evaluate()
}

func (i *Inventory) Worth() float64 {
	return i.Liquid
}
