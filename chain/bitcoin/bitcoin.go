package bitcoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"

	"github.com/ethereum/go-ethereum/log"

	"github.com/dapplink-labs/wallet-chain-utxo/chain"
	"github.com/dapplink-labs/wallet-chain-utxo/config"
	common2 "github.com/dapplink-labs/wallet-chain-utxo/rpc/common"
	"github.com/dapplink-labs/wallet-chain-utxo/rpc/utxo"
)

const ChainName = "Bitcoin"

type ChainAdaptor struct {
	btcClient       *BtcClient
	btcDataClient   *BitcoinData
	thirdPartClient *BcClient
}

func NewChainAdaptor(conf *config.Config) (chain.IChainAdaptor, error) {
	btcClient, err := NewBtcClient(conf.WalletNode.Btc.RpcUrl, conf.WalletNode.Btc.RpcUser, conf.WalletNode.Btc.RpcPass)
	if err != nil {
		log.Error("new bitcoin rpc client fail", "err", err)
		return nil, err
	}
	btcDataClient, err := NewBitcoinDataClient(conf.WalletNode.Btc.DataApiUrl, conf.WalletNode.Btc.DataApiKey)
	if err != nil {
		log.Error("new bitcoin data client fail", "err", err)
		return nil, err
	}
	bcClient, err := NewBlockChainClient(conf.WalletNode.Btc.TpApiUrl)
	if err != nil {
		log.Error("new blockchain client fail", "err", err)
		return nil, err
	}
	return &ChainAdaptor{
		btcClient:       btcClient,
		btcDataClient:   btcDataClient,
		thirdPartClient: bcClient,
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
	switch req.Format {
	case "p2pkh":
		return nil, nil
	case "p2wpkh":
		return nil, nil
	case "p2sh":
		return nil, nil
	case "p2tr":
		return nil, nil
	default:
		return nil, nil
	}
}

func (c *ChainAdaptor) ValidAddress(req *utxo.ValidAddressRequest) (*utxo.ValidAddressResponse, error) {
	switch req.Format {
	case "p2pkh":
		return nil, nil
	case "p2wpkh":
		return nil, nil
	case "p2sh":
		return nil, nil
	case "p2tr":
		return nil, nil
	default:
		return nil, nil
	}
}

func (c *ChainAdaptor) GetFee(req *utxo.FeeRequest) (*utxo.FeeResponse, error) {
	gasFeeResp, err := c.btcDataClient.GetFee()
	if err != nil {
		return &utxo.FeeResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	return &utxo.FeeResponse{
		Code:       common2.ReturnCode_SUCCESS,
		Msg:        "get fee success",
		BestFee:    gasFeeResp.BestTransactionFee,
		BestFeeSat: gasFeeResp.BestTransactionFeeSat,
		SlowFee:    gasFeeResp.SlowGasPrice,
		NormalFee:  gasFeeResp.StandardGasPrice,
		FastFee:    gasFeeResp.RapidGasPrice,
	}, nil
}

func (c *ChainAdaptor) GetAccount(req *utxo.AccountRequest) (*utxo.AccountResponse, error) {
	balance, err := c.thirdPartClient.GetAccountBalance(req.Address)
	if err != nil {
		return &utxo.AccountResponse{
			Code:    common2.ReturnCode_ERROR,
			Msg:     "get btc balance fail",
			Balance: "0",
		}, err
	}
	return &utxo.AccountResponse{
		Code:    common2.ReturnCode_SUCCESS,
		Msg:     "get btc balance success",
		Balance: balance,
	}, nil
}

func (c *ChainAdaptor) GetUnspentOutputs(req *utxo.UnspentOutputsRequest) (*utxo.UnspentOutputsResponse, error) {
	utxoList, err := c.thirdPartClient.GetAccountUtxo(req.Address)
	if err != nil {
		return &utxo.UnspentOutputsResponse{
			Code:           common2.ReturnCode_ERROR,
			Msg:            err.Error(),
			UnspentOutputs: nil,
		}, err
	}
	var unspentOutputList []*utxo.UnspentOutput
	for _, value := range utxoList {
		unspentOutput := &utxo.UnspentOutput{
			TxHashBigEndian: value.TxHashBigEndian,
			TxHash:          value.TxHash,
			TxOutputN:       value.TxOutputN,
			Script:          value.Script,
			Value:           value.Value,
			TxIndex:         value.TxIndex,
		}
		unspentOutputList = append(unspentOutputList, unspentOutput)
	}
	return &utxo.UnspentOutputsResponse{
		Code:           common2.ReturnCode_SUCCESS,
		Msg:            "get unspent outputs success",
		UnspentOutputs: unspentOutputList,
	}, nil
}

func (c *ChainAdaptor) GetBlockByNumber(req *utxo.BlockNumberRequest) (*utxo.BlockResponse, error) {
	blockHash, err := c.btcClient.Client.GetBlockHash(req.Height)
	if err != nil {
		log.Error("get block hash by number fail", "err", err)
		return &utxo.BlockResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "get block hash fail",
		}, err
	}
	var params []json.RawMessage
	numBlocksJSON, _ := json.Marshal(blockHash)
	params = []json.RawMessage{numBlocksJSON}
	block, _ := c.btcClient.Client.RawRequest("getblock", params)
	var resultBlock BlockData
	err = json.Unmarshal(block, &resultBlock)
	if err != nil {
		log.Error("Unmarshal json fail", "err", err)
	}
	for _, txid := range resultBlock.Tx {
		txIdJson, _ := json.Marshal(txid)
		boolJSON, _ := json.Marshal(true)
		dataJSON := []json.RawMessage{txIdJson, boolJSON}
		tx, err := c.btcClient.Client.RawRequest("getrawtransaction", dataJSON)
		if err != nil {
			fmt.Println("get raw transaction fail", "err", err)
		}
		var rawTx RawTransactionData
		err = json.Unmarshal(tx, &rawTx)
		if err != nil {
			log.Error("json unmarshal fail", "err", err)
			return nil, err
		}
		for _, v := range rawTx.Vin {
			fmt.Println("v.TxId==", v.TxId)
		}
	}
	return &utxo.BlockResponse{}, err
}

func (c *ChainAdaptor) GetBlockByHash(req *utxo.BlockHashRequest) (*utxo.BlockResponse, error) {
	var params []json.RawMessage
	numBlocksJSON, _ := json.Marshal(req.Hash)
	params = []json.RawMessage{numBlocksJSON}
	block, _ := c.btcClient.Client.RawRequest("getblock", params)
	var resultBlock BlockData
	err := json.Unmarshal(block, &resultBlock)
	if err != nil {
		log.Error("Unmarshal json fail", "err", err)
	}
	for _, txid := range resultBlock.Tx {
		txIdJson, _ := json.Marshal(txid)
		boolJSON, _ := json.Marshal(true)
		dataJSON := []json.RawMessage{txIdJson, boolJSON}
		tx, err := c.btcClient.Client.RawRequest("getrawtransaction", dataJSON)
		if err != nil {
			fmt.Println("get raw transaction fail", "err", err)
		}
		var rawTx RawTransactionData
		err = json.Unmarshal(tx, &rawTx)
		if err != nil {
			log.Error("json unmarshal fail", "err", err)
			return nil, err
		}
		for _, v := range rawTx.Vin {
			fmt.Println("v.TxId==", v.TxId)
		}
	}
	return &utxo.BlockResponse{}, nil
}

func (c *ChainAdaptor) GetBlockHeaderByHash(req *utxo.BlockHeaderHashRequest) (*utxo.BlockHeaderResponse, error) {
	hash, err := chainhash.NewHashFromStr(req.Hash)
	if err != nil {
		log.Error("format string to hash fail", "err", err)
	}
	blockHeader, err := c.btcClient.Client.GetBlockHeader(hash)
	if err != nil {
		return &utxo.BlockHeaderResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "get block header fail",
		}, err
	}
	return &utxo.BlockHeaderResponse{
		Code:       common2.ReturnCode_SUCCESS,
		Msg:        "get block header success",
		ParentHash: blockHeader.PrevBlock.String(),
		Number:     string(blockHeader.Version),
		BlockHash:  req.Hash,
		MerkleRoot: blockHeader.MerkleRoot.String(),
	}, nil
}

func (c *ChainAdaptor) GetBlockHeaderByNumber(req *utxo.BlockHeaderNumberRequest) (*utxo.BlockHeaderResponse, error) {
	blockNumber := req.Height
	if req.Height == 0 {
		latestBlock, err := c.btcClient.Client.GetBlockCount()
		if err != nil {
			return &utxo.BlockHeaderResponse{
				Code: common2.ReturnCode_ERROR,
				Msg:  "get latest block fail",
			}, err
		}
		blockNumber = latestBlock
	}
	blockHash, err := c.btcClient.Client.GetBlockHash(blockNumber)
	if err != nil {
		log.Error("get block hash by number fail", "err", err)
		return &utxo.BlockHeaderResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "get block hash fail",
		}, err
	}
	blockHeader, err := c.btcClient.Client.GetBlockHeader(blockHash)
	if err != nil {
		return &utxo.BlockHeaderResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "get block header fail",
		}, err
	}
	return &utxo.BlockHeaderResponse{
		Code:       common2.ReturnCode_SUCCESS,
		Msg:        "get block header success",
		ParentHash: blockHeader.PrevBlock.String(),
		Number:     strconv.FormatInt(blockNumber, 10),
		BlockHash:  blockHash.String(),
		MerkleRoot: blockHeader.MerkleRoot.String(),
	}, nil
}

func (c *ChainAdaptor) SendTx(req *utxo.SendTxRequest) (*utxo.SendTxResponse, error) {
	r := bytes.NewReader([]byte(req.RawTx))
	var msgTx wire.MsgTx
	err := msgTx.Deserialize(r)
	if err != nil {
		return &utxo.SendTxResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	txHash, err := c.btcClient.SendRawTransaction(&msgTx, true)
	if err != nil {
		return &utxo.SendTxResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  err.Error(),
		}, err
	}
	if strings.Compare(msgTx.TxHash().String(), txHash.String()) != 0 {
		log.Error("broadcast transaction, tx hash mismatch", "local hash", msgTx.TxHash().String(), "hash from net", txHash.String(), "signedTx", req.RawTx)
	}
	return &utxo.SendTxResponse{
		Code:   common2.ReturnCode_SUCCESS,
		Msg:    "send tx success",
		TxHash: txHash.String(),
	}, nil
}

func (c *ChainAdaptor) GetTxByAddress(req *utxo.TxAddressRequest) (*utxo.TxAddressResponse, error) {
	transaction, err := c.thirdPartClient.GetTransactionsByAddress(req.Address, strconv.Itoa(int(req.Page)), strconv.Itoa(int(req.Pagesize)))
	if err != nil {
		return &utxo.TxAddressResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "get transaction list fail",
			Tx:   nil,
		}, err
	}
	var tx_list []*utxo.TxMessage
	for _, ttxs := range transaction.Txs {
		var from_addrs []*utxo.Address
		var to_addrs []*utxo.Address
		var value_list []*utxo.Value
		var direction int32
		for _, inputs := range ttxs.Inputs {
			from_addrs = append(from_addrs, &utxo.Address{Address: inputs.PrevOut.Addr})
		}
		tx_fee := ttxs.Fee
		for _, out := range ttxs.Out {
			to_addrs = append(to_addrs, &utxo.Address{Address: out.Addr})
			value_list = append(value_list, &utxo.Value{Value: out.Value.String()})
		}
		datetime := ttxs.Time.String()
		if strings.EqualFold(req.Address, from_addrs[0].Address) {
			direction = 0
		} else {
			direction = 1
		}
		tx := &utxo.TxMessage{
			Hash:     ttxs.Hash,
			Froms:    from_addrs,
			Tos:      to_addrs,
			Values:   value_list,
			Fee:      tx_fee.String(),
			Status:   utxo.TxStatus_Success,
			Type:     direction,
			Height:   ttxs.BlockHeight.String(),
			Datetime: datetime,
		}
		tx_list = append(tx_list, tx)
	}
	return &utxo.TxAddressResponse{
		Code: common2.ReturnCode_SUCCESS,
		Msg:  "get transaction list success",
		Tx:   tx_list,
	}, nil
}

func (c *ChainAdaptor) GetTxByHash(req *utxo.TxHashRequest) (*utxo.TxHashResponse, error) {
	transaction, err := c.thirdPartClient.GetTransactionsByHash(req.Hash)
	if err != nil {
		return &utxo.TxHashResponse{
			Code: common2.ReturnCode_ERROR,
			Msg:  "get transaction list fail",
			Tx:   nil,
		}, err
	}
	var from_addrs []*utxo.Address
	var to_addrs []*utxo.Address
	var value_list []*utxo.Value
	for _, inputs := range transaction.Inputs {
		from_addrs = append(from_addrs, &utxo.Address{Address: inputs.PrevOut.Addr})
	}
	tx_fee := transaction.Fee
	for _, out := range transaction.Out {
		to_addrs = append(to_addrs, &utxo.Address{Address: out.Addr})
		value_list = append(value_list, &utxo.Value{Value: out.Value.String()})
	}
	datetime := transaction.Time.String()
	txMsg := &utxo.TxMessage{
		Hash:     transaction.Hash,
		Froms:    from_addrs,
		Tos:      to_addrs,
		Values:   value_list,
		Fee:      tx_fee.String(),
		Status:   utxo.TxStatus_Success,
		Type:     0,
		Height:   transaction.BlockHeight.String(),
		Datetime: datetime,
	}
	return &utxo.TxHashResponse{
		Code: common2.ReturnCode_SUCCESS,
		Msg:  "get transaction success",
		Tx:   txMsg,
	}, nil
}

func (c *ChainAdaptor) CreateUnSignTransaction(req *utxo.UnSignTransactionRequest) (*utxo.UnSignTransactionResponse, error) {
	return &utxo.UnSignTransactionResponse{
		Code:     common2.ReturnCode_SUCCESS,
		Msg:      "create un sign transaction success",
		UnSignTx: "",
	}, nil
}

func (c *ChainAdaptor) BuildSignedTransaction(req *utxo.SignedTransactionRequest) (*utxo.SignedTransactionResponse, error) {
	return &utxo.SignedTransactionResponse{
		Code:     common2.ReturnCode_SUCCESS,
		Msg:      "build signed transaction success",
		SignedTx: "",
	}, nil
}

func (c *ChainAdaptor) DecodeTransaction(req *utxo.DecodeTransactionRequest) (*utxo.DecodeTransactionResponse, error) {
	return &utxo.DecodeTransactionResponse{
		Code:   common2.ReturnCode_SUCCESS,
		Msg:    "decode transaction success",
		TxList: nil,
	}, nil
}

func (c *ChainAdaptor) VerifySignedTransaction(req *utxo.VerifyTransactionRequest) (*utxo.VerifyTransactionResponse, error) {
	return &utxo.VerifyTransactionResponse{
		Code:   common2.ReturnCode_SUCCESS,
		Msg:    "verify transaction success",
		Verify: true,
	}, nil
}
