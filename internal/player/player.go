package player

import (
	"clicker/internal/generators"
	"fmt"
	"time"
)

type Player struct {
	Liquid     float64               `json:"liquid"`
	Generators generators.Generators `json:"generators"`
}

func New() *Player {
	now := time.Now()
	return &Player{
		Liquid: 2,
		Generators: generators.Generators{
			&generators.Generator{
				Name:          generators.A,
				Count:         0,
				Cost:          1,
				GainPerSecond: 5,
				Gained:        0,
				LastEvaluated: now,
			},
			&generators.Generator{
				Name:          generators.B,
				Cost:          1,
				Count:         0,
				GainPerSecond: 2,
				Gained:        0,
				LastEvaluated: now,
			},
		},
	}
}

func (p *Player) Loop() {
	for {
		time.Sleep(time.Millisecond * 300)
		p.Evaluate()
		fmt.Println("Player worth: ", p.Worth())
	}
}

func (p *Player) Evaluate() {
	p.Liquid += p.Generators.Evaluate()
}

func (p *Player) Worth() float64 {
	return p.Liquid
}

func (p *Player) Add(name string) error {
	fmt.Println("Buying ", name)
	cost, err := p.Generators.Cost(name)
	if err == nil && cost <= p.Liquid {
		p.Generators.Add(name)
		p.Liquid -= cost
	}
	return err
}
