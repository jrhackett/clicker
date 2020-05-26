package main

import (
	"clicker/internal/generator"
	"clicker/internal/inventory"
	"clicker/internal/player"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	player := player.Player{
		Inventory: inventory.Inventory{
			Liquid: 0,
			Generators: generator.Generators{
				&generator.Generator{
					Name:          "Test",
					Count:         1,
					GainPerSecond: 5,
					Gained:        0,
					LastEvaluated: now,
				},
				&generator.Generator{
					Name:          "Test 2",
					Count:         1,
					GainPerSecond: 2,
					Gained:        0,
					LastEvaluated: now,
				},
			},
		},
	}

	fmt.Println("Player inventory: %+v", player.Inventory)

	for {
		time.Sleep(time.Millisecond * 500)
		player.Evaluate()
		fmt.Println("Player worth: ", player.Worth())
	}
}
