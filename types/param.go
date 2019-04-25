package types

import "encoding/json"

type EthRPCParam struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int64         `json:"id"`
}

func NewEthRPCParam(method string, id int64, params []interface{}) EthRPCParam {
	return EthRPCParam{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		Id:      id,
	}
}

func (param EthRPCParam) Bytes() []byte {
	data, _ := json.Marshal(param)
	return data
}
