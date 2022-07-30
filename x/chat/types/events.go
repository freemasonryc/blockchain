package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)


const (


	EventTypeFromAddress    = "from_address"
	EventTypeNodeAddress    = "node_address"
	EventTypeMortGageAmount = "mortgage_amount"
	EventTypeMortGageDenom  = "mortgage_denom"
	EventTypeMortgageRemain = "mortgage_remain"
	EventTypeGetMobile      = "get_mobile"



	EventTypeDevide         = "event_devide"
	EventTypeDevideRegister = "event_devide_register"
	EventTypeDevideMortgate = "event_devide_mortgate"
	EventTypeDevideSendGift = "event_devide_send_gift"



	MortgageEventTypeType           = "mortgage_type"
	MortgageEventTypeFromAddress    = "mortgage_from_address"
	MortgageEventTypeMortgageInfo   = "mortgage_info"
	MortgageEventTypeDenom          = "mortgage_denom"
	MortgageEventTypeMortgageAmount = "mortgage_mortgage_amount"
	MortgageEventTypeFromBalance    = "mortgage_from_balance"
	MortgateEventTypeMortgageRemain = "mortgage_mortgage_remain"



	SendGiftEventTypeFromAddress  = "send_gift_from_address"
	SendGiftEventTypeToAddress    = "send_gift_to_address"
	SendGiftEventTypeGateAddress  = "send_gift_gate_address"
	SendGiftEventTypeGiftId       = "send_gift_gift_id"
	SendGiftEventTypeGiftValue    = "send_gift_gift_value"
	SendGiftEventTypeGiftDenom    = "send_gift_gift_denom"
	SendGiftEventTypeGiftAmount   = "send_gift_gift_amount"
	SendGiftEventTypeGiftValueAll = "send_gift_gift_value_all"
	SendGiftEventTypeGiftReceive  = "send_gift_gift_receive"



	GetRewardEventTypeFromAddress       = "get_rewards_from_address"
	GetRewardEventTypeMortgageAmountAdd = "get_rewards_mortgage_amount_add"
	GetRewardEventTypeMortgageAmountNew = "get_rewards_mortgage_amount_new"
	GetRewardEventTypeDenom             = "get_rewards_denom"
)


type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}
