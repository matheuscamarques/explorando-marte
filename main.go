package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/matheuscamarques/explorando-marte/modules"
)

func main() {
	//read stdin input and print it
	var satelite modules.SateliteStruct
	var planalto modules.PlanaltoStruct
	var sonda modules.SondaStruct
	reader := bufio.NewReader(os.Stdin)
	first := 0
	count := 0

	for {
		input, _ := reader.ReadString('\n')
		if len(input) <= 1 {
			break
		}
		if first == 0 {
			first = 1
			// get input values 5 5
			s := strings.Split(
				strings.TrimSuffix(input, "\n"), " ",
			)
			a, _ := strconv.Atoi(s[0])
			b, _ := strconv.Atoi(s[1])
			planalto = modules.PlanaltoStruct{LimiteHorizontal: a, LimiteVertical: b}
			satelite = modules.SateliteStruct{Planalto: planalto}
		} else {
			if count%2 == 0 {
				s := strings.Split(strings.TrimSuffix(input, "\n"), " ")
				x, _ := strconv.Atoi(s[0])
				y, _ := strconv.Atoi(s[1])
				dir := s[2]
				sonda = modules.SondaStruct{PosX: x, PosY: y, Dir: dir}
				count++
			} else {
				s := strings.TrimSuffix(input, "\n")
				satelite.PosicionarSonda(
					satelite.Command(sonda, s),
				)
				count++
			}
		}

	}

	satelite.ImprimirSondas()
}
