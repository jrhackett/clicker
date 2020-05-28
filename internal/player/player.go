package player

import (
	"clicker/internal/generators"
	"time"
)

type (
	Player struct {
		Name       string                `json:"name"`
		Liquid     float64               `json:"liquid"`
		Generators generators.Generators `json:"generators"`
	}

	Players []*Player
)

func New(name string) *Player {
	return &Player{
		Name:       name,
		Liquid:     2,
		Generators: generators.New(),
	}
}

func (p *Player) Loop() {
	for {
		time.Sleep(time.Millisecond * 300)
		p.Evaluate()
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
	if err == nil && cost <= p.Liquid {
		p.Generators.Add(name, count)
		p.Liquid -= cost
	}
	return err
}

func (p *Player) Cost(name string, count int) (float64, error) {
	return p.Generators.Cost(name, count)
}

func (p *Player) Hit() {
	p.Liquid++
}
