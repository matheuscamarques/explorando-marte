package modules

import (
	"reflect"
	"testing"
)

var satelite = SateliteStruct{
	Planalto: PlanaltoStruct{
		LimiteHorizontal: 5,
		LimiteVertical:   5,
	},
}

func TestSateliteStruct_Command(t *testing.T) {
	var sat = satelite

	type args struct {
		sonda      SondaStruct
		sondaFinal SondaStruct
		commando   string
	}
	tests := []struct {
		name    string
		args    args
		want    SondaStruct
		wantErr bool
	}{
		{
			name: "Test 01 - Commando Correto",
			args: args{
				sonda: SondaStruct{
					PosX: 1,
					PosY: 1,
					Dir:  N,
				},
				sondaFinal: SondaStruct{
					PosX: 1,
					PosY: 2,
					Dir:  N,
				},
				commando: "LMLMLMLMM",
			},
			wantErr: false,
		},
		{
			name: "Test 02 - Comando Invalido",
			args: args{
				sonda: SondaStruct{
					PosX: 2,
					PosY: 2,
					Dir:  N,
				},
				sondaFinal: SondaStruct{
					PosX: 2,
					PosY: 2,
					Dir:  N,
				},
				commando: "LMLMLMLMMRO",
			},
			wantErr: true,
		},
		{
			name: "Test 03 - Comando Mudar Direçao sem se mecher",
			args: args{
				sonda: SondaStruct{
					PosX: 2,
					PosY: 2,
					Dir:  N,
				},
				sondaFinal: SondaStruct{
					PosX: 2,
					PosY: 2,
					Dir:  N,
				},
				commando: "LLLL",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			tt.args.sonda, err = sat.Command(tt.args.sonda, tt.args.commando)

			if (err != nil) != tt.wantErr {
				t.Errorf("SateliteStruct.Command() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (tt.args.sonda.PosX != tt.args.sondaFinal.PosX) || (tt.args.sonda.PosY != tt.args.sondaFinal.PosY) || (tt.args.sonda.Dir != tt.args.sondaFinal.Dir) {
				t.Errorf("%v SateliteStruct.Command() = %v, want %v", tt.args.sonda, tt.args.sondaFinal, tt.name)
			}
		})
	}
}

func TestParseDirection(t *testing.T) {
	type args struct {
		direction string
	}
	tests := []struct {
		name string
		args args
		want direcao
	}{
		{
			name: "Test 01 - Direção Norte",
			args: args{
				direction: "N",
			},
			want: N,
		},
		{
			name: "Test 02 - Direção Sul",
			args: args{
				direction: "S",
			},
			want: S,
		},
		{
			name: "Test 03 - Direção Leste",
			args: args{
				direction: "E",
			},
			want: E,
		},
		{
			name: "Test 04 - Direção Oeste",
			args: args{
				direction: "W",
			},
			want: W,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDirection(tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trimBreakLine(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 01 - Trim Break Line",
			args: args{
				str: "LMLMLMLMM\n",
			},
			want: "LMLMLMLMM",
		},
		{
			name: "Test 02 - Trim Break Line",
			args: args{
				str: "LMLMLMLMM\r\n",
			},
			want: "LMLMLMLMM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimBreakLine(tt.args.str); got != tt.want {
				t.Errorf("trimBreakLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSateliteStruct_GirarSonda(t *testing.T) {
	type args struct {
		sonda SondaStruct
		l     lado
	}
	tests := []struct {
		name string
		sat  *SateliteStruct
		args args
		want SondaStruct
	}{
		{
			name: "Test 01 - Girar Sonda do NORTE para Esquerda",
			sat:  &satelite,
			args: args{
				sonda: SondaStruct{
					PosX: 1,
					PosY: 2,
					Dir:  N,
				},
				l: L,
			},
			want: SondaStruct{
				PosX: 1,
				PosY: 2,
				Dir:  W,
			},
		},
		{
			name: "Test 02 - Girar Sonda do NORTE para Direita",
			sat:  &satelite,
			args: args{
				sonda: SondaStruct{
					PosX: 1,
					PosY: 2,
					Dir:  N,
				},
				l: R,
			},
			want: SondaStruct{
				PosX: 1,
				PosY: 2,
				Dir:  E,
			},
		},
		{
			name: "Test 03 - Girar Sonda do SUL para Direita",
			sat:  &satelite,
			args: args{
				sonda: SondaStruct{
					PosX: 3,
					PosY: 3,
					Dir:  S,
				},
				l: R,
			},
			want: SondaStruct{
				PosX: 3,
				PosY: 3,
				Dir:  W,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sat.GirarSonda(tt.args.sonda, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SateliteStruct.GirarSonda() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSateliteStruct_MoverSonda(t *testing.T) {
	type args struct {
		sonda SondaStruct
	}
	tests := []struct {
		name string
		sat  *SateliteStruct
		args args
		want SondaStruct
	}{
		{
			name: "Test 01 - Mover Sonda para Norte",
			sat:  &satelite,
			args: args{
				sonda: SondaStruct{
					PosX: 1,
					PosY: 2,
					Dir:  N,
				},
			},
			want: SondaStruct{
				PosX: 1,
				PosY: 3,
				Dir:  N,
			},
		},
		{
			name: "Test 02 - Mover Sonda para Sul",
			sat:  &satelite,
			args: args{
				sonda: SondaStruct{
					PosX: 3,
					PosY: 3,
					Dir:  S,
				},
			},
			want: SondaStruct{
				PosX: 3,
				PosY: 2,
				Dir:  S,
			},
		},
		{
			name: "Test 03 - Mover Sonda para Leste",
			sat:  &satelite,
			args: args{
				sonda: SondaStruct{
					PosX: 3,
					PosY: 3,
					Dir:  E,
				},
			},
			want: SondaStruct{
				PosX: 4,
				PosY: 3,
				Dir:  E,
			},
		},
		{
			name: "Test 04 - Mover Sonda para Oeste",
			sat:  &satelite,
			args: args{
				sonda: SondaStruct{
					PosX: 3,
					PosY: 3,
					Dir:  W,
				},
			},
			want: SondaStruct{
				PosX: 2,
				PosY: 3,
				Dir:  W,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sat.MoverSonda(tt.args.sonda); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SateliteStruct.MoverSonda() = %v, want %v", got, tt.want)
			}
		})
	}
}
