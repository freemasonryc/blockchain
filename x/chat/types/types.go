package types

import (
	"freemasonry.cc/blockchain/core"
	"github.com/cosmos/cosmos-sdk/types"
)

const (
	MsgTypeRegister   = "chat/MsgRegister"
	MsgTypeMortgage   = "chat/MsgMortgage"
	MsgTypeSetChatFee = "chat/MsgTypeSetChatFee"
	MsgTypeSendGift   = "chat/MsgTypSendGift"
)


type UserInfo struct {
	FromAddress    string     `json:"from_address" yaml:"from_address"`
	NodeAddress    string     `json:"node_address" yaml:"node_address"`
	MortgageAmount types.Coin `json:"mortgage_amount" yaml:"mortgage_amount"`
	CanRedemAmount types.Coin `json:"can_redem_amount" yaml:"can_redem_amount"`
	Mobile         []int64    `json:"mobile" yaml:"mobile"`
	ChatFee        types.Coin `json:"chat_fee" yaml:"chat_fee"`
}

const (
	TransferTypeToModule  = "to_module"
	TransferTypeToAccount = "to_account"
)

type RegisterData struct {
	core.TxBase
	core.TradeBase
	FromAddress    string     `json:"from_address"`    
	NodeAddress    string     `json:"node_address"`    
	MortgageAmount types.Coin `json:"mortgage_amount"` 
}

type TranserType string
