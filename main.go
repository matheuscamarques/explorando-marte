package main

import (
	"fmt"
)

const (
	N = 'N'
	S = 'S'
	W = 'W'
	E = 'E'
)

const (
	L = 'L'
	R = 'R'
	M = 'M'
)

type SondaStruct struct {
	PosX int
	PosY int
	Dir  rune
}

type PlanaltoStruct struct {
	limite_horizontal int
	limite_vertical   int
}

type SateliteStruct struct {
	Planalto PlanaltoStruct
	sondas   []SondaStruct
}

func (sat *SateliteStruct) PosicionarSonda(sonda SondaStruct) {
	sat.sondas = append(sat.sondas, sonda)
}

func (sat *SateliteStruct) MoverSonda(sonda SondaStruct) SondaStruct {
	switch sonda.Dir {
	case N:
		if sonda.PosY < sat.Planalto.limite_vertical {
			sonda.PosY++
		}
	case S:
		if sonda.PosY > 0 {
			sonda.PosY--
		}
	case W:
		if sonda.PosX > 0 {
			sonda.PosX--
		}
	case E:
		if sonda.PosX < sat.Planalto.limite_horizontal {
			sonda.PosX++
		}
	}
	return sonda
}

func (sat *SateliteStruct) GirarSonda(sonda SondaStruct, lado rune) SondaStruct {
	switch lado {
	case L:
		switch sonda.Dir {
		case N:
			sonda.Dir = W
		case S:
			sonda.Dir = E
		case W:
			sonda.Dir = S
		case E:
			sonda.Dir = N
		}
	case R:
		switch sonda.Dir {
		case N:
			sonda.Dir = E
		case S:
			sonda.Dir = W
		case W:
			sonda.Dir = N
		case E:
			sonda.Dir = S
		}
	}

	return sonda
}

func (sat *SateliteStruct) ImprimirSondas() {
	for _, sonda := range sat.sondas {
		fmt.Printf("%d %d %s\n", sonda.PosX, sonda.PosY, string(sonda.Dir))
	}
}

func (sat *SateliteStruct) Command(sonda SondaStruct, commando string) SondaStruct {
	for _, c := range commando {
		if c == L || c == R {
			sonda = sat.GirarSonda(sonda, c)
		}
		if c == M {
			sonda = sat.MoverSonda(sonda)
		}

	}
	return sonda
}

func main() {

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
