package main

import (
	"fmt"
)

func main() {
	// read stdin input and print it
	var input string
	for {
		// read input
		_, err := fmt.Scan(&input)
		if err != nil {
			break
		}
		// print input
		fmt.Println(input)
	}

	var planalto PlanaltoStruct = PlanaltoStruct{5, 5}
	var satelite SateliteStruct = SateliteStruct{Planalto: planalto}

	var commando = "LMLMLMLMM"

	satelite.PosicionarSonda(
		satelite.Command(SondaStruct{1, 2, N}, commando),
	)

	commando = "MMRMMRMRRM"
	satelite.PosicionarSonda(
		satelite.Command(SondaStruct{3, 3, E}, commando),
	)

	satelite.ImprimirSondas()
	return
}
