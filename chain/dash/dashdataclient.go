package dash

import (
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
	"github.com/dapplink-labs/chain-explorer-api/explorer/oklink"
)

type DashDataClient struct {
	DashDataCli *oklink.ChainExplorerAdaptor
}

func NewDashDataClient(baseUrl, apiKey string) (*DashDataClient, error) {
	dashCli, err := oklink.NewChainExplorerAdaptor(apiKey, baseUrl, false, time.Second*15)
	if err != nil {
		log.Error("New dash client fail", "err", err)
		return nil, err
	}
	return &DashDataClient{DashDataCli: dashCli}, err
}

func (bd *DashDataClient) GetFee() (*gas_fee.GasEstimateFeeResponse, error) {
	gefr := &gas_fee.GasEstimateFeeRequest{
		ChainShortName: "Dash",
		ExplorerName:   "Dash",
	}
	okResp, err := bd.DashDataCli.GetEstimateGasFee(gefr)
	if err != nil {
		return nil, err
	}
	return okResp, nil
}

func (bd *DashDataClient) GetAccountBalance(address string) (*account.AccountBalanceResponse, error) {
	accountItem := []string{address}
	symbol := []string{"DASH"}
	contractAddress := []string{"0x00"}
	page := []string{"1"}
	limit := []string{"10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "DASH",
		ExplorerName:    "Dash",
		Account:         accountItem,
		ContractAddress: contractAddress,
		Symbol:          symbol,
		Page:            page,
		Limit:           limit,
	}
	balanceResponse, err := bd.DashDataCli.GetAccountBalance(acbr)
	if err != nil {
		return nil, err
	}
	return balanceResponse, nil
}
