package model

type Address struct {
	Street  string `json:"street"`
	Number  int32  `json:"number"`
	City    string `json:"city"`
	Country string `json:"country"`
}
