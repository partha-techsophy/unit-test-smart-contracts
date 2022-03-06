# unit-test-smart-contracts

Run this 
go mod init github.com/partha-techsophy/unit-test-smart-contracts


## Pre requisite
* Go version - go1.17.1
* solc version - 0.8.12+
* abigen version 1.10.16-stable-20356e57

## Compiling the Smart Contract

```
$ solc --optimize --abi ./contract/my_smart_contract.sol -o build
$ solc --optimize --bin ./contract/my_smart_contract.sol -o build
$ abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/my_smart_contract.go

```