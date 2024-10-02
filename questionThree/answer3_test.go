package main

import "testing"

func TestCar_CalculateDistance(t *testing.T) {
	type args struct {
		gasLiters float64
	}
	tests := []struct {
		name string
		c    *Car
		args args
		want string
	}{
		{
			name: "Distance KM",
			c:    NewCar(),
			args: args{10.0},
			want: "123 KM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.CalculateDistance(tt.args.gasLiters); got != tt.want {
				t.Errorf("Car.CalculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
