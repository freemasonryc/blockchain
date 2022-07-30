package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)


const (
	EventTypeTokenLock             = "token_lock"
	EventTypeTokenUnlock           = "token_unlock"
	EventTypeMint                  = "mint"
	EventTypeConvertCoin           = "convert_coin"
	EventTypeConvertERC20          = "convert_erc20"
	EventTypeBurn                  = "burn"
	EventTypeRegisterCoin          = "register_coin"
	EventTypeRegisterERC20         = "register_erc20"
	EventTypeToggleTokenConversion = "toggle_token_conversion" 

	AttributeKeyCosmosCoin   = "cosmos_coin"
	AttributeKeyERC20Token   = "erc20_token" 
	AttributeKeyReceiver     = "receiver"
	AttributeKeyReturnAmount = "return_amount"

	ERC20EventTransfer = "Transfer"
)


type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}
