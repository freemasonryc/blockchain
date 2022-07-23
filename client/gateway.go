package client

import (
	"freemasonry.cc/blockchain/util"
	"freemasonry.cc/blockchain/x/comm/types"
	"github.com/sirupsen/logrus"
)

type GatewayClient struct {
	ServerUrl string
	logPrefix string
}


func (this *GatewayClient) QueryGateway(gatewayAddress, gatewayNum string) (data *types.Gateway, err error) {
	log := util.BuildLog(util.GetStructFuncName(this), util.LmChainClient).WithFields(logrus.Fields{"gatewayAddress": gatewayAddress})
	params := types.QueryGatewayInfoParams{GatewayAddress: gatewayAddress, GatewayNumIndex: gatewayNum}
	bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		log.WithError(err).Error("MarshalJSON")
		return nil, err
	}
	resBytes, _, err := clientCtx.QueryWithData("custom/comm/"+types.QueryGatewayInfo, bz)
	if err != nil {
		log.WithError(err).Error("QueryWithData")
		return nil, err
	}
	if resBytes != nil {
		err := util.Json.Unmarshal(resBytes, data)
		if err != nil {
			return nil, err
		}
	}
	return
}
