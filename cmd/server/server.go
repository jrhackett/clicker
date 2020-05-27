package main

import (
	"bufio"
	"clicker/api"
	"clicker/internal/generators"
	"clicker/internal/player"
	"fmt"
	"os"
)

func main() {
	p := player.New()
	go p.Loop()

	go userStdInput(p)

	api.Serve(p)
}

func userStdInput(p *player.Player) {
	reader := bufio.NewReader(os.Stdin)

	for {
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}

		switch char {
		case 'A', 'a':
			fmt.Println("Buying A")
			p.Add(string(generators.A))
		case 'B', 'b':
			fmt.Println("Buying B")
			p.Add(string(generators.B))
		}

		if err != nil {
			fmt.Println("error %s: %+v", char, err)
		}
	}
}
