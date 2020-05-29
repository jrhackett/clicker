package main

import (
	"clicker/api"
	"clicker/internal/player"
)

func main() {
	p := player.New("Jake")
	go p.Loop()

	api.Serve(p)
}
