package main

import (
	"fmt"
	"log"
	"strings"
)

type CurrencyConverter struct {
	rate    float64
	typeCur string
}

func NewCurrencyConverter(rate float64, typeCur string) *CurrencyConverter {
	return &CurrencyConverter{
		rate:    rate,
		typeCur: typeCur,
	}
}

func (cc *CurrencyConverter) Convert(thb float64) (string, error) {
	converted := thb * cc.rate
	return formatCurrency(converted, cc.typeCur), nil
}

func formatCurrency(amount float64, typeCur string) string {
	amountStr := fmt.Sprintf("%.2f", amount)

	parts := strings.Split(amountStr, ".")

	intPart := parts[0]
	decimalPart := ""
	if len(parts) > 1 {
		decimalPart = "." + parts[1]
	}

	formattedIntPart := ""
	for i, digit := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			formattedIntPart += ","
		}
		formattedIntPart += string(digit)
	}
	return fmt.Sprintf("%s " + typeCur, formattedIntPart + decimalPart)
}

func main() {
	thbAmount := 100000.0
	conversionRate := 33.80
	typeCurrency := "WON"
	converter := NewCurrencyConverter(conversionRate, typeCurrency)
	convertedAmount, err := converter.Convert(thbAmount)
	if err != nil {
		log.Fatalf("Error converting currency: %v", err)
	}
	fmt.Printf("Input: Number on THB (%.2f)\nOutput: String: Number of converted (%s)\n", thbAmount, convertedAmount)
}
