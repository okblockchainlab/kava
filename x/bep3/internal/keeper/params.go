package keeper

import (
	"github.com/Kava-Labs/kava/x/bep3/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the total set of bep3 parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSubspace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the bep3 parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}