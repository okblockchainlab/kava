package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/kava-labs/kava/x/bep3/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper of the bep3 store
type Keeper struct {
	key           sdk.StoreKey
	cdc           *codec.Codec
	paramSubspace subspace.Subspace
	supplyKeeper  types.SupplyKeeper
	codespace     sdk.CodespaceType
}

// NewKeeper creates a bep3 keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, sk types.SupplyKeeper, paramstore subspace.Subspace, codespace sdk.CodespaceType) Keeper {
	if addr := sk.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}
	keeper := Keeper{
		key:           key,
		cdc:           cdc,
		paramSubspace: paramstore.WithKeyTable(types.ParamKeyTable()),
		supplyKeeper:  sk,
		codespace:     codespace,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetAtomicSwap puts the AtomicSwap into the store, and updates any indexes.
func (k Keeper) SetAtomicSwap(ctx sdk.Context, atomicSwap types.AtomicSwap, swapID []byte) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.AtomicSwapKeyPrefix)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(atomicSwap)
	store.Set(swapID, bz)
}

// GetAtomicSwap gets an AtomicSwap from the store.
func (k Keeper) GetAtomicSwap(ctx sdk.Context, swapID []byte) (types.AtomicSwap, bool) {
	var atomicSwap types.AtomicSwap

	store := prefix.NewStore(ctx.KVStore(k.key), types.AtomicSwapKeyPrefix)
	bz := store.Get(swapID)
	if bz == nil {
		return atomicSwap, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &atomicSwap)
	return atomicSwap, true
}

// IterateAtomicSwaps provides an iterator over all stored AtomicSwaps.
// For each AtomicSwap, cb will be called. If cb returns true, the iterator will close and stop.
func (k Keeper) IterateAtomicSwaps(ctx sdk.Context, cb func(atomicSwap types.AtomicSwap) (stop bool)) {
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.key), types.AtomicSwapKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var atomicSwap types.AtomicSwap
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &atomicSwap)

		if cb(atomicSwap) {
			break
		}
	}
}

// GetAllAtomicSwaps returns all AtomicSwaps from the store
func (k Keeper) GetAllAtomicSwaps(ctx sdk.Context) (atomicSwaps types.AtomicSwaps) {
	k.IterateAtomicSwaps(ctx, func(atomicSwap types.AtomicSwap) bool {
		atomicSwaps = append(atomicSwaps, atomicSwap)
		return false
	})
	return
}

// SetAssetSupply updates an asset's current active supply
func (k Keeper) SetAssetSupply(ctx sdk.Context, asset sdk.Coin, coinID []byte) {
	store := prefix.NewStore(ctx.KVStore(k.key), types.AssetSupplyKeyPrefix)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(asset)
	store.Set(coinID, bz)
}

// GetAssetSupply gets an asset's current supply from the store.
func (k Keeper) GetAssetSupply(ctx sdk.Context, denom []byte) (sdk.Coin, bool) {
	var asset sdk.Coin

	store := prefix.NewStore(ctx.KVStore(k.key), types.AssetSupplyKeyPrefix)
	bz := store.Get(denom)
	if bz == nil {
		return sdk.Coin{}, false
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &asset)
	return asset, true
}

// IterateAssetSupplies provides an iterator over all stored AssetSupplies.
// For each AssetSupply, cb will be called. If cb returns true, the iterator will close and stop.
func (k Keeper) IterateAssetSupplies(ctx sdk.Context, cb func(asset sdk.Coin) (stop bool)) {
	iterator := sdk.KVStorePrefixIterator(ctx.KVStore(k.key), types.AssetSupplyKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var asset sdk.Coin
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &asset)

		if cb(asset) {
			break
		}
	}
}

// GetAllAssetSupplies returns all asset supplies from the store as an array of sdk.Coin
func (k Keeper) GetAllAssetSupplies(ctx sdk.Context) (assets []sdk.Coin) {
	k.IterateAssetSupplies(ctx, func(asset sdk.Coin) bool {
		assets = append(assets, asset)
		return false
	})
	return
}
