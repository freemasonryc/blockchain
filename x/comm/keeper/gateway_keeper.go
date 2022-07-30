package keeper

import (
	"freemasonry.cc/blockchain/util"
	"freemasonry.cc/blockchain/x/comm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	GatewayKey = "gateway_" 

	GatewayNumKey = "gateway_num" 

	GatewayRedeemNumKey = "gateway_redeem_num" 
)

func (k Keeper) SetGateway(ctx sdk.Context, msg types.MsgGatewayRegister, coin sdk.Coin, valAddress string) error {
	kvStore := k.KVHelper(ctx)
	params := k.GetParams(ctx)
	num := coin.Amount.Quo(params.MinDelegate)
	gatewayInfo := types.Gateway{}
	if kvStore.Has(GatewayKey + valAddress) {
		err := kvStore.GetUnmarshal(GatewayKey+valAddress, &gatewayInfo)
		if err != nil {
			return err
		}
	}
	gatewayInfo.GatewayAddress = valAddress
	gatewayInfo.GatewayName = msg.GatewayName
	gatewayInfo.GatewayUrl = msg.GatewayUrl
	gatewayInfo.GatewayQuota = num.Int64()
	if len(msg.IndexNumber) > 0 {
		if (num.Sub(sdk.NewInt(int64(len(gatewayInfo.GatewayNum))))).LT(sdk.NewInt(int64(len(msg.IndexNumber)))) {
			return types.ErrGatewayNum
		}
		gatewayNumArray, err := k.GatewayNumFilter(ctx, valAddress, msg.IndexNumber)
		if err != nil {
			return err
		}
		gatewayInfo.GatewayNum = append(gatewayInfo.GatewayNum, gatewayNumArray...)
		
		err = k.SetGatewayNum(ctx, gatewayNumArray)
		if err != nil {
			return err
		}
		
		err = k.GatewayRedeemNumFilter(ctx, gatewayNumArray)
		if err != nil {
			return err
		}
	}
	
	return kvStore.Set(GatewayKey+valAddress, gatewayInfo)
}


func (k Keeper) GatewayNumFilter(ctx sdk.Context, validatorAddress string, indexNum []string) ([]types.GatewayNumIndex, error) {
	var gatewayNumArray []types.GatewayNumIndex
	for _, val := range indexNum {
		
		gatewayNum, isRegister, err := k.GetGatewayNum(ctx, val)
		if err != nil {
			return nil, err
		}
		if !isRegister { 
			return nil, types.ErrGatewayNumber
		}
		if gatewayNum == nil {
			gatewayNum = &types.GatewayNumIndex{
				GatewayAddress: validatorAddress,
				NumberIndex:    val,
			}
		}
		gatewayNum.Status = 0
		gatewayNum.Validity = 0
		gatewayNumArray = append(gatewayNumArray, *gatewayNum)
	}
	return gatewayNumArray, nil
}


func (k Keeper) UpdateGatewayInfo(ctx sdk.Context, gateway types.Gateway) error {
	kvStore := k.KVHelper(ctx)
	keys := GatewayKey + gateway.GatewayAddress
	err := kvStore.Set(keys, gateway)
	if err != nil {
		return err
	}
	return nil
}


func (k Keeper) SetGatewayNum(ctx sdk.Context, gatewayNumArray []types.GatewayNumIndex) error {
	kvStore := k.KVHelper(ctx)
	gatewayMap := make(map[string]types.GatewayNumIndex)
	if kvStore.Has(GatewayNumKey) {
		err := kvStore.GetUnmarshal(GatewayNumKey, &gatewayMap)
		if err != nil {
			return err
		}
	}
	for _, val := range gatewayNumArray {
		gatewayMap[val.NumberIndex] = val
	}
	err := kvStore.Set(GatewayNumKey, gatewayMap)
	if err != nil {
		return err
	}
	return nil
}


func (k Keeper) SetGatewayRedeemNum(ctx sdk.Context, gatewayNumArray []types.GatewayNumIndex) error {
	kvStore := k.KVHelper(ctx)
	gatewayMap := make(map[string]types.GatewayNumIndex)
	if kvStore.Has(GatewayRedeemNumKey) {
		err := kvStore.GetUnmarshal(GatewayRedeemNumKey, &gatewayMap)
		if err != nil {
			return err
		}
	}
	for _, val := range gatewayNumArray {
		gatewayMap[val.NumberIndex] = val
	}
	err := kvStore.Set(GatewayRedeemNumKey, gatewayMap)
	if err != nil {
		return err
	}
	return nil
}


func (k Keeper) GatewayRedeemNumFilter(ctx sdk.Context, gatewayNumArray []types.GatewayNumIndex) error {
	kvStore := k.KVHelper(ctx)
	if !kvStore.Has(GatewayRedeemNumKey) {
		return nil
	}
	gatewayNumMap := make(map[string]types.GatewayNumIndex)
	err := kvStore.GetUnmarshal(GatewayRedeemNumKey, &gatewayNumMap)
	if err != nil {
		return err
	}
	for _, val := range gatewayNumArray {
		if _, ok := gatewayNumMap[val.NumberIndex]; ok {
			delete(gatewayNumMap, val.NumberIndex)
		}
	}
	return kvStore.Set(GatewayRedeemNumKey, gatewayNumMap)
}


func (k Keeper) GetGatewayRedeemNum(ctx sdk.Context) (map[string]types.GatewayNumIndex, error) {
	kvStore := k.KVHelper(ctx)
	if !kvStore.Has(GatewayRedeemNumKey) {
		return nil, nil
	}
	gatewayNumMap := make(map[string]types.GatewayNumIndex)
	err := kvStore.GetUnmarshal(GatewayRedeemNumKey, &gatewayNumMap)
	if err != nil {
		return nil, err
	}
	return gatewayNumMap, nil
}


func (k Keeper) GetGatewayList(ctx sdk.Context) ([]types.Gateway, error) {
	store := k.KVHelper(ctx)
	gatewayArray := []types.Gateway{}
	iterator := store.KVStorePrefixIterator(GatewayKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var gateway types.Gateway
		err := util.Json.Unmarshal(iterator.Value(), &gateway)
		if err != nil {
			return nil, err
		}
		if gateway.GatewayAddress == "" {
			continue
		}
		gatewayArray = append(gatewayArray, gateway)
	}
	return gatewayArray, nil
}


func (k Keeper) GetGatewayInfo(ctx sdk.Context, gatewayAddress string) (*types.Gateway, error) {
	kvStore := k.KVHelper(ctx)
	keys := GatewayKey + gatewayAddress
	if !kvStore.Has(keys) {
		return nil, types.ErrGatewayNotExist
	}
	gateway := new(types.Gateway)
	err := kvStore.GetUnmarshal(keys, gateway)
	if err != nil {
		return nil, err
	}
	return gateway, nil
}


func (k Keeper) GetGatewayInfoByNum(ctx sdk.Context, gatewayNum string) (*types.Gateway, error) {
	kvStore := k.KVHelper(ctx)
	if !kvStore.Has(GatewayNumKey) {
		return nil, nil
	}
	
	gatewayNumMap, err := k.GetGatewayNumMap(ctx)
	if err != nil {
		return nil, err
	}
	if gatewayNumMap == nil {
		return nil, nil
	}
	var gatewayAddress string
	if _, ok := gatewayNumMap[gatewayNum]; ok {
		gatewayAddress = gatewayNumMap[gatewayNum].GatewayAddress
	}
	if gatewayAddress == "" {
		return nil, nil
	}
	return k.GetGatewayInfo(ctx, gatewayAddress)
}


func (k Keeper) GetGatewayNumMap(ctx sdk.Context) (map[string]types.GatewayNumIndex, error) {
	kvStore := k.KVHelper(ctx)
	if !kvStore.Has(GatewayNumKey) {
		return nil, nil
	}
	gatewayNumMap := make(map[string]types.GatewayNumIndex)
	err := kvStore.GetUnmarshal(GatewayNumKey, &gatewayNumMap)
	if err != nil {
		return nil, err
	}
	return gatewayNumMap, nil
}


func (k Keeper) GetGatewayNum(ctx sdk.Context, gatewayNum string) (*types.GatewayNumIndex, bool, error) {
	kvStore := k.KVHelper(ctx)
	if !kvStore.Has(GatewayNumKey) {
		return nil, true, nil
	}
	gatewayNumMap := make(map[string]types.GatewayNumIndex)
	err := kvStore.GetUnmarshal(GatewayNumKey, &gatewayNumMap)
	if err != nil {
		return nil, false, err
	}
	if val, ok := gatewayNumMap[gatewayNum]; ok {
		if val.Status != 2 {
			return &val, false, nil
		}
		return &val, true, nil
	}
	return nil, true, nil
}
