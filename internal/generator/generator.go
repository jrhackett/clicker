package generator

import (
	"time"
)

type (
	Generator struct {
		Name          string
		Count         int64
		GainPerSecond float64
		Gained        float64
		LastEvaluated time.Time
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
