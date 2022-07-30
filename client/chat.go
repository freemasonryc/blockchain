package client

import (
	"freemasonry.cc/blockchain/util"
	"freemasonry.cc/blockchain/x/chat/types"
	"github.com/sirupsen/logrus"
)

type ChatInfo struct {
	TxClient      *TxClient
	AccountClient *AccountClient
	ServerUrl     string
	logPrefix     string
}

type ChatClient struct {
	TxClient      *TxClient
	AccountClient *AccountClient
	ServerUrl     string
	logPrefix     string
}










//





//



//








//








//




//




//



type GetUserInfo struct {
	Status   int            `json:"status"`
	Message  string         `json:"message"`
	UserInfo types.UserInfo `json:"user_info"`
}


func (this *ChatClient) QueryUserInfo(address string) (data *GetUserInfo, err error) {
	log := util.BuildLog(util.GetStructFuncName(this), util.LmChainClient).WithFields(logrus.Fields{"address": address})
	params := types.QueryUserInfoParams{Address: address}
	bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		log.WithError(err).Error("MarshalJSON")
		return nil, err
	}
	resBytes, _, err := clientCtx.QueryWithData("custom/chat/"+types.QueryUserInfo, bz)
	if err != nil {
		log.WithError(err).Error("QueryWithData")
		return nil, err
	}
	userInfo := &types.UserInfo{}
	if resBytes != nil {
		err = util.Json.Unmarshal(resBytes, userInfo)
		if err != nil {
			log.WithError(err).Error("Unmarshal")
			return nil, err
		}
	}

	data = &GetUserInfo{}
	if userInfo.FromAddress == "" {
		data.Status = 0
		return
	}
	data.Status = 1
	data.UserInfo.Mobile = userInfo.Mobile
	data.UserInfo.NodeAddress = userInfo.NodeAddress
	data.UserInfo.FromAddress = userInfo.FromAddress
	data.UserInfo.MortgageAmount = userInfo.MortgageAmount
	data.UserInfo.CanRedemAmount = userInfo.CanRedemAmount
	data.UserInfo.ChatFee = userInfo.ChatFee

	return
}
