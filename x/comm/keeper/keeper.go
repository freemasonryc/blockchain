package keeper

import (
	"fmt"
	"freemasonry.cc/blockchain/core"
	chatkeeper "freemasonry.cc/blockchain/x/chat/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/tendermint/tendermint/libs/log"

	"freemasonry.cc/blockchain/x/comm/types"
)


type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	stakingKeeper stakingKeeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	chatKeeper    chatkeeper.Keeper
}


func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	ck chatkeeper.Keeper,
	stakingKeeper stakingKeeper.Keeper,
) Keeper {

	



	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: ak,
		bankKeeper:    bk,
		chatKeeper:    ck,
		stakingKeeper: stakingKeeper,
	}
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) KVHelper(ctx sdk.Context) StoreHelper {
	store := ctx.KVStore(k.storeKey)
	return StoreHelper{
		store,
	}
}

func (k Keeper) AddressBookSet(ctx sdk.Context, fromAddress string, AddressBook []string) error {
	store := k.KVHelper(ctx)
	key := types.KeyPrefixAddressBook + fromAddress
	err := store.Set(key, AddressBook)
	if err != nil {
		return err
	}
	return nil
}


func (k Keeper) GatewayBonus(ctx sdk.Context) error {
	
	if ctx.BlockHeight()%14400 == 0 {
		
		index := ctx.BlockHeight() / 5256000
		am := sdk.NewInt(10000)
		amount := sdk.NewCoin(sdk.DefaultBondDenom, am.Quo(sdk.NewInt(1<<index)))
		if amount.IsZero() {
			return nil
		}
		return k.bankKeeper.SendCoins(ctx, core.ContractGatewayBonus, core.ContractAddressFee, sdk.NewCoins(amount))
	}
	return nil
}
