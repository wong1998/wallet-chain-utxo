package dash

import (
	"github.com/btcsuite/btcd/rpcclient"

	"github.com/ethereum/go-ethereum/log"
)

type DashClient struct {
	*rpcclient.Client
	compressed bool
}

func NewDashClient(RpcUrl, RpcUser, RpcPass string) (*DashClient, error) {
	client, err := rpcclient.New(&rpcclient.ConnConfig{
		Host:         RpcUrl,
		User:         RpcUser,
		Pass:         RpcPass,
		HTTPPostMode: true,
		DisableTLS:   true,
	}, nil)
	if err != nil {
		log.Error("new dash rpc client fail", "err", err)
		return nil, err
	}
	return &DashClient{
		Client:     client,
		compressed: true,
	}, nil
}
