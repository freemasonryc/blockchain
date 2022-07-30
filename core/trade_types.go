package core

import (
	"freemasonry.cc/trerr"
)



var (
	TradeTypeTransfer            = RegisterTranserType("transfer", "轉帳", "transfer accounts")
	TradeTypeDelegation          = RegisterTranserType("bonded", "POS抵押", "POS mortgage")
	TradeTypeDelegationFee       = RegisterTranserType("bonded-fee", "POS抵押手續費", "POS mortgage service charge")
	TradeTypeUnbondDelegation    = RegisterTranserType("unbonded", "POS贖回", "POS redemption")
	TradeTypeUnbondDelegationFee = RegisterTranserType("unbonded-fee", "POS贖回手續費", "POS redemption fee")
	TradeTypeFee                 = RegisterTranserType("fee", "手續費支出", "Service charge expenditure")
	TradeTypeDelegationReward    = RegisterTranserType("delegate-reward", "POS抵押獎勵", "POS mortgage reward")
	TradeTypeCommissionReward    = RegisterTranserType("commission-reward", "POS傭金獎勵", "POS Commission reward")
	TradeTypeCommunityReward     = RegisterTranserType("community-reward", "社區獎勵", "Community reward")
	TradeTypeValidatorUnjail     = RegisterTranserType("unjail", "POS解除監禁", "POS release")
	TradeTypeValidatorMinerBonus = RegisterTranserType("validator-miner-bonus", "POS激勵", "POS incentive")
	TradeTypeValidatorCreate     = RegisterTranserType("validator-create", "創建驗證器", "Create validator")
	TradeTypeValidatorEditor     = RegisterTranserType("validator-editor", "修改驗證器", "Editor validator")
	TradeTypeCrossChainOut       = RegisterTranserType("cross-chain-out", "跨鏈轉出", "Cross Chain Out")
	TradeTypeCrossChainFee       = RegisterTranserType("cross-chain-fee", "跨鏈手續費", "Cross Chain Fee")
	TradeTypeCrossChainIn        = RegisterTranserType("cross-chain-in", "跨鏈轉入", "Cross Chain In")
	TradeTypeGatewayRegister     = RegisterTranserType("gateway-register", "網關註冊", "Gateway register")
	TradeTypeChatRegister        = RegisterTranserType("event_devide_register", "聊天注册分成", "Chat register share")
	TradeTypeChatMortgage        = RegisterTranserType("event_devide_mortgate", "聊天质押分成", "Chat mortgage share")
	TradeTypeChatSendGift        = RegisterTranserType("event_devide_send_gift", "聊天送礼物分成", "Chat send gift share")
)

var tradeTypeText = make(map[string]string)
var tradeTypeTextEn = make(map[string]string)


func RegisterTranserType(key, value, enValue string) TranserType {
	tradeTypeTextEn[key] = enValue
	tradeTypeText[key] = value
	return TranserType(key)
}


func GetTranserTypeConfig() map[string]string {
	if trerr.Language == "EN" {
		return tradeTypeTextEn
	} else {
		return tradeTypeText
	}
}

type TranserType string

func (this TranserType) GetValue() string {
	if text, ok := tradeTypeText[string(this[:])]; ok {
		return text
	} else {
		return ""
	}
}

func (this TranserType) GetKey() string {
	return string(this[:])
}
