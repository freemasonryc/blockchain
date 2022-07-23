package keeper

import (
	"context"
	"fmt"
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

func (k Keeper) Test(context.Context, *types.MsgTest) (*types.MsgTestResponse, error) {
	return &types.MsgTestResponse{}, nil
}

func (k Keeper) SendGift(goCtx context.Context, msg *types.MsgSendGift) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	giftValueAll := sdk.NewCoin(msg.GiftValue.Denom, msg.GiftValue.Amount.Mul(sdk.NewInt(msg.GiftAmount)))

	
	canRemainCoin, err := k.MortgageSendCoin(ctx, types.TransferTypeToAccount, msg.ToAddress, msg.FromAddress, msg.NodeAddress, giftValueAll)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"send_gift",
			sdk.NewAttribute("from_address", msg.FromAddress),
			sdk.NewAttribute("to_address", msg.ToAddress),
			sdk.NewAttribute("node_address", msg.NodeAddress),
			sdk.NewAttribute("gift_id", strconv.FormatInt(msg.GiftId, 10)),
			sdk.NewAttribute("gift_value", msg.GiftValue.Amount.String()),
			sdk.NewAttribute("gift_amount", strconv.FormatInt(msg.GiftAmount, 10)),
			sdk.NewAttribute("gift_value_all", giftValueAll.Amount.String()),
			sdk.NewAttribute("gift_receive", canRemainCoin.Amount.String()),
			sdk.NewAttribute("gift_value_denom", msg.GiftValue.Denom),
		),
	)

	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) SetChatFee(goCtx context.Context, msg *types.MsgSetChatFee) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	userInfo, err := k.GerRegisterInfo(ctx, msg.FromAddress)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserNotFound
	}

	
	userInfo.ChatFee = msg.Fee

	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserUpdate
	}

	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"set_chat_fee",
			sdk.NewAttribute("from_address", msg.FromAddress),
			sdk.NewAttribute("fee_amount", msg.Fee.Amount.String()),
			sdk.NewAttribute("fee_denom", msg.Fee.Denom),
		),
	)

	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) MortGage(goCtx context.Context, msg *types.MsgMortgage) (*types.MsgEmptyResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	
	userInfo, err := k.GerRegisterInfo(ctx, msg.FromAddress)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserNotFound
	}

	
	canRemainCoin, err := k.MortgageSendCoin(ctx, types.TransferTypeToModule, types.ModuleName, msg.FromAddress, msg.NodeAddress, msg.MortgageAmount)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	
	userInfo.CanRedemAmount = userInfo.CanRedemAmount.Add(canRemainCoin)
	userInfo.MortgageAmount = userInfo.MortgageAmount.Add(msg.MortgageAmount)
	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrUserUpdate
	}

	
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"mortgage",
			sdk.NewAttribute("from_address", msg.FromAddress),
			sdk.NewAttribute("node_address", msg.NodeAddress),
			sdk.NewAttribute("mortgage_amount", msg.MortgageAmount.Amount.String()),
			sdk.NewAttribute("mortgage_denom", msg.MortgageAmount.Denom),
			sdk.NewAttribute("mortgage_remain", canRemainCoin.Amount.String()),
		),
	)

	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) Register(goCtx context.Context, msg *types.MsgRegister) (*types.MsgEmptyResponse, error) {

	fmt.Println("test----------------------------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println(msg.NodeAddress, msg.FromAddress, msg.MortgageAmount)
	

	ctx := sdk.UnwrapSDKContext(goCtx)

	
	_, err := k.GerRegisterInfo(ctx, msg.FromAddress)
	if err == nil {
		return &types.MsgEmptyResponse{}, types.ErrUserHasExisted
	}

	canRemainCoin, err := k.MortgageSendCoin(ctx, types.TransferTypeToModule, types.ModuleName, msg.FromAddress, msg.NodeAddress, msg.MortgageAmount)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}

	
	mobile := k.RegisterMobile(ctx, msg.NodeAddress, msg.FromAddress)
	userMobile := []int64{
		mobile,
	}

	
	var userInfo types.UserInfo
	userInfo.NodeAddress = msg.NodeAddress
	userInfo.Mobile = userMobile
	userInfo.FromAddress = msg.FromAddress
	userInfo.MortgageAmount = msg.MortgageAmount
	userInfo.CanRedemAmount = canRemainCoin
	err = k.SetRegisterInfo(ctx, userInfo)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrRegister
	}

	
	ctx.EventManager().EmitEvents(
		[]sdk.Event{
			sdk.NewEvent(
				"register",
				sdk.NewAttribute("from_address", msg.FromAddress),
				sdk.NewAttribute("node_address", msg.NodeAddress),
				sdk.NewAttribute("mortgage_amount", msg.MortgageAmount.Amount.String()),
				sdk.NewAttribute("mortgage_denom", msg.MortgageAmount.Denom),
				sdk.NewAttribute("mortgage_remain", canRemainCoin.Amount.String()),
				sdk.NewAttribute("get_mobile", sdk.NewInt(mobile).String()),
			),
			sdk.NewEvent(
				"mortgage",
				sdk.NewAttribute("from_address", msg.FromAddress),
				sdk.NewAttribute("node_address", msg.NodeAddress),
				sdk.NewAttribute("mortgage_amount", msg.MortgageAmount.Amount.String()),
				sdk.NewAttribute("mortgage_denom", msg.MortgageAmount.Denom),
				sdk.NewAttribute("mortgage_remain", canRemainCoin.Amount.String()),
			),
		},
	)

	return &types.MsgEmptyResponse{}, nil
}
