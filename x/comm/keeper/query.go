package keeper

import (
	"freemasonry.cc/blockchain/core"
	"freemasonry.cc/blockchain/util"
	"freemasonry.cc/blockchain/x/comm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)
		switch path[0] {
		case types.QueryGatewayInfo:
			return QueryGatewayInfo(ctx, req, k, legacyQuerierCdc)
		case types.QueryGatewayList:
			return QueryGatewayList(ctx, k)
		case types.QueryGatewayNum:
			return QueryGatewayNum(ctx, k)
		case types.QueryValidatorByConsAddress:
			return queryValidatorByConsAddress(ctx, req, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}


func queryValidatorByConsAddress(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	log := core.BuildLog(core.GetFuncName(), core.LmChainKeeper)
	var params types.QueryValidatorByConsAddrParams

	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		log.WithError(err).Error("UnmarshalJSON")
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	validator, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, params.ValidatorConsAddress)
	if !found {
		return nil, stakingTypes.ErrNoValidatorFound
	}
	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, validator)
	if err != nil {
		log.WithError(err).Error("MarshalJSONIndent")
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}


func QueryGatewayList(ctx sdk.Context, k Keeper) ([]byte, error) {
	log := util.BuildLog(util.GetFuncName(), util.LmChainKeeper)
	gatewayList, err := k.GetGatewayList(ctx)
	if err != nil {
		log.WithError(err).Error("GetGatewayList")
		return nil, err
	}
	gatewayListByte, err := util.Json.Marshal(gatewayList)
	if err != nil {
		log.WithError(err).Error("Marshal")
		return nil, err
	}
	return gatewayListByte, nil
}


func QueryGatewayInfo(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	log := util.BuildLog(util.GetFuncName(), util.LmChainKeeper)
	var params types.QueryGatewayInfoParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		log.WithError(err).Error("UnmarshalJSON")
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	var gateway *types.Gateway

	if params.GatewayAddress != "" {
		gateway, err = k.GetGatewayInfo(ctx, params.GatewayAddress)
		if err != nil {
			log.WithError(err).Error("GetGatewayInfo")
			return nil, err
		}
	}

	if params.GatewayNumIndex != "" {
		gateway, err = k.GetGatewayInfoByNum(ctx, params.GatewayNumIndex)
		if err != nil {
			log.WithError(err).Error("GetGatewayInfoByNum")
			return nil, err
		}
	}
	if gateway != nil {
		gatewayByte, err := util.Json.Marshal(gateway)
		if err != nil {
			log.WithError(err).Error("Marshal")
			return nil, err
		}
		return gatewayByte, nil
	}
	return nil, nil
}


func QueryGatewayNum(ctx sdk.Context, k Keeper) ([]byte, error) {
	log := util.BuildLog(util.GetFuncName(), util.LmChainKeeper)
	gatewayMap, err := k.GetGatewayNumMap(ctx)
	if err != nil {
		log.WithError(err).Error("GetGatewayNumMap")
		return nil, err
	}
	gatewayMapByte, err := util.Json.Marshal(gatewayMap)
	if err != nil {
		log.WithError(err).Error("Marshal")
		return nil, err
	}
	return gatewayMapByte, nil
}
