package dash

import (
	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/wallet-chain-utxo/chain"
	"github.com/dapplink-labs/wallet-chain-utxo/config"
	common2 "github.com/dapplink-labs/wallet-chain-utxo/rpc/common"
	"github.com/dapplink-labs/wallet-chain-utxo/rpc/utxo"
)

const ChainName = "Dash"

type ChainAdaptor struct {
	dashClient     *DashClient
	dashDataClient *DashDataClient
}

func NewChainAdaptor(conf *config.Config) (chain.IChainAdaptor, error) {
	dashClient, err := NewDashClient(conf.WalletNode.Dash.RpcUrl, conf.WalletNode.Dash.RpcUser, conf.WalletNode.Dash.RpcPass)
	if err != nil {
		log.Error("new dash rpc client fail", "err", err)
		return nil, err
	}
	dashDataClient, err := NewDashDataClient(conf.WalletNode.Btc.DataApiUrl, conf.WalletNode.Btc.DataApiKey)
	if err != nil {
		log.Error("new dash data client fail", "err", err)
		return nil, err
	}
	return &ChainAdaptor{
		dashClient:     dashClient,
		dashDataClient: dashDataClient,
	}, nil
}

func (c *ChainAdaptor) GetSupportChains(req *utxo.SupportChainsRequest) (*utxo.SupportChainsResponse, error) {
	return &utxo.SupportChainsResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "Support this chain",
		Support: true,
	}, nil
}

func (c *ChainAdaptor) ConvertAddress(req *utxo.ConvertAddressRequest) (*utxo.ConvertAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) ValidAddress(req *utxo.ValidAddressRequest) (*utxo.ValidAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) GetFee(req *utxo.FeeRequest) (*utxo.FeeResponse, error) {
	gasFeeResp, err := c.dashDataClient.GetFee()
	if err != nil {
		return &utxo.FeeResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	return &utxo.FeeResponse{
		Code:       common2.ReturnCode_SUCCESS,
		Msg:        "get dash fee success",
		BestFee:    gasFeeResp.BestTransactionFee,
		BestFeeSat: gasFeeResp.BestTransactionFeeSat,
		SlowFee:    gasFeeResp.SlowGasPrice,
		NormalFee:  gasFeeResp.StandardGasPrice,
		FastFee:    gasFeeResp.RapidGasPrice,
	}, nil
}

func (c *ChainAdaptor) GetAccount(req *utxo.AccountRequest) (*utxo.AccountResponse, error) {
	balance, err := c.dashDataClient.GetAccountBalance(req.Address)
	if err != nil {
		return &utxo.AccountResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "Get dash account info fail",
		}, err
	}
	return &utxo.AccountResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "Get dash account info success",
		Balance: balance.BalanceStr,
	}, nil
}

func (c *ChainAdaptor) GetUnspentOutputs(req *utxo.UnspentOutputsRequest) (*utxo.UnspentOutputsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) GetBlockByNumber(req *utxo.BlockNumberRequest) (*utxo.BlockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) GetBlockByHash(req *utxo.BlockHashRequest) (*utxo.BlockResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) GetBlockHeaderByHash(req *utxo.BlockHeaderHashRequest) (*utxo.BlockHeaderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) GetBlockHeaderByNumber(req *utxo.BlockHeaderNumberRequest) (*utxo.BlockHeaderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) SendTx(req *utxo.SendTxRequest) (*utxo.SendTxResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) GetTxByAddress(req *utxo.TxAddressRequest) (*utxo.TxAddressResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) GetTxByHash(req *utxo.TxHashRequest) (*utxo.TxHashResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) CreateUnSignTransaction(req *utxo.UnSignTransactionRequest) (*utxo.UnSignTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) BuildSignedTransaction(req *utxo.SignedTransactionRequest) (*utxo.SignedTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) DecodeTransaction(req *utxo.DecodeTransactionRequest) (*utxo.DecodeTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChainAdaptor) VerifySignedTransaction(req *utxo.VerifyTransactionRequest) (*utxo.VerifyTransactionResponse, error) {
	//TODO implement me
	panic("implement me")
}
