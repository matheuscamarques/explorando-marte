package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/matheuscamarques/explorando-marte/modules"
)

func main() {
	var (
		satelite modules.SateliteStruct
		planalto *modules.PlanaltoStruct
		sonda    *modules.SondaStruct
		err      error
	)

	var first bool = true
	count := 0
	var str string = getUserInput()

	for _, input := range strings.Split(str, "\n") {
		if len(input) <= 1 {
			break
		}

		if first {
			first = false
			planalto, err = criarPlanalto(input)
			if err != nil {
				panic(err)
			}
			satelite = modules.SateliteStruct{Planalto: *planalto}
		} else {
			if count%2 == 0 {
				sonda, err = criarSonda(input)
				if err != nil {
					panic(err)
				}
				count++
			} else {
				tempSonda, err := satelite.Command(*sonda, input)
				if err != nil {
					panic(err)
				}

				satelite.PosicionarSonda(tempSonda)
				count++
			}
		}
	}
	satelite.ImprimirSondas()
}

func criarPlanalto(input string) (*modules.PlanaltoStruct, error) {
	s := strings.Split(
		strings.TrimSuffix(input, "\n"), " ",
	)

	a, err := strconv.Atoi(s[0])
	if err != nil {
		return nil, err
	}

	b, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, err
	}

	planalto := modules.PlanaltoStruct{LimiteHorizontal: a, LimiteVertical: b}

	return &planalto, nil
}

func criarSonda(input string) (*modules.SondaStruct, error) {
	s := strings.Split(strings.TrimSuffix(input, "\n"), " ")
	if len(s) != 3 {
		return nil, fmt.Errorf("command invalid")
	}

	x, err := strconv.Atoi(s[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, err
	}

	dir := modules.ParseDirection(s[2])
	if dir == modules.Invalid {
		return nil, fmt.Errorf("invalid direction")
	}

	return &modules.SondaStruct{PosX: x, PosY: y, Dir: dir}, nil
}

func getUserInput() (input string) {
	inputReader := bufio.NewReader(os.Stdin)

	for {
		consoleInput, _ := inputReader.ReadString('\n')

		if consoleInput == "\n" {
			break
		}

		input += consoleInput
	}

	return
}
