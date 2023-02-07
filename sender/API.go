package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GlobalQuote struct {
	Stock Stock `json:"Global Quote"`
}

type Stock struct {
	Symbol string `json:"01. symbol"`
	Open   string `json:"02. open"`
	High   string `json:"03. high"`
	Low    string `json:"04. low"`
	Day    string `json:"07. latest trading day"`
}

func getStockQuotes() ([]byte, error) {
	nameStock := [3]string{"AAPL", "IBM", "BA"}
	var GlobalQuote GlobalQuote
	var infoStock []Stock
	for _, value := range nameStock {
		URL := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=ZR48ZK13NDQYGPLU", value)
		resp, err := http.Get(URL)
		if err != nil {
			log.Fatalln(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		errUnmarshal := json.Unmarshal(body, &GlobalQuote)
		if errUnmarshal != nil {
			log.Fatalln(err)
		}
		infoStock = append(infoStock, GlobalQuote.Stock)
	}
	parseData, errMarshal := json.Marshal(infoStock)

	if errMarshal != nil {
		log.Fatalln(errMarshal)
	}
	return parseData, nil
}
