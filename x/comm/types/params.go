package types

import (
	"freemasonry.cc/blockchain/cmd/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyCommunityAddress  = []byte("CommunityAddress")
	KeyEcologicalAddress = []byte("EcologicalAddress")
	KeyCoin              = []byte("Coin")
)


func NewParams(
	communityAddress string,
	ecologicalAddress string,
	coin sdk.Coin,
) Params {
	return Params{
		CommunityAddress:  communityAddress,
		EcologicalAddress: ecologicalAddress,
		MinMortgageCoin:   coin,
	}
}

func DefaultParams() Params {

	
	return Params{
		CommunityAddress:  "",
		EcologicalAddress: "",
		MinMortgageCoin:   sdk.NewCoin(config.BaseDenom, sdk.NewInt(1)),
	}
}

func (p Params) Validate() error {
	if err := validateAddrString(p.CommunityAddress); err != nil {
		return err
	}

	if err := validateAddrString(p.EcologicalAddress); err != nil {
		return err
	}

	if err := validateCoin(p.MinMortgageCoin); err != nil {
		return err
	}
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCommunityAddress, &p.CommunityAddress, validateAddrString),
		paramtypes.NewParamSetPair(KeyEcologicalAddress, &p.EcologicalAddress, validateAddrString),
		paramtypes.NewParamSetPair(KeyCoin, &p.MinMortgageCoin, validateCoin),
	}
}

func validateAddrString(i interface{}) error {

	return nil
}

func validateCoin(i interface{}) error {

	return nil
}
