package modules

import (
	"reflect"
	"testing"
)

/* func TestSateliteStruct_PosicionarSonda(t *testing.T) {
	type fields struct {
		Planalto PlanaltoStruct
		sondas   []SondaStruct
	}
	type args struct {
		sonda SondaStruct
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Valid",
			fields: ,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &SateliteStruct{
				Planalto: tt.fields.Planalto,
				sondas:   tt.fields.sondas,
			}
			sat.PosicionarSonda(tt.args.sonda)
		})
	}
} */

func TestSateliteStruct_MoverSonda(t *testing.T) {
	type fields struct {
		Planalto PlanaltoStruct
		sondas   []SondaStruct
	}
	type args struct {
		sonda SondaStruct
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SondaStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &SateliteStruct{
				Planalto: tt.fields.Planalto,
				sondas:   tt.fields.sondas,
			}
			if got := sat.MoverSonda(tt.args.sonda); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SateliteStruct.MoverSonda() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSateliteStruct_GirarSonda(t *testing.T) {
	type fields struct {
		Planalto PlanaltoStruct
		sondas   []SondaStruct
	}
	type args struct {
		sonda SondaStruct
		lado  rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SondaStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &SateliteStruct{
				Planalto: tt.fields.Planalto,
				sondas:   tt.fields.sondas,
			}
			if got := sat.GirarSonda(tt.args.sonda, tt.args.lado); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SateliteStruct.GirarSonda() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSateliteStruct_ImprimirSondas(t *testing.T) {
	type fields struct {
		Planalto PlanaltoStruct
		sondas   []SondaStruct
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &SateliteStruct{
				Planalto: tt.fields.Planalto,
				sondas:   tt.fields.sondas,
			}
			sat.ImprimirSondas()
		})
	}
}

func TestSateliteStruct_Command(t *testing.T) {
	type fields struct {
		Planalto PlanaltoStruct
		sondas   []SondaStruct
	}
	type args struct {
		sonda    SondaStruct
		commando string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   SondaStruct
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sat := &SateliteStruct{
				Planalto: tt.fields.Planalto,
				sondas:   tt.fields.sondas,
			}
			if got := sat.Command(tt.args.sonda, tt.args.commando); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SateliteStruct.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}
