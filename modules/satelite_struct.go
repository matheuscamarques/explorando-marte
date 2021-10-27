package modules

import (
	"fmt"
	"strings"
)

type SateliteStruct struct {
	Planalto PlanaltoStruct
	sondas   []SondaStruct
}

func (sat *SateliteStruct) PosicionarSonda(sonda SondaStruct) error {
	if sonda.PosX < 0 || sonda.PosX > sat.Planalto.LimiteHorizontal {
		return fmt.Errorf("position invalid in X axis")
	} else if sonda.PosY < 0 || sonda.PosY > sat.Planalto.LimiteVertical {
		return fmt.Errorf("position invalid in Y axis")
	}

	sat.sondas = append(sat.sondas, sonda)
	return nil
}

func (sat *SateliteStruct) MoverSonda(sonda SondaStruct) SondaStruct {
	switch sonda.Dir {
	case N:
		if sonda.PosY < sat.Planalto.LimiteVertical {
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
		if sonda.PosX < sat.Planalto.LimiteHorizontal {
			sonda.PosX++
		}
	}
	return sonda
}

func (sat *SateliteStruct) GirarSonda(sonda SondaStruct, l lado) SondaStruct {
	switch l {
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

func trimBreakLine(str string) string {
	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")

	return str
}

func (sat *SateliteStruct) Command(sonda SondaStruct, commando string) (SondaStruct, error) {
	commando = trimBreakLine(commando)

	tempSonda := sonda

	for _, c := range commando {
		switch c {
		case 'L':
			tempSonda = sat.GirarSonda(sonda, L)
		case 'R':
			tempSonda = sat.GirarSonda(sonda, R)
		case 'M':
			tempSonda = sat.MoverSonda(sonda)
		default:
			return sonda, fmt.Errorf("command invalid")
		}
	}

	return tempSonda, nil
}

func ParseDirection(direction string) direcao {
	switch direction {
	case "N":
		return N
	case "S":
		return S
	case "W":
		return W
	case "E":
		return E
	}

	return Invalid
}
