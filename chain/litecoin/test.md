# Bitcoin rpc api test

## 1.get fee
- request
```
grpcurl -plaintext -d '{
  "chain": "Litecoin",
  "coin": "LTC",
  "network": "mainet"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getFee
```
- response
```
{
  "code": "SUCCESS",
  "msg": "get litcoin fee success",
  "best_fee": "0.00011985",
  "best_fee_sat": "11985",
  "slow_fee": "",
  "normal_fee": "",
  "fast_fee": ""
}
```

## 2.get account
- request
```
grpcurl -plaintext -d '{
  "chain": "Litecoin",
  "network": "mainet",
  "address": "MVRidwRCeGpmfDoeZ2wc1LtjMEuQ5gSxnq"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getAccount
```

- response
```
{
  "code": "SUCCESS",
  "msg": "Get litecoin account info success",
  "network": "",
  "balance": "0.00297156"
}
```