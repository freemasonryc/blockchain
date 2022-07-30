package types

import (
	"github.com/cosmos/cosmos-sdk/types"
)

const (
	MsgTypeRegister        = "chat/MsgRegister"
	MsgTypeMortgage        = "chat/MsgMortgage"
	MsgTypeSetChatFee      = "chat/MsgTypeSetChatFee"
	MsgTypeSendGift        = "chat/MsgTypSendGift"
	MsgTypeAddressBookSave = "chat/MsgTypeAddressBookSave"
	MsgTypeGetRewards      = "chat/MsgTypeGetRewards"
	MsgTypeMobileTransfer  = "chat/MsgTypeMobileTransfer"
)


type UserInfo struct {
	FromAddress    string     `json:"from_address" yaml:"from_address"`
	NodeAddress    string     `json:"node_address" yaml:"node_address"`
	MortgageAmount types.Coin `json:"mortgage_amount" yaml:"mortgage_amount"`
	CanRedemAmount types.Coin `json:"can_redem_amount" yaml:"can_redem_amount"`
	Mobile         []string   `json:"mobile" yaml:"mobile"`
	ChatFee        types.Coin `json:"chat_fee" yaml:"chat_fee"`
}

const (
	TransferTypeToModule  = "to_module"
	TransferTypeToAccount = "to_account"
)










type MortgageInfo struct {
	MortgageRemain     types.Coin           `json:"mortgage_remain"`      
	MortgageDevideInfo []MortgageDevideInfo `json:"mortgage_devide_info"` 
}

type MortgageDevideInfo struct {
	MortgageAddress string `json:"mortgage_address"` 
	MortgageAmount  string `json:"mortgage_amount"`  
	ShowBalance     bool   `json:"show_balance"`     
}

type LastReceiveLog struct {
	Height int64      `json:"height"`
	Value  types.Coin `json:"value"`
}

type MortgageAddLog struct {
	Height        int64      `json:"height"`
	MortgageValue types.Coin `json:"mortgage_value"`
}
