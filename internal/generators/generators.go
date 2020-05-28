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
		InitialCost   float64       `json:"-"`
		GainPerSecond float64       `json:"gain_per_second"`
		Gained        float64       `json:"gained"`
		RateOfGrowth  float64       `json:"-"`
		LastEvaluated time.Time     `json:"last_evaluated"`
	}

	Generators []*Generator
)

func New() Generators {
	now := time.Now()
	return Generators{
		&Generator{
			Name:          A,
			Count:         0,
			InitialCost:   5,
			GainPerSecond: 2,
			Gained:        0,
			RateOfGrowth:  0.3,
			LastEvaluated: now,
		},
		&Generator{
			Name:          B,
			Count:         0,
			InitialCost:   40,
			GainPerSecond: 5,
			Gained:        0,
			RateOfGrowth:  0.4,
			LastEvaluated: now,
		},
		&Generator{
			Name:          C,
			Count:         0,
			InitialCost:   100,
			GainPerSecond: 10,
			Gained:        0,
			RateOfGrowth:  0.6,
			LastEvaluated: now,
		},
		&Generator{
			Name:          D,
			Count:         0,
			InitialCost:   250,
			GainPerSecond: 22,
			Gained:        0,
			RateOfGrowth:  0.5,
			LastEvaluated: now,
		},
		&Generator{
			Name:          E,
			Count:         0,
			InitialCost:   400,
			GainPerSecond: 45,
			Gained:        0,
			RateOfGrowth:  0.7,
			LastEvaluated: now,
		},
		&Generator{
			Name:          F,
			Count:         0,
			InitialCost:   1000,
			GainPerSecond: 95,
			Gained:        0,
			RateOfGrowth:  0.6,
			LastEvaluated: now,
		},
		&Generator{
			Name:          G,
			Count:         0,
			InitialCost:   5000,
			GainPerSecond: 200,
			Gained:        0,
			RateOfGrowth:  0.8,
			LastEvaluated: now,
		},
	}
}

func (gs *Generators) Evaluate() float64 {
	var earned float64
	for _, g := range *gs {
		earned += g.Evaluate()
	}
	return earned
}

func (gs *Generators) Add(name string, count int) error {
	i, err := gs.generatorIndex(name)
	if err != nil {
		return err
	}

	g := (*gs)[i]

	g.Count = g.Count + int64(count)

	return nil
}

func (gs *Generators) Cost(name string, count int) (float64, error) {
	i, err := gs.generatorIndex(name)
	if err != nil {
		return 0, err
	}

	g := (*gs)[i]

	return g.TotalCost(int(g.Count) + count), nil
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

func (g *Generator) Evaluate() float64 {
	now := time.Now()
	since := now.Sub(g.LastEvaluated)
	gained := float64(g.Count) * g.GainPerSecond * since.Seconds()

	g.Gained += gained
	g.LastEvaluated = now

	return gained
}

func (g *Generator) TotalCost(count int) float64 {
	return g.InitialCost * math.Pow(math.E, float64(count-1)*g.RateOfGrowth)
}
