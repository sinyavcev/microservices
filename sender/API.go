package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Stock struct {
	GlobalQuote struct {
		Symbol string `json:"01. symbol"`
		Open   string `json:"02. open"`
		High   string `json:"03. high"`
		Low    string `json:"04. low"`
		Price  string `json:"05. price"`
		Volume string `json:"06. volume"`
	} `json:"Global Quote"`
}

func getStockQuotes(symbol string) ([]byte, error) {
	var stock Stock
	URL := fmt.Sprintf("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=ZR48ZK13NDQYGPLU", symbol)
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	errUnmarshal := json.Unmarshal(body, &stock)
	if errUnmarshal != nil {
		log.Fatalln(err)
	}
	var parseData, errMarshal = json.Marshal(&stock)
	if errMarshal != nil {
		log.Fatalln(errMarshal)
	}
	return parseData, nil
}
