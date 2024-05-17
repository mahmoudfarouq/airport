package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_computeItinerary(t *testing.T) {
	tests := map[string]struct {
		args [][]string
		want []string
	}{
		"empty request": {
			args: [][]string{},
			want: []string{},
		},
		"single tuple request": {
			args: [][]string{
				{"A", "B"}, {"B", "C"},
			},
			want: []string{"A", "B", "C"},
		},
		"multiple tuple request": {
			args: [][]string{
				{"B", "C"}, {"A", "B"}, {"C", "E"},
			},
			want: []string{"A", "B", "C", "E"},
		},
		"task example": {
			args: [][]string{
				{"LAX", "DXB"}, {"JFK", "LAX"}, {"SFO", "SJC"}, {"DXB", "SFO"},
			},
			want: []string{"JFK", "LAX", "DXB", "SFO", "SJC"},
		},
	}

	for name, tt := range tests {
		tt := tt

		t.Run(name, func(t *testing.T) {
			got := computeItinerary(tt.args)
			require.Equal(t, tt.want, got)
		})
	}
}
