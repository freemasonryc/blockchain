package keeper

import (
	"context"
	types2 "freemasonry.cc/blockchain/x/chat/types"
	"freemasonry.cc/blockchain/x/comm/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakeingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var _ types.MsgServer = &Keeper{}

type msgServer struct {
	Keeper
	logPrefix string
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper, logPrefix: "comm | msgServer | "}
}


func (k Keeper) GatewayRegister(goCtx context.Context, msg *types.MsgGatewayRegister) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}
	delInt, ok := sdk.NewIntFromString(msg.Delegation)
	if !ok {
		return &types.MsgEmptyResponse{}, types.ErrDelegationCoin
	}
	delegation := sdk.NewCoin(sdk.DefaultBondDenom, delInt)
	if delegation.Amount.LT(sdk.NewInt(1000000)) { 

	}

	valAddress := sdk.ValAddress(addr)
	_, found := k.stakingKeeper.GetValidator(ctx, valAddress)
	if !found { 
		pk, ok := msg.Pubkey.GetCachedValue().(cryptotypes.PubKey)
		if !ok {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "Expecting cryptotypes.PubKey, got %T", pk)
		}

		if _, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pk)); found {
			return nil, stakeingTypes.ErrValidatorPubKeyExists
		}
		bondDenom := k.stakingKeeper.BondDenom(ctx)
		if delegation.Denom != bondDenom {
			return nil, sdkerrors.Wrapf(
				sdkerrors.ErrInvalidRequest, "invalid coin denomination: got %s, expected %s", delegation.Denom, bondDenom,
			)
		}

		description := stakeingTypes.NewDescription(msg.GatewayName, "gateway", "", "", "")

		validator, err := stakeingTypes.NewValidator(valAddress, pk, description)
		if err != nil {
			return &types.MsgEmptyResponse{}, err
		}
		commission := stakeingTypes.NewCommissionWithTime(
			msg.Commission.Rate, msg.Commission.MaxRate,
			msg.Commission.MaxChangeRate, ctx.BlockHeader().Time,
		)
		validator, err = validator.SetInitialCommission(commission)
		
		validator.MinSelfDelegation = delegation.Amount
		k.stakingKeeper.SetValidator(ctx, validator)
		k.stakingKeeper.SetValidatorByConsAddr(ctx, validator)
		k.stakingKeeper.SetNewValidatorByPowerIndex(ctx, validator)
		k.stakingKeeper.AfterValidatorCreated(ctx, validator.GetOperator())
		
		_, err = k.stakingKeeper.Delegate(ctx, addr, delegation.Amount, stakeingTypes.Unbonded, validator, true)
		if err != nil {
			return &types.MsgEmptyResponse{}, err
		}
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				stakeingTypes.EventTypeCreateValidator,
				sdk.NewAttribute(stakeingTypes.AttributeKeyValidator, validator.GetOperator().String()),
				sdk.NewAttribute(sdk.AttributeKeyAmount, delegation.String()),
			),
		})
	}
	
	err = k.SetGateway(ctx, *msg, delegation, valAddress.String())
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}
	return &types.MsgEmptyResponse{}, nil
}

func (k Keeper) AddressBookSave(goCtx context.Context, msg *types.MsgAddressBookSave) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	
	_, err := k.chatKeeper.GerRegisterInfo(ctx, msg.FromAddress)
	if err != nil {
		return &types.MsgEmptyResponse{}, types2.ErrUserNotFound
	}

	
	err = k.AddressBookSet(ctx, msg.FromAddress, msg.AddressBook)
	if err != nil {
		return &types.MsgEmptyResponse{}, types.ErrAddressBookSet
	}

	return &types.MsgEmptyResponse{}, nil
}
