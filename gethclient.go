package ethclient

import (
	"encoding/json"
	"github.com/zhouxun1995/ethclient/types"

	"github.com/zhouxun1995/xutil/xhttp"
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
