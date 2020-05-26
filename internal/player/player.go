package player

import "clicker/internal/inventory"

type Player struct {
	Inventory inventory.Inventory
}

func (p *Player) Evaluate() {
	p.Inventory.Evaluate()
}

func (p *Player) Worth() float64 {
	return p.Inventory.Worth()
}
