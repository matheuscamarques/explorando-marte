package main

import (
	"testing"

	"github.com/matheuscamarques/explorando-marte/modules"
)

func Test_criarPlanalto(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *modules.PlanaltoStruct
		wantErr bool
	}{
		{
			name: "Valid Command",
			args: args{
				input: "5 5",
			},
			wantErr: false,
		},
		{
			name: "Invalid Command",
			args: args{
				input: "A B",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := criarPlanalto(tt.args.input)
			if (err == nil) == tt.wantErr {
				t.Errorf("criarPlanalto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_criarSonda(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    *modules.SondaStruct
		wantErr bool
	}{
		{
			name: "Valid Command",
			args: args{
				input: "1 2 N",
			},
			wantErr: false,
		},
		{
			name: "Invalid Command",
			args: args{
				input: "A B",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := criarSonda(tt.args.input)
			if (err == nil) == tt.wantErr {
				t.Errorf("criarSonda() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
