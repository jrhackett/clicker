package generators

import (
	"fmt"
	"math"
	"time"
)

type (
	Generator struct {
		Name          GeneratorName `json:"name"`
		Count         int64         `json:"count"`
		Cost          float64       `json:"cost"`
		GainPerSecond float64       `json:"gain_per_second"`
		Gained        float64       `json:"gained"`
		LastEvaluated time.Time     `json:"last_evaluated"`
	}

	Generators []*Generator
)

func (g *Generator) Evaluate() float64 {
	now := time.Now()
	since := now.Sub(g.LastEvaluated)
	gained := float64(g.Count) * g.GainPerSecond * since.Seconds()

	g.Gained += gained
	g.LastEvaluated = now

	return gained
}

func (gs *Generators) Evaluate() float64 {
	var earned float64
	for _, g := range *gs {
		earned += g.Evaluate()
	}
	return earned
}

func (gs *Generators) Add(name string) error {
	i, err := gs.generatorIndex(name)
	if err != nil {
		return err
	}

	g := (*gs)[i]

	g.Count++
	g.Cost = math.Pow(math.E, float64(g.Count)*0.8)

	return nil
}

func (gs *Generators) Cost(name string) (float64, error) {
	i, err := gs.generatorIndex(name)
	if err != nil {
		return 0, err
	}

	return (*gs)[i].Cost, nil
}

func (gs *Generators) generatorIndex(name string) (int, error) {
	generator := GeneratorName(name)

	for i, g := range *gs {
		if g.Name == generator {
			return i, nil
		}
	}

	return 0, GeneratorNameDoesNotExist{
		Name: name,
		Err:  fmt.Errorf("generator does not exist"),
	}
}
