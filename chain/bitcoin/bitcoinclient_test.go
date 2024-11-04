package bitcoin

import (
	"fmt"
	"testing"
)

func TestGetLatestBlock(t *testing.T) {
	host := "nd-314-634-477.p2pify.com"
	user := "lucid-swanson"
	pass := "salon-ahead-vanish-dial-curdle-arise"
	clients, err := NewBtcClient(host, user, pass)
	if err != nil {
		return
	}

	fmt.Println(clients)
}
