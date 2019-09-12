package main

import (
	"testing"
)

func Test_intervalCalc(t *testing.T) {
	type args struct {
		cost    float64
		advCost float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"cost є (1.1,∞)", args{1.2, 0}, -20},
		{"cost є [1.1,∞]", args{1.1, 0}, -20},
		{"cost є (1, 1.1)", args{1.05, 0}, -10},
		{"cost є [1, 1.1)", args{1, 0}, -10},
		{"cost є (0, 1) & cost>=advCost & diff=5", args{0.7, 0.65}, -5},
		{"cost є (0, 1) & cost>=advCost & diff>5", args{0.7, 0.5}, -5},
		{"cost є (0, 1) & cost>=advCost & diff<5", args{0.7, 0.69}, -1},
		{"cost є (0, 1) & cost<advCost & diff=10", args{0.7, 0.8}, 10},
		{"cost є (0, 1) & cost<advCost & diff>=10", args{0.7, 0.9}, 10},
		{"cost є (0, 1) & cost<advCost & diff<10", args{0.7, 0.75}, 5},
		{"cost є (-∞, 0)", args{-0.7, 0.69}, 10},
		{"advCost є (-∞, 0)", args{0.7, -0.8}, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intervalCalc(tt.args.cost, tt.args.advCost); got != tt.want {
				t.Errorf("intervalCalc() = %v, want %v", got, tt.want)
			}
		})
	}
}
