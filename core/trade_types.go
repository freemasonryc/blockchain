package core

import (
	"freemasonry.cc/trerr"
)





var (
	TradeTypeTransfer               = RegisterTranserType("transfer", "轉帳", "transfer accounts")
	TradeTypeCopyrightBuy           = RegisterTranserType("down", "版權購買", "Copyright purchase")
	TradeTypeCopyrightPublish       = RegisterTranserType("publish", "版權鑄造費", "Copyright foundry fee")
	TradeTypeCopyrightEdit          = RegisterTranserType("edit", "修改版權", "Modify copyright")
	TradeTypeCopyrightSharesReward  = RegisterTranserType("bonus", "硬碟頻寬分享獎勵", "Hard disk bandwidth sharing reward")
	TradeTypeCopyrightDelete        = RegisterTranserType("delete", "删除版權", "Delete copyright")
	TradeTypeCopyrightSell          = RegisterTranserType("copyright-share", "版權銷售收入", "Copyright sales revenue")
	TradeTypeCopyrightBuyMortgage   = RegisterTranserType("mining-mortg", "版權抵押購買", "Copyright mortgage purchase")
	TradeTypeCopyrightBuyRedeem     = RegisterTranserType("mining-redeem", "版權抵押購買贖回", "Copyright mortgage purchase and redemption")
	TradeTypeCopyrightPartyRegister = RegisterTranserType("bind-pay", "版權方注册", "Copyright registration")
	TradeTypeDelegation             = RegisterTranserType("bonded", "POS抵押", "POS mortgage")
	TradeTypeDelegationFee          = RegisterTranserType("bonded-fee", "POS抵押手續費", "POS mortgage service charge")
	TradeTypeUnbondDelegation       = RegisterTranserType("unbonded", "POS贖回", "POS redemption")
	TradeTypeUnbondDelegationFee    = RegisterTranserType("unbonded-fee", "POS贖回手續費", "POS redemption fee")
	TradeTypeFee                    = RegisterTranserType("fee", "手續費支出", "Service charge expenditure")
	TradeTypeDelegationReward       = RegisterTranserType("delegate-reward", "POS抵押獎勵", "POS mortgage reward")
	TradeTypeCommissionReward       = RegisterTranserType("commission-reward", "POS傭金獎勵", "POS Commission reward")
	TradeTypeCommunityReward        = RegisterTranserType("community-reward", "社區獎勵", "Community reward")
	TradeTypeValidatorUnjail        = RegisterTranserType("unjail", "POS解除監禁", "POS release")
	TradeTypeNftTransfer            = RegisterTranserType("nft-transfer", "nft轉帳", "NFT transfer")
	TradeTypeCopyrightComplain      = RegisterTranserType("cp-complain", "發起版權申訴", "Copyright complaint")
	TradeTypeSpaceMinerBonus        = RegisterTranserType("space-miner-bonus", "區塊鏈空間激勵", "Blockchain spatial incentive")
	TradeTypeCopyrightComplainReply = RegisterTranserType("cp-complain-reply", "版權應訴", "Copyright response")
	TradeTypeCopyrightComplainVotes = RegisterTranserType("cp-complain-votes", "版權申訴投票", "Copyright appeal vote")
	TradeTypeValidatorMinerBonus    = RegisterTranserType("validator-miner-bonus", "POS激勵", "POS incentive")
	TradeTypeValidatorCreate        = RegisterTranserType("validator-create", "創建驗證器", "Create validator")
	TradeTypeValidatorEditor        = RegisterTranserType("validator-editor", "修改驗證器", "Editor validator")
	TradeTypeInviteReward           = RegisterTranserType("invite-reward", "打赏擴容", "Reward expansion")
	TradeTypeCopyrightVotes         = RegisterTranserType("cp-votes", "社區治理投票", "Community governance voting")
	TradeTypeCopyrightVoteReward    = RegisterTranserType("copyright-vote-reward", "社區治理獎勵", "Community governance incentives")
	TradeTypeUsdtForFsv             = RegisterTranserType("usdt-for-fsv", "usdt兌fsv", "Usdt versus FSV")
	TradeTypeCrossChainOut          = RegisterTranserType("cross-chain-out", "跨鏈轉出", "Cross Chain Out")
	TradeTypeCrossChainFee          = RegisterTranserType("cross-chain-fee", "跨鏈手續費", "Cross Chain Fee")
	TradeTypeCrossChainIn           = RegisterTranserType("cross-chain-in", "跨鏈轉入", "Cross Chain In")
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
