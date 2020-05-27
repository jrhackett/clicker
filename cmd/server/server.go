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
			err = p.Add(string(generators.A), 1)
		case 'B', 'b':
			err = p.Add(string(generators.B), 1)
		}

		if err != nil {
			fmt.Println("error %s: %+v", char, err)
		}
	}
}
