package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var errorMessage = fmt.Errorf("both numbers must be non-negative")

func TestAddPositive(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr error
	}{
		{
			name:    "should return added number and no error",
			args:    args{2, 2},
			want:    4,
			wantErr: nil,
		},
		{
			name:    "should return error and zero",
			args:    args{-2, 1},
			want:    0,
			wantErr: errorMessage,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddPositive(tt.args.x, tt.args.y)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestSqrtNumbers(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should return 2",
			args: args{4},
			want: "2",
		},
		{
			name: "should return 2i",
			args: args{-4},
			want: "2i",
		},
		{
			name: "should return i",
			args: args{-1},
			want: "i",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SqrtNumber(tt.args.x)
			assert.Equal(t, tt.want, got)
		})
	}
}
