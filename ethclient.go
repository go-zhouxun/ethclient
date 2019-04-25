package ethclient

import (
	"bytes"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"

	"github.com/zhouxun1995/ethclient/types"
	"github.com/zhouxun1995/xlog"
	"github.com/zhouxun1995/xutil/hexutil"
)

type EthClient struct {
	gethClient GethClient
	logger     xlog.XLog
}

func NewEthClient(gethUrl string, logger xlog.XLog) *EthClient {
	return &EthClient{
		gethClient: GethClient{host: gethUrl},
		logger:     logger,
	}
}

func (ethClient EthClient) LastBlockNumber() int64 {
	params := types.NewEthRPCParam("eth_blockNumber", 1, []interface{}{})
	resp, err := ethClient.gethClient.SendRequest(params)
	if err != nil {
		ethClient.logger.Error("ethclient get eth last block number failed, %v", err)
		return -1
	}
	result, ok := resp.Result.(string)
	if !ok {
		return -1
	}
	height, err := hexutil.DecodeUint64(result)
	if err != nil {
		return -1
	}
	return int64(height)
}

func (client EthClient) GetBlockByNumber(height int64) *types.EthBlock {
	params := types.NewEthRPCParam("eth_getBlockByNumber", 1, []interface{}{
		hexutil.EncodeUint64(uint64(height)),
		false,
	})

	resp, err := client.gethClient.SendRequest(params)
	if err != nil {
		client.logger.Error("ethclient get block by number failed, %v", err)
		return nil
	}
	result, ok := resp.Result.(bool)
	if ok && !result {
		return nil
	}
	resultBody, _ := json.Marshal(resp.Result)
	var block types.EthBlock
	err = json.Unmarshal(resultBody, &block)
	if err != nil {
		client.logger.Error("ethclient get block by number unmarshal failed, %v", err)
		return nil
	}
	return &block
}

func (client EthClient) GetTransaction(hash string) *types.ETHTransaction {
	param := types.NewEthRPCParam("eth_getTransactionByHash", 1, []interface{}{hash})
	resp, err := client.gethClient.SendRequest(param)
	if err != nil {
		client.logger.Error("ethclient get transaction failed, %v", err)
		return nil
	}
	result, ok := resp.Result.(bool)
	if ok && !result {
		return nil
	}
	resultBody, _ := json.Marshal(resp.Result)

	var tx types.ETHTransaction
	err = json.Unmarshal(resultBody, &tx)
	if err != nil {
		client.logger.Error("ethclient get transaction unmarshal failed, %v", err)
		return nil
	}

	return &tx
}

func (client EthClient) GetTransactionReceipt(hash string) *types.ETHTxReceipt {
	params := types.NewEthRPCParam("eth_getTransactionReceipt", 1, []interface{}{hash})
	resp, err := client.gethClient.SendRequest(params)
	if err != nil {
		client.logger.Error("ethclient get transaction receipt failed, %v", err)
		return nil
	}
	result, ok := resp.Result.(bool)
	if ok && !result {
		return nil
	}
	resultBody, _ := json.Marshal(resp.Result)

	var receipt types.ETHTxReceipt
	err = json.Unmarshal(resultBody, &receipt)
	if err != nil {
		client.logger.Error("ethclient get transaction receipt unmarshal failed, %v", err)
		return nil
	}

	return &receipt
}

func (client EthClient) SendRawTransaction(transaction *types2.Transaction) string {
	var buffer bytes.Buffer
	err := transaction.EncodeRLP(&buffer)
	if err != nil {
		client.logger.Error("ethclient send raw transaction failed, %v", err)
	}
	param := types.NewEthRPCParam("eth_sendRawTransaction", 1, []interface{}{hexutil.Encode(buffer.Bytes())})
	resp, err := client.gethClient.SendRequest(param)
	if err != nil {
		client.logger.Error("ethclient send raw transaction failed, %v", err)
		return ""
	}
	hash, ok := resp.Result.(common.Hash)
	if !ok {
		return ""
	}
	return hexutil.Encode(hash[:])
}

func (client EthClient) GetNonce(address string) uint64 {
	if !strings.HasPrefix(address, "0x") {
		address = "0x" + address
	}
	param := types.NewEthRPCParam("eth_getTransactionCount", 1, []interface{}{address, "latest"})
	resp, err := client.gethClient.SendRequest(param)
	if err != nil {
		client.logger.Error("ethclient get nonce failed, %v", err)
		return 0
	}
	str, ok := resp.Result.(string)
	if !ok {
		client.logger.Error("ethclient get nonce failed, %v", err)
		return 0
	}
	bi := new(big.Int)
	bi, ok = bi.SetString(str, 0)
	if !ok {
		return 0
	}
	return bi.Uint64()
}
