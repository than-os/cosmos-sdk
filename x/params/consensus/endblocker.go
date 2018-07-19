package consensus

import (
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	params "github.com/cosmos/cosmos-sdk/x/params/store"
)

// Keys for parameter access
const (
	DefaultParamSpace = "ConsensusParams"

	BlockSizeSpace   = "BlockSize"
	TxSizeSpace      = "TxSize"
	BlockGossipSpace = "BlockGossip"

	MaxBytesKey      = "MaxBytes"
	MaxTxsKey        = "MaxTxs"
	MaxGasKey        = "MaxGas"
	PartSizeBytesKey = "PartSizeBytes"
)

func BlockMaxBytesKey() params.Key      { return params.NewKey(BlockSizeSpace, MaxBytesKey) }
func BlockMaxTxsKey() params.Key        { return params.NewKey(BlockSizeSpace, MaxTxsKey) }
func BlockMaxGasKey() params.Key        { return params.NewKey(BlockSizeSpace, MaxGasKey) }
func TxMaxBytesKey() params.Key         { return params.NewKey(TxSizeSpace, MaxBytesKey) }
func TxMaxGasKey() params.Key           { return params.NewKey(TxSizeSpace, MaxGasKey) }
func BlockPartSizeBytesKey() params.Key { return params.NewKey(BlockGossipSpace, PartSizeBytesKey) }

var (
	blockMaxBytesKey      = BlockMaxBytesKey()
	blockMaxTxsKey        = BlockMaxTxsKey()
	blockMaxGasKey        = BlockMaxGasKey()
	txMaxBytesKey         = TxMaxBytesKey()
	txMaxGasKey           = TxMaxGasKey()
	blockPartSizeBytesKey = BlockPartSizeBytesKey()
)

// nolint
func EndBlock(ctx sdk.Context, store params.Store) (updates *abci.ConsensusParams) {
	if store.Modified(ctx, blockMaxBytesKey) {
		updates = new(abci.ConsensusParams)
		updates.BlockSize = new(abci.BlockSize)
		store.MustGet(ctx, blockMaxBytesKey, &updates.BlockSize.MaxBytes)
	}

	if store.Modified(ctx, blockMaxTxsKey) {
		if updates == nil {
			updates = new(abci.ConsensusParams)
		}
		if updates.BlockSize == nil {
			updates.BlockSize = new(abci.BlockSize)
		}
		store.MustGet(ctx, blockMaxTxsKey, &updates.BlockSize.MaxTxs)
	}

	if store.Modified(ctx, blockMaxGasKey) {
		if updates == nil {
			updates = new(abci.ConsensusParams)
		}
		if updates.BlockSize == nil {
			updates.BlockSize = new(abci.BlockSize)
		}
		store.MustGet(ctx, blockMaxGasKey, &updates.BlockSize.MaxTxs)
	}

	if store.Modified(ctx, txMaxBytesKey) {
		if updates == nil {
			updates = new(abci.ConsensusParams)
		}
		updates.TxSize = new(abci.TxSize)
		store.MustGet(ctx, txMaxBytesKey, &updates.BlockSize.MaxTxs)
	}

	if store.Modified(ctx, txMaxGasKey) {
		if updates == nil {
			updates = new(abci.ConsensusParams)
		}
		if updates.TxSize == nil {
			updates.TxSize = new(abci.TxSize)
		}
		store.MustGet(ctx, txMaxGasKey, &updates.BlockSize.MaxTxs)
	}

	return
}
