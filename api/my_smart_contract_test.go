package api

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

// Test MySmartContract contract gets deployed correctly
func TestDeployMySmartContract(t *testing.T) {
	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	//Default chainID
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("Failed to create NewKeyedTransactorWithChainID : %v", err)
	}

	auth.GasLimit = uint64(300000)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, auth.GasLimit)

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())

	if err != nil {
		t.Fatalf("Failed to get SuggestGasPrice : %v", err)
	}

	auth.GasPrice = gasPrice

	//Deploy contract
	address, _, _, err := DeployApi(
		auth,
		blockchain,
	)

	// commit all pending transactions
	blockchain.Commit()

	if err != nil {
		t.Fatalf("Failed to deploy the MySmartContract contract: %v", err)
	}

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty address byte array instead")
	}

}

func TestHello(t *testing.T) {
	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	//Default chainID
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("Failed to create NewKeyedTransactorWithChainID : %v", err)
	}

	auth.GasLimit = uint64(300000)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, auth.GasLimit)

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())

	if err != nil {
		t.Fatalf("Failed to get SuggestGasPrice : %v", err)
	}

	auth.GasPrice = gasPrice

	//Deploy contract
	_, _, api, err := DeployApi(
		auth,
		blockchain,
	)

	if err != nil {
		t.Fatalf("Failed to deploy the MySmartContract contract: %v", err)
	}

	// commit all pending transactions
	blockchain.Commit()

	msg, err := api.Hello(&bind.CallOpts{})

	if err != nil {
		t.Fatalf("Failed to call Hello : %v", err)
	}

	assert.Equal(t, "Hello World", msg)
}

func TestGreet(t *testing.T) {
	//Setup simulated block chain
	key, _ := crypto.GenerateKey()
	//Default chainID
	chainID := big.NewInt(1337)
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		t.Fatalf("Failed to create NewKeyedTransactorWithChainID : %v", err)
	}

	auth.GasLimit = uint64(300000)
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, auth.GasLimit)

	gasPrice, err := blockchain.SuggestGasPrice(context.Background())

	if err != nil {
		t.Fatalf("Failed to get SuggestGasPrice : %v", err)
	}

	auth.GasPrice = gasPrice

	//Deploy contract
	_, _, api, err := DeployApi(
		auth,
		blockchain,
	)

	if err != nil {
		t.Fatalf("Failed to deploy the MySmartContract contract: %v", err)
	}

	// commit all pending transactions
	blockchain.Commit()

	msg, err := api.Greet(&bind.CallOpts{}, "Ganesh")

	if err != nil {
		t.Fatalf("Failed to call Greet : %v", err)
	}

	assert.Equal(t, "Ganesh", msg)
}
