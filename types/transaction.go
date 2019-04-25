package types

import "encoding/json"

type TXData struct {
	Method string
	To     string
	Amount int64
}

type ETHTransaction struct {
	From        string `json: "from"`
	Gas         string `json: "gas"`
	GasPrice    string `json: "gasPrice"`
	Hash        string `json: "hash"`
	Input       string `json: "input"`
	Nonce       string `json: "nonce"`
	To          string `json: "to"`
	Value       string `json: "value"`
	ParsedValue TXData `json:"-"`
	R           string `json: "r"`
	S           string `json: "s"`
	V           string `json: "v"`
}

func (tx ETHTransaction) Bytes() []byte {
	data, _ := json.Marshal(tx)
	return data
}
