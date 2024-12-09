package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type CoinPriceResponse struct {
	Data []struct {
		ID     string  `json:"id"`
		Name   string  `json:"name"`
		Symbol string  `json:"symbol"`
		Price  float64 `json:"priceUsd,string"`
	} `json:"data"`
}



func getCryptoPrices() (map[string]float64, error) {

	URL := "https://api.coincap.io/v2/assets/"
	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	
	var prices CoinPriceResponse
	
	if err := json.Unmarshal(body, &prices); err != nil {
		panic(err)
	}
	
	pricesMap := make(map[string]float64)

	for _, value := range prices.Data {
		fmt.Printf("Price of %s (%s): $%.2f\n", value.Name, value.Symbol, value.Price)
		pricesMap[value.Name] = value.Price
		
	}


	return pricesMap,nil
}

func main() {
	
	pricesMap, err := getCryptoPrices()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Prices: %v\n", pricesMap["Bitcoin"])
}