package keeper

import (
	"context"
	"encoding/json"
	"freemasonry.cc/blockchain/x/chat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

var _ types.MsgServer = &Keeper{}

type msgServer struct {
	Keeper
	logPrefix string
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper, logPrefix: "chat | msgServer | "}
}


func (k Keeper) MobileTransfer(goCtx context.Context, msg *types.MsgMobileTransfer) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	
	userInfo, err := k.GetRegisterInfo(ctx, msg.GetFromAddress())
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserNotFound
	}

	isHave := false

	for _, mobile := range userInfo.Mobile {
		if mobile == msg.Mobile {
			isHave = true
			break
		}
	}

	if !isHave {
		return &types.MsgEmptyResponse{}, types.ErrUserNotHaveMobile
	}

	
	toUserInfo, err := k.GetRegisterInfo(ctx, msg.GetToAddress())
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserNotFound
	}

	
	userInfoNew := make([]string, 0)
	for _, mobile := range userInfo.Mobile {
		if mobile != msg.Mobile {
			userInfoNew = append(userInfoNew, mobile)
		}
	}
	userInfo.Mobile = userInfoNew
	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserUpdate
	}

	toUserInfo.Mobile = append(toUserInfo.Mobile, msg.Mobile)

	err = k.SetRegisterInfo(ctx, toUserInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserUpdate
	}

	return nil, nil
}


func (k Keeper) GetRewards(goCtx context.Context, msg *types.MsgGetRewards) (*types.MsgEmptyResponse, error) {
	

	ctx := sdk.UnwrapSDKContext(goCtx)

	
	chatParams := k.GetParams(ctx)

	
	reward, err := k.CaculateChatReward(ctx, msg.FromAddress, chatParams.ChatRewardLog)
	if err != nil {
		return nil, err
	}

	
	userinfo, err := k.GetRegisterInfo(ctx, msg.FromAddress)
	if err != nil {
		return nil, err
	}

	CanRedemAmountAdd := reward.Sub(userinfo.CanRedemAmount)

	userinfo.CanRedemAmount = reward

	err = k.SetRegisterInfo(ctx, userinfo)
	if err != nil {
		return nil, err
	}

	
	err = k.SetMortgageLog(ctx, msg.FromAddress, reward)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	
	k.SetLastGetHeight(ctx, msg.FromAddress, ctx.BlockHeight(), userinfo.CanRedemAmount)

	
	ctx.EventManager().EmitEvents(
		[]sdk.Event{
			sdk.NewEvent(
				types.TypeMsgGetRewards,
				sdk.NewAttribute(types.GetRewardEventTypeFromAddress, msg.FromAddress),
				sdk.NewAttribute(types.GetRewardEventTypeMortgageAmountNew, userinfo.CanRedemAmount.Amount.String()),
				sdk.NewAttribute(types.GetRewardEventTypeMortgageAmountAdd, CanRedemAmountAdd.Amount.String()),
				sdk.NewAttribute(types.GetRewardEventTypeDenom, userinfo.CanRedemAmount.Denom),
			),
		},
	)
	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) AddressBookSave(goCtx context.Context, msg *types.MsgAddressBookSave) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	userInfo, err := k.GetRegisterInfo(ctx, msg.GetFromAddress())
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserNotFound
	}

	userInfo.Mobile = msg.AddressBook

	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserUpdate
	}

	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) SendGift(goCtx context.Context, msg *types.MsgSendGift) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	giftValueAll := sdk.NewCoin(msg.GiftValue.Denom, msg.GiftValue.Amount.Mul(sdk.NewInt(msg.GiftAmount)))

	
	mortgateInfo, err := k.MortgageSendCoin(ctx, types.TransferTypeToAccount, msg.ToAddress, msg.FromAddress, msg.NodeAddress, giftValueAll)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	mortgageInfoJson, err := json.Marshal(mortgateInfo.MortgageDevideInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrDevideError
	}

	accFromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrDevideError
	}

	fromBalance := k.bankKeeper.GetAllBalances(ctx, accFromAddress)

	ctx.EventManager().EmitEvents(
		[]sdk.Event{
			sdk.NewEvent(
				types.TypeMsgSendGift,
				sdk.NewAttribute(types.SendGiftEventTypeFromAddress, msg.FromAddress),
				sdk.NewAttribute(types.SendGiftEventTypeToAddress, msg.ToAddress),
				sdk.NewAttribute(types.SendGiftEventTypeGateAddress, msg.NodeAddress),
				sdk.NewAttribute(types.SendGiftEventTypeGiftId, strconv.FormatInt(msg.GiftId, 10)),
				sdk.NewAttribute(types.SendGiftEventTypeGiftValue, msg.GiftValue.Amount.String()),
				sdk.NewAttribute(types.SendGiftEventTypeGiftDenom, msg.GiftValue.Denom),
				sdk.NewAttribute(types.SendGiftEventTypeGiftAmount, strconv.FormatInt(msg.GiftAmount, 10)),
				sdk.NewAttribute(types.SendGiftEventTypeGiftValueAll, giftValueAll.Amount.String()),
				sdk.NewAttribute(types.SendGiftEventTypeGiftReceive, mortgateInfo.MortgageRemain.Amount.String()),
			),
			sdk.NewEvent(
				types.EventTypeDevide,
				sdk.NewAttribute(types.MortgageEventTypeType, types.EventTypeDevideSendGift),
				sdk.NewAttribute(types.MortgageEventTypeFromAddress, msg.FromAddress),
				sdk.NewAttribute(types.MortgageEventTypeMortgageInfo, string(mortgageInfoJson)),
				sdk.NewAttribute(types.MortgageEventTypeDenom, msg.GiftValue.Denom),
				sdk.NewAttribute(types.MortgageEventTypeMortgageAmount, giftValueAll.Amount.String()),
				sdk.NewAttribute(types.MortgageEventTypeFromBalance, fromBalance.String()),
				sdk.NewAttribute(types.MortgateEventTypeMortgageRemain, mortgateInfo.MortgageRemain.Amount.String()),
			),
		},
	)

	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) SetChatFee(goCtx context.Context, msg *types.MsgSetChatFee) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	userInfo, err := k.GetRegisterInfo(ctx, msg.FromAddress)
	
	
	

	
	userInfo.ChatFee = msg.Fee

	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserUpdate
	}

	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.TypeMsgSetChatFee,
			sdk.NewAttribute("from_address", msg.FromAddress),
			sdk.NewAttribute("fee_amount", msg.Fee.Amount.String()),
			sdk.NewAttribute("fee_denom", msg.Fee.Denom),
		),
	)

	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) MortGage(goCtx context.Context, msg *types.MsgMortgage) (*types.MsgEmptyResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	
	userInfo, err := k.GetRegisterInfo(ctx, msg.FromAddress)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserNotFound
	}

	
	mortgateInfo, err := k.MortgageSendCoin(ctx, types.TransferTypeToModule, types.ModuleName, msg.FromAddress, msg.NodeAddress, msg.MortgageAmount)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	
	userInfo.CanRedemAmount = userInfo.CanRedemAmount.Add(mortgateInfo.MortgageRemain)
	userInfo.MortgageAmount = userInfo.MortgageAmount.Add(msg.MortgageAmount)
	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserUpdate
	}

	
	err = k.SetMortgageLog(ctx, msg.FromAddress, userInfo.CanRedemAmount)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	mortgageInfoJson, err := json.Marshal(mortgateInfo.MortgageDevideInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrRegister
	}

	accFromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrDevideError
	}
	fromBalance := k.bankKeeper.GetAllBalances(ctx, accFromAddress)

	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDevide,
			sdk.NewAttribute(types.MortgageEventTypeType, types.EventTypeDevideMortgate),
			sdk.NewAttribute(types.MortgageEventTypeFromAddress, msg.FromAddress),
			sdk.NewAttribute(types.MortgageEventTypeMortgageInfo, string(mortgageInfoJson)),
			sdk.NewAttribute(types.MortgageEventTypeDenom, msg.MortgageAmount.Denom),
			sdk.NewAttribute(types.MortgageEventTypeMortgageAmount, msg.MortgageAmount.Amount.String()),
			sdk.NewAttribute(types.MortgageEventTypeFromBalance, fromBalance.String()),
			sdk.NewAttribute(types.MortgateEventTypeMortgageRemain, mortgateInfo.MortgageRemain.Amount.String()),
		),
	)

	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgEmptyResponse, error) {

	
	//
	

	ctx := sdk.UnwrapSDKContext(goCtx)

	
	_, err := k.GetRegisterInfo(ctx, msg.FromAddress)
	if err == nil {
		return &types.MsgEmptyResponse{}, types.ErrUserHasExisted
	}

	mortgateInfo, err := k.MortgageSendCoin(ctx, types.TransferTypeToModule, types.ModuleName, msg.FromAddress, msg.NodeAddress, msg.MortgageAmount)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	
	mobile, err := k.RegisterMobile(ctx, msg.NodeAddress, msg.FromAddress, msg.MobilePrefix)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	userMobile := []string{
		mobile,
	}

	
	var userInfo types.UserInfo
	userInfo.NodeAddress = msg.NodeAddress
	userInfo.Mobile = userMobile
	userInfo.FromAddress = msg.FromAddress
	userInfo.MortgageAmount = msg.MortgageAmount
	userInfo.CanRedemAmount = mortgateInfo.MortgageRemain
	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrRegister
	}

	
	err = k.SetMortgageLog(ctx, msg.FromAddress, userInfo.CanRedemAmount)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	mortgageInfoJson, err := json.Marshal(mortgateInfo.MortgageDevideInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrRegister
	}

	accFromAddress, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrDevideError
	}
	fromBalance := k.bankKeeper.GetAllBalances(ctx, accFromAddress)

	
	ctx.EventManager().EmitEvents(
		[]sdk.Event{
			sdk.NewEvent(
				types.TypeMsgRegister,
				sdk.NewAttribute(types.EventTypeFromAddress, msg.FromAddress),
				sdk.NewAttribute(types.EventTypeNodeAddress, msg.NodeAddress),
				sdk.NewAttribute(types.EventTypeMortGageAmount, msg.MortgageAmount.Amount.String()),
				sdk.NewAttribute(types.EventTypeMortGageDenom, msg.MortgageAmount.Denom),
				sdk.NewAttribute(types.EventTypeMortgageRemain, mortgateInfo.MortgageRemain.Amount.String()),
				sdk.NewAttribute(types.EventTypeGetMobile, mobile),
			),
			sdk.NewEvent(
				types.EventTypeDevide,
				sdk.NewAttribute(types.MortgageEventTypeType, types.EventTypeDevideRegister),
				sdk.NewAttribute(types.MortgageEventTypeFromAddress, msg.FromAddress),
				sdk.NewAttribute(types.MortgageEventTypeMortgageInfo, string(mortgageInfoJson)),
				sdk.NewAttribute(types.MortgageEventTypeDenom, msg.MortgageAmount.Denom),
				sdk.NewAttribute(types.MortgageEventTypeMortgageAmount, msg.MortgageAmount.Amount.String()),
				sdk.NewAttribute(types.MortgageEventTypeFromBalance, fromBalance.String()),
				sdk.NewAttribute(types.MortgateEventTypeMortgageRemain, mortgateInfo.MortgageRemain.Amount.String()),
			),
		},
	)

	return &types.MsgEmptyResponse{}, nil
}
