package keeper

import (
	"context"
	"encoding/base64"
	"freemasonry.cc/blockchain/x/comm/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"time"
)

var _ types.MsgServer = &msgServer{}

type msgServer struct {
	Keeper
	logPrefix string
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper, logPrefix: "comm | msgServer | "}
}


func (k msgServer) GatewayDelegation(goCtx context.Context, msg *types.MsgGatewayDelegate) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	valAddr, valErr := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if valErr != nil {
		return nil, valErr
	}
	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return nil, stakingTypes.ErrNoValidatorFound
	}
	delegatorAddress, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}
	
	if !msg.Amount.IsZero() {
		err = k.delegate(ctx, delegatorAddress, valAddr, validator, msg.Amount)
		if err != nil {
			return nil, err
		}
	}
	
	if validator.GetOperator().String() == sdk.ValAddress(delegatorAddress).String() {
		
		gateway, err := k.GetGatewayInfo(ctx, msg.ValidatorAddress)
		if err != nil {
			if err == types.ErrGatewayNotExist {
				return &types.MsgEmptyResponse{}, nil
			}
			return nil, err
		}
		
		delegation, found := k.stakingKeeper.GetDelegation(ctx, delegatorAddress, valAddr)
		if !found {
			return nil, stakingTypes.ErrNoDelegation
		}
		params := k.GetParams(ctx)
		num := delegation.Shares.QuoInt(params.MinDelegate)
		gateway.GatewayQuota = num.TruncateInt64()
		if len(msg.IndexNumber) > 0 {
			
			if num.Sub(sdk.NewInt(int64(len(gateway.GatewayNum))).ToDec()).LT(sdk.NewDec(int64(len(msg.IndexNumber)))) {
				return nil, types.ErrGatewayNum
			}
			indexNumArray, err := k.GatewayNumFilter(ctx, msg.ValidatorAddress, msg.IndexNumber)
			if err != nil {
				return nil, err
			}
			
			err = k.SetGatewayNum(ctx, indexNumArray)
			if err != nil {
				return nil, err
			}
			
			err = k.GatewayRedeemNumFilter(ctx, indexNumArray)
			if err != nil {
				return nil, err
			}
			gateway.GatewayNum = append(gateway.GatewayNum, indexNumArray...)
		}
		
		err = k.UpdateGatewayInfo(ctx, *gateway)
		if err != nil {
			return nil, err
		}
	}
	return &types.MsgEmptyResponse{}, nil
}


func (k msgServer) GatewayUndelegate(goCtx context.Context, msg *types.MsgGatewayUndelegate) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, err := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		return nil, err
	}
	delegatorAddress, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		return nil, err
	}
	
	delegation, found := k.stakingKeeper.GetDelegation(ctx, delegatorAddress, addr)
	if !found {
		return nil, stakingTypes.ErrNoDelegation
	}
	
	shares, err := k.stakingKeeper.ValidateUnbondAmount(
		ctx, delegatorAddress, addr, msg.Amount.Amount,
	)
	if err != nil {
		return nil, err
	}
	bondDenom := k.stakingKeeper.BondDenom(ctx)
	if msg.Amount.Denom != bondDenom {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest, "invalid coin denomination: got %s, expected %s", msg.Amount.Denom, bondDenom,
		)
	}
	validator, found := k.stakingKeeper.GetValidator(ctx, addr)
	if !found {
		return nil, stakingTypes.ErrNoDelegatorForAddress
	}
	completionTime, returnAmount, err := k.Keeper.Undelegate(ctx, delegatorAddress, addr, validator, shares)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			stakingTypes.EventTypeUnbond,
			sdk.NewAttribute(stakingTypes.AttributeKeyValidator, msg.ValidatorAddress),
			sdk.NewAttribute(sdk.AttributeKeyAmount, msg.Amount.Amount.String()),    
			sdk.NewAttribute(types.AttributeKeyReturnAmount, returnAmount.String()), 
			sdk.NewAttribute(stakingTypes.AttributeKeyCompletionTime, completionTime.Format(time.RFC3339)),

			sdk.NewAttribute(stakingTypes.AttributeKeyDelegatorAddr, delegatorAddress.String()), 
			sdk.NewAttribute(stakingTypes.AttributeKeyNewShares, shares.String()),               
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, stakingTypes.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.DelegatorAddress),
		),
	})

	
	if validator.GetOperator().String() == sdk.ValAddress(delegatorAddress).String() {
		
		gatewayInfo, err := k.GetGatewayInfo(ctx, validator.GetOperator().String())
		if err != nil {
			if err == types.ErrGatewayNumNotFound {
				return &types.MsgEmptyResponse{}, nil
			}
			return nil, err
		}
		params := k.GetParams(ctx)
		
		balanceShares := delegation.Shares.Sub(shares)
		
		num := balanceShares.QuoInt(params.MinDelegate)

		gatewayInfo.GatewayQuota = num.TruncateInt64()
		
		holdNum := int64(len(gatewayInfo.GatewayNum))
		
		if holdNum-int64(len(msg.IndexNumber)) > gatewayInfo.GatewayQuota {
			return nil, types.ErrGatewayNum
		}
		if len(msg.IndexNumber) > 0 {
			var indexNumArray []types.GatewayNumIndex
			for _, val := range msg.IndexNumber {
				
				indexNum, _, err := k.GetGatewayNum(ctx, val)
				if err != nil {
					return nil, err
				}
				if indexNum == nil {
					return nil, types.ErrGatewayNumNotFound
				}
				indexNum.Status = 1 
				indexNum.Validity = indexNum.Validity + params.Validity
				indexNumArray = append(indexNumArray, *indexNum)
				
				for i, gatewayNum := range gatewayInfo.GatewayNum {
					if gatewayNum.NumberIndex == val {
						gatewayInfo.GatewayNum = append(gatewayInfo.GatewayNum[:i], gatewayInfo.GatewayNum[i+1:]...)
					}
				}
			}
			
			err := k.UpdateGatewayInfo(ctx, *gatewayInfo)
			if err != nil {
				return nil, err
			}
			
			err = k.SetGatewayNum(ctx, indexNumArray)
			if err != nil {
				return nil, err
			}
			
			err = k.SetGatewayRedeemNum(ctx, indexNumArray)
			if err != nil {
				return nil, err
			}
		}
		
		err = k.UpdateGatewayInfo(ctx, *gatewayInfo)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgEmptyResponse{}, nil
}


func (k msgServer) GatewayRegister(goCtx context.Context, msg *types.MsgGatewayRegister) (*types.MsgEmptyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	addr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}
	delInt, ok := sdk.NewIntFromString(msg.Delegation)
	if !ok {
		return &types.MsgEmptyResponse{}, sdkerrors.Wrapf(
			types.ErrDelegationCoin, "invalid delegation : got %s", msg.Delegation)
	}
	params := k.GetParams(ctx)
	delegation := sdk.NewCoin(sdk.DefaultBondDenom, delInt)

	valAddress := sdk.ValAddress(addr)
	validator, found := k.stakingKeeper.GetValidator(ctx, valAddress)
	if !found { 
		err := k.createValidator(ctx, addr, valAddress, params, *msg, delegation)
		if err != nil {
			return &types.MsgEmptyResponse{}, err
		}
		
		err = k.SetGatewayDelegateLastTime(ctx, msg.Address, valAddress.String())
		if err != nil {
			return &types.MsgEmptyResponse{}, err
		}
	} else {
		
		if !delegation.IsZero() {
			err = k.delegate(ctx, addr, valAddress, validator, delegation)
			if err != nil {
				return nil, err
			}
		}
		
		delegate, found := k.stakingKeeper.GetDelegation(ctx, addr, valAddress)
		if !found {
			return nil, stakingTypes.ErrNoDelegation
		}
		if delegate.Shares.LT(params.MinDelegate.ToDec()) {
			return &types.MsgEmptyResponse{}, types.ErrGatewayDelegation
		}
		delegation = sdk.NewCoin(sdk.DefaultBondDenom, delegate.Shares.TruncateInt())
	}
	
	err = k.SetGateway(ctx, *msg, delegation, valAddress.String())
	if err != nil {
		return &types.MsgEmptyResponse{}, err
	}
	return &types.MsgEmptyResponse{}, nil
}


func ParseBech32ValConsPubkey(validatorInfoPubKeyBase64 string) (cryptotypes.PubKey, error) {
	validatorInfoPubKeyBytes, err := base64.StdEncoding.DecodeString(validatorInfoPubKeyBase64)
	if err != nil {
		return nil, err
	}
	pbk := ed25519.PubKey(validatorInfoPubKeyBytes) 
	pubkey, err := cryptocodec.FromTmPubKeyInterface(pbk)
	if err != nil {
		return nil, err
	}
	return pubkey, nil
}
