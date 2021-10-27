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
		sonda    SondaStruct
		commando string
	}
	tests := []struct {
		name    string
		args    args
		want    SondaStruct
		wantErr bool
	}{
		{
			name: "Test 01",
			args: args{
				sonda: SondaStruct{
					PosX: 1,
					PosY: 1,
					Dir:  N,
				},
				commando: "LMLMLMLMM",
			},
			wantErr: false,
		},
		{
			name: "Test 02",
			args: args{
				sonda: SondaStruct{
					PosX: 2,
					PosY: 2,
					Dir:  N,
				},
				commando: "LMLMLMLMMRO",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := sat.Command(tt.args.sonda, tt.args.commando)

			if (err != nil) != tt.wantErr {
				t.Errorf("SateliteStruct.Command() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
