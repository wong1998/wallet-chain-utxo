package bitcoincash

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/log"
)

type BitcoinCashClient struct {
	*rpcclient.Client
}

func NewBitcoinCashClient(RpcUrl, RpcUser, RpcPass string) (*BitcoinCashClient, error) {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         RpcUrl,
		User:         RpcUser,
		Pass:         RpcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if err != nil {
		log.Error("new bitcoin cash rpc client fail", "err", err)
		return nil, err
	}
	return &BitcoinCashClient{
		Client: client,
	}, nil
}
