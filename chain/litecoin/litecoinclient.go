package litecoin

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/ethereum/go-ethereum/log"
)

type LitecoinClient struct {
	*rpcclient.Client
	compressed bool
}

func NewLitecoinClient(RpcUrl, RpcUser, RpcPass string) (*LitecoinClient, error) {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         RpcUrl,
		User:         RpcUser,
		Pass:         RpcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if err != nil {
		log.Error("new litecoin rpc client fail", "err", err)
		return nil, err
	}
	return &LitecoinClient{
		Client:     client,
		compressed: true,
	}, nil
}
