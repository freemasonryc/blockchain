package keeper

import (
	

	"freemasonry.cc/blockchain/util"
	"freemasonry.cc/blockchain/x/chat/types"
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
		case types.QueryUserInfo: 
			return QueryUserInfo(ctx, req, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

func QueryUserInfo(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	log := util.BuildLog(util.GetFuncName(), util.LmChainKeeper)
	var params types.QueryUserInfoParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		log.WithError(err).Error("UnmarshalJSON")
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	userInfo, err := k.GerRegisterInfo(ctx, params.Address)
	if err != nil {
		return nil, err
	}

	userInfoByte, err := util.Json.Marshal(userInfo)
	if err != nil {
		return nil, err
	}

	return userInfoByte, nil
}
