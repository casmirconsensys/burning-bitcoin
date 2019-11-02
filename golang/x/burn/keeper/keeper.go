package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/summa-tx/bitcoin-spv/golang/btcspv"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper instantiates a new keeper
func NewKeeper(storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

func (k Keeper) setSomeThing(ctx sdk.Context, thing []byte) {
	digest := btcspv.Hash256(thing)
	store := ctx.KVStore(k.storeKey)
	store.Set(digest[:], thing[:])
}

// SetSomeThing sets a thing
func (k Keeper) SetSomeThing(ctx sdk.Context, thing []byte) {
	k.setSomeThing(ctx, thing)
}

// GetSomeThing gets a thing
func (k Keeper) GetSomeThing(ctx sdk.Context, digest [32]byte) []byte {
	store := ctx.KVStore(k.storeKey)
	return store.Get(digest[:])
}
