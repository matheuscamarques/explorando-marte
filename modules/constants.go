package modules

type direcao uint8

type lado uint8

const (
	Invalid direcao = iota
	N
	S
	W
	E
)

const (
	L lado = iota
	R
	M
)
