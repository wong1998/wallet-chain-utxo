# Bitcoin rpc api test

## 1.get fee
- request
```
grpcurl -plaintext -d '{
  "chain": "Bitcoin",
  "coin": "Btc",
  "network": "mainnet"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getFee
```
- response
```
{
  "code": "SUCCESS",
  "msg": "get fee success",
  "best_fee": "0.00001427",
  "best_fee_sat": "1427",
  "slow_fee": "",
  "normal_fee": "",
  "fast_fee": ""
}
```

## 2.get account
- request
```
grpcurl -plaintext -d '{
  "chain": "Bitcoin",
  "network": "mainet",
  "address": "bc1qa2eu6p5rl9255e3xz7fcgm6snn4wl5kdfh7zpt05qp5fad9dmsys0qjg0e"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getAccount
```

- response
```
{
  "code": "SUCCESS",
  "msg": "get btc balance success",
  "network": "",
  "balance": "5171031881312"
}
```