package keeper

import (
	"freemasonry.cc/blockchain/util"
	"freemasonry.cc/blockchain/x/comm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		case types.QueryGatewayNum: 
			return QueryGatewayNum(ctx, k)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
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
