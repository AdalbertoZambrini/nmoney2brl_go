package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rates := map[string]float64{
		"USD": 0.151, "EUR": 0.137, "JPY": 16.29, "GBP": 0.13,
		"CHF": 0.1402, "AUD": 0.2712, "CAD": 0.2374, "CNY": 1.251,
		"HKD": 1.326, "NZD": 0.2922, "SEK": 1.655, "NOK": 1.806,
		"DKK": 1.122, "SGD": 0.2249, "KRW": 242.97, "ZAR": 3.239,
		"MXN": 3.454, "INR": 14.71, "ILS": 0.63, "THB": 5.74,
		"IDR": 2875.0, "MYR": 0.754, "PHP": 9.74, "PLN": 0.644,
		"CZK": 3.77, "HUF": 61.59, "TRY": 6.49, "BGN": 0.293,
		"RON": 0.746,
	}

	if len(os.Args) != 3 {
		fmt.Println("Uso: ./convert [valor_em_brl] [moeda_destino]")
		os.Exit(1)
	}

	valorBRL, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Erro: valor invalido. Certifique-se de usar ponto para decimais.")
		os.Exit(1)
	}

	moedaDestino := strings.ToUpper(os.Args[2])

	taxa, existe := rates[moedaDestino]
	if !existe {
		fmt.Println("Erro: moeda de destino nao encontrada.")
		os.Exit(1)
	}

	resultado := valorBRL * taxa
	fmt.Printf("%.2f\n", resultado)
}
