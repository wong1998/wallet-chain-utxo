package bitcoincash

import (
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"time"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
	"github.com/dapplink-labs/chain-explorer-api/explorer/oklink"
)

type BitcoinCashDataClient struct {
	BitcoinCashDataCli *oklink.ChainExplorerAdaptor
}

func NewBitcoinCashDataClient(baseUrl, apiKey string) (*BitcoinCashDataClient, error) {
	btcCli, err := oklink.NewChainExplorerAdaptor(apiKey, baseUrl, false, time.Second*15)
	if err != nil {
		log.Error("New bitcoin cash client fail", "err", err)
		return nil, err
	}
	return &BitcoinCashDataClient{BitcoinCashDataCli: btcCli}, err
}

func (bdc *BitcoinCashDataClient) GetFee() (*gas_fee.GasEstimateFeeResponse, error) {
	gefr := &gas_fee.GasEstimateFeeRequest{
		ChainShortName: "BCH",
		ExplorerName:   "BCH",
	}
	okResp, err := bdc.BitcoinCashDataCli.GetEstimateGasFee(gefr)
	if err != nil {
		return nil, err
	}
	return okResp, nil
}

func (bdc *BitcoinCashDataClient) GetAccountBalance(address string) (*account.AccountBalanceResponse, error) {
	accountItem := []string{address}
	symbol := []string{"BCH"}
	page := []string{"1"}
	contractAddress := []string{"0x00"}
	limit := []string{"10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "BCH",
		ExplorerName:    "BitcoinCash",
		Account:         accountItem,
		ContractAddress: contractAddress,
		Symbol:          symbol,
		Page:            page,
		Limit:           limit,
	}
	balanceResponse, err := bdc.BitcoinCashDataCli.GetAccountBalance(acbr)
	if err != nil {
		return nil, err
	}
	return balanceResponse, nil
}
