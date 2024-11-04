package litecoin

import (
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
	"github.com/dapplink-labs/chain-explorer-api/explorer/oklink"
)

type LitecoinDataClient struct {
	LitecoinDataCli *oklink.ChainExplorerAdaptor
}

func NewLitecoinDataClient(baseUrl, apiKey string) (*LitecoinDataClient, error) {
	ltcCli, err := oklink.NewChainExplorerAdaptor(apiKey, baseUrl, false, time.Second*15)
	if err != nil {
		log.Error("New lit coin client fail", "err", err)
		return nil, err
	}
	return &LitecoinDataClient{LitecoinDataCli: ltcCli}, err
}

func (bd *LitecoinDataClient) GetFee() (*gas_fee.GasEstimateFeeResponse, error) {
	gefr := &gas_fee.GasEstimateFeeRequest{
		ChainShortName: "LTC",
		ExplorerName:   "LTC",
	}
	okResp, err := bd.LitecoinDataCli.GetEstimateGasFee(gefr)
	if err != nil {
		return nil, err
	}
	return okResp, nil
}

func (bd *LitecoinDataClient) GetAccountBalance(address string) (*account.AccountBalanceResponse, error) {
	accountItem := []string{address}
	symbol := []string{"LTC"}
	contractAddress := []string{"0x00"}
	page := []string{"1"}
	limit := []string{"10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "LTC",
		ExplorerName:    "Litecoin",
		Account:         accountItem,
		Symbol:          symbol,
		ContractAddress: contractAddress,
		Page:            page,
		Limit:           limit,
	}
	balanceResponse, err := bd.LitecoinDataCli.GetAccountBalance(acbr)
	if err != nil {
		return nil, err
	}
	return balanceResponse, nil
}
