package l2

import (
	"github.com/bitdao-io/bitnetwork/indexer/bindings/l2erc20"
	"github.com/bitdao-io/bitnetwork/indexer/db"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func QueryERC20(address common.Address, client *ethclient.Client) (*db.Token, error) {
	contract, err := l2erc20.NewL2ERC20(address, client)
	if err != nil {
		return nil, err
	}

	name, err := contract.Name(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	symbol, err := contract.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	decimals, err := contract.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return &db.Token{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}, nil
}