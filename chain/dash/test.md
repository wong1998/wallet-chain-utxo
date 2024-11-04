# Bitcoin rpc api test

## 1.get fee
- request
```
grpcurl -plaintext -d '{
  "chain": "Dash",
  "coin": "Dash",
  "network": "mainet"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getFee
```
- response
```
{
  "code": "SUCCESS",
  "msg": "get dash fee success",
  "best_fee": "0.00003733",
  "best_fee_sat": "3733",
  "slow_fee": "",
  "normal_fee": "",
  "fast_fee": ""
}
```

## 2.get account
- request
```
grpcurl -plaintext -d '{
  "chain": "Dash",
  "network": "mainet",
  "address": "Xrh6ouWeDBR8nLo95du2RLq8PbPZMe5vks"
}' 127.0.0.1:8289 dapplink.utxo.WalletUtxoService.getAccount
```

- response
```
{
  "code": "SUCCESS",
  "msg": "Get dash account info success",
  "network": "",
  "balance": "0.000505"
}
```