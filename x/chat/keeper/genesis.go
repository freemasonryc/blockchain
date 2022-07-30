package keeper

import (
	types2 "freemasonry.cc/blockchain/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)


func (k Keeper) InitGenesis(ctx sdk.Context) {
	k.SetParams(ctx, types2.DefaultParams())
}


func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {

	return &types.GenesisState{}
}
