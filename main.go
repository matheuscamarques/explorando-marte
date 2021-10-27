package main

import (
	"github.com/matheuscamarques/explorando-marte/modules"
)

func main() {
	// read stdin input and print it
	//var input string
	//for {
	//	// read input
	//	_, err := fmt.Scan(&input)
	//	if err != nil {
	//		break
	//	}
	//	// print input
	//	fmt.Println(input)
	//}

	planalto := modules.PlanaltoStruct{LimiteHorizontal: 5, LimiteVertical: 5}
	satelite := modules.SateliteStruct{Planalto: planalto}

	var commando = "LMLMLMLMM"

	satelite.PosicionarSonda(
		satelite.Command(modules.SondaStruct{1, 2, modules.N}, commando),
	)

	commando = "MMRMMRMRRM"
	satelite.PosicionarSonda(
		satelite.Command(modules.SondaStruct{3, 3, modules.E}, commando),
	)

	satelite.ImprimirSondas()
	return
}
