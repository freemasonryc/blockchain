package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)


type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	sk         types.StakingKeeper
	paramspace types.ParamSubspace
}


func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, sk types.StakingKeeper, paramspace types.ParamSubspace) Keeper {

	if !paramspace.HasKeyTable() {
		paramspace = paramspace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:   key,
		cdc:        cdc,
		sk:         sk,
		paramspace: paramspace,
	}
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}


func (k Keeper) AddPubkey(ctx sdk.Context, pubkey cryptotypes.PubKey) error {
	bz, err := k.cdc.MarshalInterface(pubkey)
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	key := types.AddrPubkeyRelationKey(pubkey.Address())
	store.Set(key, bz)
	return nil
}


func (k Keeper) GetPubkey(ctx sdk.Context, a cryptotypes.Address) (cryptotypes.PubKey, error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.AddrPubkeyRelationKey(a))
	if bz == nil {
		return nil, fmt.Errorf("address %s not found", sdk.ConsAddress(a))
	}
	var pk cryptotypes.PubKey
	return pk, k.cdc.UnmarshalInterface(bz, &pk)
}



func (k Keeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress, fraction sdk.Dec, power, distributionHeight int64) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSlash,
			sdk.NewAttribute(types.AttributeKeyAddress, consAddr.String()),
			sdk.NewAttribute(types.AttributeKeyPower, fmt.Sprintf("%d", power)),
			sdk.NewAttribute(types.AttributeKeyReason, types.AttributeValueDoubleSign),
		),
	)

	k.sk.Slash(ctx, consAddr, distributionHeight, power, fraction)
}



func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSlash,
			sdk.NewAttribute(types.AttributeKeyJailed, consAddr.String()),
		),
	)

	k.sk.Jail(ctx, consAddr)
}

func (k Keeper) deleteAddrPubkeyRelation(ctx sdk.Context, addr cryptotypes.Address) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.AddrPubkeyRelationKey(addr))
}
