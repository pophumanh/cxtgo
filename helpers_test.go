package cxtgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAmountToLotSize(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		lot       float64
		precision int
		amount    float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test with lot of zero and invalid amount",
			args: args{
				lot:       0.00100000,
				precision: 8,
				amount:    0.00010000,
			},
			want: 0,
		},
		{
			name: "test with lot",
			args: args{
				lot:       0.00100000,
				precision: 3,
				amount:    1.39,
			},
			want: 1.389,
		},
		{
			name: "test with big decimal",
			args: args{
				lot:       0.00100000,
				precision: 8,
				amount:    11.31232419283240912834434,
			},
			want: 11.312,
		},
		{
			name: "test with big number",
			args: args{
				lot:       0.0010000,
				precision: 8,
				amount:    14000.14000,
			},
			want: 14000.140,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.want, AmountToLotSize(tt.args.lot, tt.args.precision, tt.args.amount))
		})
	}
}

func TestFromString(t *testing.T) {
	assert := assert.New(t)
	type args struct {
		input    string
		splitter string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test with empty arguments",
			args: args{},
			want: 0,
		},
		{
			name: "test with input and no splitter",
			args: args{
				input: "000",
			},
			want: 0,
		},
		{
			name: "test with input and comma splitter",
			args: args{
				input:    "0,0",
				splitter: ",",
			},
			want: 1,
		},
		{
			name: "test with input and dot splitter and high precision",
			args: args{
				input:    "0000000000000000000000000000000.00000000000",
				splitter: ".",
			},
			want: 11,
		},
		{
			name: "test with usual input",
			args: args{
				input:    "0.000055",
				splitter: ".",
			},
			want: 5,
		},
		{
			name: "test with input and multiple dot splitters",
			args: args{
				input:    "0000000000000000000000000000000.....00000000000",
				splitter: ".",
			},
			want: 0,
		},
		{
			name: "test with input and multiple dot splitters",
			args: args{
				input:    "000.0.0.0",
				splitter: ".",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(tt.want, Zeros(tt.args.input, tt.args.splitter), tt.name)
		})
	}
}
