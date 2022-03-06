# unit-test-smart-contracts

## Pre requisite
* Go version - go1.17.1
* solc version - 0.8.12+
* abigen version - 1.10.16-stable-20356e57
* Ganache version - ganache-2.5.4

## Compiling the Smart Contract

```
$ solc --optimize --abi ./contract/my_smart_contract.sol -o build
$ solc --optimize --bin ./contract/my_smart_contract.sol -o build
$ abigen --abi=./build/MySmartContract.abi --bin=./build/MySmartContract.bin --pkg=api --out=./api/my_smart_contract.go

```

## To run unit tests

* Go to folder ./api
* Run the following command
```
go test -v
```
Test will display the following output

```
=== RUN   TestDeployMySmartContract
--- PASS: TestDeployMySmartContract (0.00s)
=== RUN   TestHello
--- PASS: TestHello (0.00s)
=== RUN   TestGreet
--- PASS: TestGreet (0.00s)
PASS
ok  	github.com/partha-techsophy/unit-test-smart-contracts/api	0.019s
```

## To deploy contract

* Download Ganache
* Run Ganache
* Note down the Private key of one account
* Replace "PRIVATE_KEY" at line number 21
* run the following command
```
go run my_smartcontract_deploy.go
```