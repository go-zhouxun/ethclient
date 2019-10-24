package types

import "encoding/json"

type ETHTxReceipt struct {
	TransactionHash   string `json:"transactionHash"`
	TransactionIndex  string `json:"transactionIndex"`
	BlockNumber       string `json:"blockNumber"`
	BlockHash         string `json:"blockHash"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	ContractAddress   string `json:"contractAddress"`
	Status            HexInt `json:"status"`
}

func (receipt ETHTxReceipt) Bytes() []byte {
	data, _ := json.Marshal(receipt)
	return data
}
