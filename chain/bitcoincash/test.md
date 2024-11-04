# Bitcoin rpc api test

## 1.get fee
- request
```
grpcurl -plaintext -d '{
  "chain": "BitcoinCash",
  "coin": "BCH",
  "network": "mainet"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getFee
```
- response
```
{
  "code": "SUCCESS",
  "msg": "get bch fee success",
  "best_fee": "0.00004815",
  "best_fee_sat": "4815",
  "slow_fee": "",
  "normal_fee": "",
  "fast_fee": ""
}
```


## 2.get account
- request
```
grpcurl -plaintext -d '{
  "chain": "BitcoinCash",
  "network": "mainet",
  "address": "pqwjhsy0qsj8e6drfm5gua6lcdlqagczjq7vh5kak6"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getAccount
```

- response
```
{
  "code": "SUCCESS",
  "msg": "Get bch account info success",
  "network": "",
  "balance": "222.8"
}
```