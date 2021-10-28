package modules

import (
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
