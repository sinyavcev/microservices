package model

type Stock struct {
	Symbol string `json:"01. symbol" bson:"symbol"`
	Open   string `json:"02. open" bson:"open"`
	High   string `json:"03. high" bson:"high"`
	Low    string `json:"04. low" bson:"low"`
	Day    string `json:"07. latest trading day" bson:"day"`
}
