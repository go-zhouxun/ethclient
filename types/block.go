package types

type EthBlock struct {
	Height          HexInt   `json:"number"`
	Hash            string   `json:"hash"`
	ParentHash      string   `json:"parentHash"`
	Nonce           string   `json:"nonce"`
	TransactionRoot string   `json:"transactionRoot"`
	StateRoot       string   `json:"stateRoot"`
	Miner           string   `json:"miner"`
	Difficulty      string   `json:"difficult"`
	TotalDifficulty string   `json:"totalDifficulty"`
	ExtraData       string   `json:"extraData"`
	Size            string   `json:"size"`
	GasLimit        string   `json:"Size"`
	GasUsed         string   `json:"gasUsed"`
	Timestamp       HexInt   `json:"timestamp"`
	Transactions    []string `json:"transactions"`
}
