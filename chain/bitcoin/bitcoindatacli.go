package bitcoin

import (
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
	"github.com/dapplink-labs/chain-explorer-api/explorer/oklink"
)

type BitcoinData struct {
	BitcoinDataCli *oklink.ChainExplorerAdaptor
}

func NewBitcoinDataClient(baseUrl, apiKey string) (*BitcoinData, error) {
	btcCli, err := oklink.NewChainExplorerAdaptor(apiKey, baseUrl, false, time.Second*15)
	if err != nil {
		log.Error("New bitcion client fail", "err", err)
		return nil, err
	}
	return &BitcoinData{BitcoinDataCli: btcCli}, err
}

func (bd *BitcoinData) GetFee() (*gas_fee.GasEstimateFeeResponse, error) {
	gefr := &gas_fee.GasEstimateFeeRequest{
		ChainShortName: "BTC",
		ExplorerName:   "Bitcoin",
	}
	okResp, err := bd.BitcoinDataCli.GetEstimateGasFee(gefr)
	if err != nil {
		return nil, err
	}
	return okResp, nil
}
