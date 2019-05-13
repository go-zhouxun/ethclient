package ethclient

import (
	"encoding/json"

	"github.com/go-zhouxun/ethclient/types"
	"github.com/go-zhouxun/xutil/xhttp"
)

type GethClient struct {
	host string `json:"host"`
}

func (geth GethClient) SendRequest(param types.EthRPCParam) (*types.EthRPCResp, error) {
	body, err := xhttp.HttpPost(geth.host, param.Bytes())
	if err != nil {
		return nil, err
	}
	var resp types.EthRPCResp
	err = json.Unmarshal(body, &resp)
	return &resp, err
}
