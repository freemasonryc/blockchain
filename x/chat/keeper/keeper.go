package keeper

import (
	"errors"
	"fmt"
	"freemasonry.cc/blockchain/cmd/config"
	"freemasonry.cc/blockchain/core"
	commkeeper "freemasonry.cc/blockchain/x/comm/keeper"
	types2 "freemasonry.cc/blockchain/x/comm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	"strconv"
	"strings"

	"freemasonry.cc/blockchain/x/chat/types"
)


type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	commKeeper    commkeeper.Keeper
}


func NewKeeper(
	storeKey sdk.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	bk types.BankKeeper,
	cm commkeeper.Keeper,
) Keeper {
	
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: ak,
		bankKeeper:    bk,
		commKeeper:    cm,
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

func (k Keeper) GetRegisterInfo(ctx sdk.Context, fromAddress string) (userinfo types.UserInfo, err error) {

	store := k.KVHelper(ctx)
	key := types.KeyPrefixRegisterInfo + fromAddress
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



func (k Keeper) MortgageSendCoin(ctx sdk.Context, transferType, toAddress, fromAddress, nodeAddress string, mortgageAmount sdk.Coin) (*types.MortgageInfo, error) {
	log := core.BuildLog(core.GetStructFuncName(k), core.LmChainKeeper)
	
	chatParams := k.GetParams(ctx)

	log.Info("CommunityAddress:", chatParams.CommunityAddress)
	log.Info("EcologicalAddress:", chatParams.EcologicalAddress)
	log.Info("MinMortgageCoin:", chatParams.MinMortgageCoin)

	if mortgageAmount.Amount.Sub(chatParams.MinMortgageCoin.Amount).LT(sdk.NewInt(0)) {
		return nil, types.ErrMortgageAmount
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
		return nil, types.ErrAddressFormat
	}

	
	valAddr, err := sdk.ValAddressFromBech32(nodeAddress)
	if err != nil {
		return nil, types.ErrAddressFormat
	}

	accNodeAddress := sdk.AccAddress(valAddr)
	if err != nil {
		return nil, types.ErrAddressFormat
	}

	
	accCommunityAddress, err := sdk.AccAddressFromBech32(chatParams.CommunityAddress)
	if err != nil {
		return nil, types.ErrAddressFormat
	}
	
	accEcologicalAddress, err := sdk.AccAddressFromBech32(chatParams.EcologicalAddress)
	if err != nil {
		return nil, types.ErrAddressFormat
	}
	
	
	toNodeCoin := sdk.NewCoin(config.BaseDenom, mortgageNodeDec.TruncateInt())
	toNodeCoins := sdk.NewCoins(toNodeCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accNodeAddress, toNodeCoins)
	if err != nil {
		return nil, types.ErrTransfer
	}

	
	toCommunityCoin := sdk.NewCoin(config.BaseDenom, mortgageCommunityDec.TruncateInt())
	toCommunityCoins := sdk.NewCoins(toCommunityCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accCommunityAddress, toCommunityCoins)
	if err != nil {
		return nil, types.ErrTransfer
	}

	
	toEcologicalCoin := sdk.NewCoin(config.BaseDenom, mortgageEcologicalDec.TruncateInt())
	toEcologicalCoins := sdk.NewCoins(toEcologicalCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accEcologicalAddress, toEcologicalCoins)
	if err != nil {
		return nil, types.ErrTransfer
	}

	
	toPosCoin := sdk.NewCoin(config.BaseDenom, mortgagePosDec.TruncateInt())
	toPosCoins := sdk.NewCoins(toPosCoin)
	err = k.bankKeeper.SendCoins(ctx, accFromAddress, accPosAddress, toPosCoins)
	if err != nil {
		return nil, types.ErrTransfer
	}

	
	burnCoin := sdk.NewCoin(config.BaseDenom, mortgageBurnDec.TruncateInt())
	burnCoins := sdk.NewCoins(burnCoin)
	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accFromAddress, types.ModuleBurnName, burnCoins)
	if err != nil {
		return nil, types.ErrTransfer
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleBurnName, burnCoins)
	if err != nil {
		return nil, types.ErrTransfer
	}

	canRemainCoin := sdk.NewCoin(config.BaseDenom, mortgageRemainDec.TruncateInt())
	canRemainCoins := sdk.NewCoins(canRemainCoin)
	if transferType == types.TransferTypeToModule {
		
		err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, accFromAddress, toAddress, canRemainCoins)
		if err != nil {
			return nil, types.ErrTransfer
		}
	} else if transferType == types.TransferTypeToAccount {
		

		accToAddress, err := sdk.AccAddressFromBech32(toAddress)
		if err != nil {
			return nil, types.ErrAddressFormat
		}

		err = k.bankKeeper.SendCoins(ctx, accFromAddress, accToAddress, canRemainCoins)
		if err != nil {
			return nil, types.ErrTransfer
		}
	}

	
	var mortgageinfo types.MortgageInfo
	mortgageinfo.MortgageDevideInfo = make([]types.MortgageDevideInfo, 0)

	log.Info("canRemainCoin:", canRemainCoin)
	log.Info("mortgageDevideInfo:", len(mortgageinfo.MortgageDevideInfo))
	mortgageinfo.MortgageRemain = canRemainCoin

	
	if transferType == types.TransferTypeToModule {
		toAddress = core.ContractChat.String()
	}

	mortgageinfo.MortgageDevideInfo = append(mortgageinfo.MortgageDevideInfo, types.MortgageDevideInfo{
		MortgageAddress: toAddress,
		MortgageAmount:  canRemainCoin.String(),
		ShowBalance:     true,
	})

	
	mortgageinfo.MortgageDevideInfo = append(mortgageinfo.MortgageDevideInfo, types.MortgageDevideInfo{
		MortgageAddress: chatParams.CommunityAddress,
		MortgageAmount:  toCommunityCoin.String(),
	})

	
	mortgageinfo.MortgageDevideInfo = append(mortgageinfo.MortgageDevideInfo, types.MortgageDevideInfo{
		MortgageAddress: accNodeAddress.String(),
		MortgageAmount:  toNodeCoin.String(),
	})
	
	mortgageinfo.MortgageDevideInfo = append(mortgageinfo.MortgageDevideInfo, types.MortgageDevideInfo{
		MortgageAddress: chatParams.EcologicalAddress,
		MortgageAmount:  toEcologicalCoin.String(),
	})
	
	mortgageinfo.MortgageDevideInfo = append(mortgageinfo.MortgageDevideInfo, types.MortgageDevideInfo{
		MortgageAddress: accPosAddress.String(),
		MortgageAmount:  toPosCoin.String(),
	})
	
	mortgageinfo.MortgageDevideInfo = append(mortgageinfo.MortgageDevideInfo, types.MortgageDevideInfo{
		MortgageAddress: core.ContractChatBurn.String(),
		MortgageAmount:  burnCoin.String(),
	})

	return &mortgageinfo, nil
}


func (k Keeper) RegisterMobile(ctx sdk.Context, nodeAddress, fromAddress, mobilePrefix string) (mobile string, err error) {

	log := core.BuildLog(core.GetFuncName(), core.LmChainKeeper)

	log.Info("StoreKey:", k.storeKey)

	
	GatewayNumInfo, _, err := k.commKeeper.GetGatewayNum(ctx, mobilePrefix)

	if err != nil {
		return mobile, types.ErrGeneratingMobile
	}

	if GatewayNumInfo == nil {
		return mobile, types.ErrGetMobile
	}

	if GatewayNumInfo.GatewayAddress != nodeAddress || GatewayNumInfo.Status != 0 {
		return mobile, types.ErrGateway
	}

	
	mobileSuffixInt := len(GatewayNumInfo.NumberEnd)

	log.Info("len(GatewayNumInfo.NumberEnd):", mobileSuffixInt)

	if mobileSuffixInt >= types.MobileSuffixMax {
		return mobile, types.ErrGetMobile
	}

	
	mobileSuffixString := strconv.Itoa(mobileSuffixInt)

	
	mobileSuffixString = strings.Repeat("0", types.MobileSuffixLength-len(mobileSuffixString)) + mobileSuffixString

	mobile = mobilePrefix + mobileSuffixString

	

	GatewayNumInfo.NumberEnd = append(GatewayNumInfo.NumberEnd, mobile)
	GatewayNumInfos := []types2.GatewayNumIndex{
		*GatewayNumInfo,
	}
	err = k.commKeeper.SetGatewayNum(ctx, GatewayNumInfos)
	if err != nil {
		return mobile, types.ErrMobileSetError
	}

	return mobile, nil
}


/*

*/
func (k Keeper) CaculateChatReward(ctx sdk.Context, fromAddress string, chatRewardChangeLog []types.ChatReward) (reward sdk.Coin, err error) {

	
	mortgageLog, err := k.GetMortgageLog(ctx, fromAddress)
	if err != nil {
		return reward, err
	}

	
	commParams := k.commKeeper.GetParams(ctx)

	HeightPerDay := commParams.BonusCycle

	
	LastGetInfo, err := k.GetLastGetHeight(ctx, fromAddress)
	if err != nil {
		return reward, err
	}

	
	NowHeight := ctx.BlockHeight()

	newAmountDec := LastGetInfo.Value.Amount.ToDec()
	
	
	for SendBonusGeight := GetCommonMultipleGt(LastGetInfo.Height, HeightPerDay); SendBonusGeight < NowHeight; SendBonusGeight += HeightPerDay {

		
		var ratio sdk.Dec
		for changeIndex, chatReward := range chatRewardChangeLog {

			
			if SendBonusGeight >= chatReward.Height {
				if HasIndex(changeIndex+1, chatRewardChangeLog) { 
					if SendBonusGeight >= chatRewardChangeLog[changeIndex+1].Height { 
						continue
					} else {
						ratio, err = sdk.NewDecFromStr(chatReward.Value) 
						if err != nil {
							return reward, types.ErrGetBonus
						}
					}
				} else { 
					ratio, err = sdk.NewDecFromStr(chatReward.Value)
					if err != nil {
						return reward, types.ErrGetBonus
					}
				}
			}

		}

		
		mortgageAmount := sdk.NewCoin(mortgageLog[0].MortgageValue.Denom, sdk.ZeroInt()) 
		for _, addLog := range mortgageLog {
			if addLog.Height < SendBonusGeight {
				mortgageAmount = addLog.MortgageValue
			}
		}

		
		newAmountDec = newAmountDec.Add(mortgageAmount.Amount.ToDec().Mul(ratio))
	}

	return sdk.NewCoin(LastGetInfo.Value.Denom, newAmountDec.TruncateInt()), nil
}


func GetCommonMultipleGt(a, b int64) int64 {
	return (a/b + 1) * b
}


func HasIndex(index int, data []types.ChatReward) bool {
	return len(data) > index
}

func (k Keeper) GetLastGetHeight(ctx sdk.Context, fromAddress string) (types.LastReceiveLog, error) {
	store := k.KVHelper(ctx)
	key := types.KeyPrefixLastGetRewardLog + fromAddress

	log := types.LastReceiveLog{
		Height: 1,
		Value:  sdk.NewCoin(config.BaseDenom, sdk.NewInt(0)),
	}

	if store.Has(key) {

		err := store.GetUnmarshal(key, &log)
		if err != nil {
			return log, types.ErrGetLastReveiveHeight
		}
		return log, nil
	}

	return log, nil
}

func (k Keeper) SetLastGetHeight(ctx sdk.Context, fromAddress string, height int64, mortgage sdk.Coin) error {

	store := k.KVHelper(ctx)
	key := types.KeyPrefixLastGetRewardLog + fromAddress

	err := store.Set(key, types.LastReceiveLog{
		Height: height,
		Value:  mortgage,
	})

	if err != nil {
		return types.ErrSetLastReveiveHeight
	}

	return nil
}

func (k Keeper) SetMortgageLog(ctx sdk.Context, fromAddress string, mortgageNew sdk.Coin) error {

	store := k.KVHelper(ctx)
	key := types.KeyPrefixMortgageAddLog + fromAddress

	logNew := types.MortgageAddLog{
		Height:        ctx.BlockHeight(),
		MortgageValue: mortgageNew,
	}
	log := make([]types.MortgageAddLog, 0)

	if store.Has(key) {

		err := store.GetUnmarshal(key, &log)
		if err != nil {
			return types.ErrSetMortgageLog
		}

		
		mortgageAddLogCheck := k.CheckMortgageAddLog(ctx, log, logNew)
		if !mortgageAddLogCheck {
			return types.ErrSetMortgageLog
		}

	}

	log = append(log, logNew)
	err := store.Set(key, log)
	if err != nil {
		return types.ErrSetMortgageLog
	}

	return nil
}

func (k Keeper) GetMortgageLog(ctx sdk.Context, fromAddress string) ([]types.MortgageAddLog, error) {
	store := k.KVHelper(ctx)
	key := types.KeyPrefixMortgageAddLog + fromAddress

	log := make([]types.MortgageAddLog, 0)
	if store.Has(key) {
		err := store.GetUnmarshal(key, &log)
		if err != nil {
			return log, types.ErrGetMortgageLog
		}
	} else {
		return log, types.ErrUserNotFound
	}
	return log, nil
}

func (k Keeper) CheckMortgageAddLog(ctx sdk.Context, log []types.MortgageAddLog, logNew types.MortgageAddLog) bool {

	lastMortgageInfo := log[len(log)-1]
	
	if lastMortgageInfo.Height <= logNew.Height {
		return false
	}

	
	if lastMortgageInfo.MortgageValue.Amount.LT(logNew.MortgageValue.Amount) {
		return false
	}

	return true
}
