package main

import (
	"testing"
)

func Test_formatCurrency(t *testing.T) {
	tests := []struct {
		name    string
		thb     float64
		c       *CurrencyConverter
		want    string
		wantErr bool
	}{
		{
			name: "THB to WON",
			thb:  100000.00,
			c: &CurrencyConverter{
				rate:    33.72,
				typeCur: "WON",
			},
			want: "3,372,000.00 WON",
		},
		{
			name: "THB to USD",
			thb:  100000.00,
			c: &CurrencyConverter{
				rate:    0.028,
				typeCur: "USD",
			},
			want: "2,800.00 USD",
		},
		{
			name: "THB to YEN",
			thb:  100000.00,
			c: &CurrencyConverter{
				rate:    3.4,
				typeCur: "YEN",
			},
			want: "340,000.00 YEN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := tt.c.Convert(tt.thb); got != tt.want {
				if (err != nil) != tt.wantErr {
					t.Errorf("CurrencyConverter.Convert() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Errorf("formatCurrency() = %v, want %v", got, tt.want)
			}
		})
	}
}
