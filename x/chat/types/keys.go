package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)


const (
	
	ModuleName     = "chat"
	ModuleBurnName = "chat_burn"
	
	StoreKey = ModuleName

	
	RouterKey = ModuleName
)


var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}


const (
	
	MobileSuffixLength = 5 

	
	MobileSuffixMax = 99999

	KeyPrefixRegisterInfo = "chat_register_info_"

	KeyPrefixLastGetRewardLog = "chat_last_get_reward_log_"

	
	KeyPrefixMortgageAddLog = "chat_mortgage_add_log_"
)
