package keeper

import (
	"errors"
	"fmt"
	"freemasonry.cc/blockchain/cmd/config"
	"freemasonry.cc/blockchain/core"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"freemasonry.cc/blockchain/x/chat/types"
)


type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}


func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	bk types.BankKeeper,
) Keeper {

	



	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: ak,
		bankKeeper:    bk,
	}
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) KVHelper(ctx sdk.Context) storeHelper {
	store := ctx.KVStore(k.storeKey)
	return storeHelper{
		store,
	}
}

func (k Keeper) GerRegisterInfo(ctx sdk.Context, fromAddress string) (types.UserInfo, error) {

	store := k.KVHelper(ctx)
	key := types.KeyPrefixRegisterInfo + fromAddress

	var userinfo types.UserInfo

	if store.Has(key) {

		err := store.GetUnmarshal(key, &userinfo)
		if err != nil {
			return types.UserInfo{}, err
		}

		return userinfo, nil

	} else {

		return types.UserInfo{}, errors.New("USER-INFO-NOT-FOUND")

	}
}

func (k Keeper) SetRegisterInfo(ctx sdk.Context, userInfo types.UserInfo) error {

	store := k.KVHelper(ctx)
	key := types.KeyPrefixRegisterInfo + userInfo.FromAddress

	err := store.Set(key, userInfo)
	if err != nil {
		return err
	}

	return nil
}


func (k Keeper) MortgageSendCoin(ctx sdk.Context, transferType, toAddress, fromAddress, nodeAddress string, mortgageAmount sdk.Coin) (sdk.Coin, error) {

	
	chatParams := k.GetParams(ctx)
	if mortgageAmount.Amount.Sub(chatParams.MinMortgageCoin.Amount).LT(sdk.NewInt(0)) {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrMortgageAmount
	}

	
	mortgageMoney := mortgageAmount.Amount
	mortgageMoneyDec := mortgageMoney.ToDec()

	
	mortgageNodeDec := mortgageMoneyDec.Mul(core.MortgageRatioDecNode)
	
	mortgageBurnDec := mortgageMoneyDec.Mul(core.MortgageRatioDecBurn)
	
	mortgageCommunityDec := mortgageMoneyDec.Mul(core.MortgageRatioDecCommunity)
	
	mortgageEcologicalDec := mortgageMoneyDec.Mul(core.MortgageRatioDecEcological)
	
	mortgagePosDec := mortgageMoneyDec.Mul(core.MortgageRatioDecPos)
	
	mortgageRemainDec := mortgageMoneyDec.Mul(core.MortgageRatioDecRemain)
	

	
	
	accPosAddress := authtypes.NewModuleAddress(authtypes.FeeCollectorName)
	
	accFromAddress, err := sdk.AccAddressFromBech32(fromAddress)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrAddressFormat
	}
	accNodeAddress, err := sdk.AccAddressFromBech32(nodeAddress)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrAddressFormat
	}
	
	accCommunityAddress, err := sdk.AccAddressFromBech32(chatParams.CommunityAddress)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrAddressFormat
	}
	
	accEcologicalAddress, err := sdk.AccAddressFromBech32(chatParams.EcologicalAddress)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrAddressFormat
	}
	
	
	toNodeCoin := sdk.NewCoin(config.BaseDenom, mortgageNodeDec.TruncateInt())
	toNodeCoins := sdk.NewCoins(toNodeCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accNodeAddress, toNodeCoins)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
	}

	
	toCommunityCoin := sdk.NewCoin(config.BaseDenom, mortgageCommunityDec.TruncateInt())
	toCommunityCoins := sdk.NewCoins(toCommunityCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accCommunityAddress, toCommunityCoins)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
	}

	
	toEcologicalCoin := sdk.NewCoin(config.BaseDenom, mortgageEcologicalDec.TruncateInt())
	toEcologicalCoins := sdk.NewCoins(toEcologicalCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accEcologicalAddress, toEcologicalCoins)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
	}

	
	toPosCoin := sdk.NewCoin(config.BaseDenom, mortgagePosDec.TruncateInt())
	toPosCoins := sdk.NewCoins(toPosCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accPosAddress, toPosCoins)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
	}

	
	burnCoin := sdk.NewCoin(config.BaseDenom, mortgageBurnDec.TruncateInt())
	burnCoins := sdk.NewCoins(burnCoin)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accFromAddress, types.ModuleBurnName, burnCoins)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleBurnName, burnCoins)
	if err != nil {
		return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
	}

	canRemainCoin := sdk.NewCoin(config.BaseDenom, mortgageRemainDec.TruncateInt())
	canRemainCoins := sdk.NewCoins(canRemainCoin)
	if transferType == types.TransferTypeToModule {
		
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accFromAddress, toAddress, canRemainCoins)
		if err != nil {
			return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
		}
	} else if transferType == types.TransferTypeToAccount {
		

		accToAddress, err := sdk.AccAddressFromBech32(toAddress)
		if err != nil {
			return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrAddressFormat
		}

		err = k.bankKeeper.SendCoins(ctx, accFromAddress, accToAddress, canRemainCoins)
		if err != nil {
			return sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)), types.ErrTransfer
		}
	}
	return canRemainCoin, nil
}


func (k Keeper) RegisterMobile(ctx sdk.Context, nodeAddress, fromAddress string) (mobile int64) {
	
	mobile = time.Now().Unix() + 10000000000

	

	return mobile
}
