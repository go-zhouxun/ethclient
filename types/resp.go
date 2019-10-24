package types

type EthRPCResp struct {
	JSONRPC string      `json:"jsonrpc"`
	Id      int64       `json:"id"`
	Result  interface{} `json:"result"`
}
