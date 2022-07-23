package keeper

import (
	"freemasonry.cc/blockchain/core"
	"freemasonry.cc/blockchain/x/comm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	GatewayKey = "gateway_" 

	GatewayNumKey = "gateway_num" 
)

func (k Keeper) SetGateway(ctx sdk.Context, msg types.MsgGatewayRegister, coin sdk.Coin, valAddress string) error {
	kvStore := k.KVHelper(ctx)
	num := coin.Amount.Quo(core.MustRealString2LedgerInt("10000")) 
	if num.LT(sdk.NewInt(int64(len(msg.IndexNumber)))) {
		return types.ErrGatewayNum
	}
	gatewayInfo := types.Gateway{GatewayAddress: valAddress, GatewayName: msg.GatewayName}
	var gatewayNumArray []types.GatewayNumIndex
	for _, val := range msg.IndexNumber {
		
		gatewayNum, isRegister, err := k.GetGatewayNum(ctx, val)
		if err != nil {
			return err
		}
		if !isRegister { 
			return types.ErrGatewayNumber
		}
		if gatewayNum == nil {
			gatewayNum = &types.GatewayNumIndex{
				GatewayAddress: valAddress,
				NumberIndex:    val,
			}
		}
		gatewayNum.Status = 0
		gatewayNum.Validity = 0
		gatewayNumArray = append(gatewayNumArray, *gatewayNum)
	}
	
	gatewayInfo.GatewayNum = gatewayNumArray
	err := kvStore.Set(GatewayKey+valAddress, gatewayInfo)
	if err != nil {
		return err
	}
	return k.SetGatewayNum(ctx, gatewayNumArray)
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


func (k Keeper) GetGatewayInfo(ctx sdk.Context, gatewayAddress string) (*types.Gateway, error) {
	kvStore := k.KVHelper(ctx)
	keys := GatewayKey + gatewayAddress
	if !kvStore.Has(keys) {
		return nil, nil
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
		return nil, false, nil
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
