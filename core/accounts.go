package core

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcTransferTypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

var (
	
	ContractAddressFee = authtypes.NewModuleAddress(authtypes.FeeCollectorName)

	
	ContractAddressBank = authtypes.NewModuleAddress(bankTypes.ModuleName)

	
	ContractAddressDistribution = authtypes.NewModuleAddress(distrtypes.ModuleName)

	
	ContractAddressStakingBonded = authtypes.NewModuleAddress(stakingtypes.BondedPoolName)

	
	ContractAddressStakingNotBonded = authtypes.NewModuleAddress(stakingtypes.NotBondedPoolName)

	ContractAddressGov = authtypes.NewModuleAddress(govtypes.ModuleName)

	
	ContractAddressIbcTransfer = authtypes.NewModuleAddress(ibcTransferTypes.ModuleName)

	
	ContractGatewayBonus = authtypes.NewModuleAddress(GatewayBonusAddress)
)
