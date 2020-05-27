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
	return &Player{
		Liquid:     2,
		Generators: generators.New(),
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

func (p *Player) Add(name string, count int) error {
	cost, err := p.Generators.Cost(name, count)
	fmt.Println("Buying ", name, cost, count)
	if err == nil && cost <= p.Liquid {
		p.Generators.Add(name, count)
		p.Liquid -= cost
	}
	return err
}
