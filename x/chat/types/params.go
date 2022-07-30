package types

import (
	"fmt"
	"freemasonry.cc/blockchain/cmd/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyCommunityAddress  = []byte("CommunityAddress")
	KeyEcologicalAddress = []byte("EcologicalAddress")
	KeyCoin              = []byte("Coin")
	KeyChatRewardLog     = []byte("ChatRewardLog")
	KeyMaxPhoneNumber    = []byte("MaxPhoneNumber")
)


func NewParams(
	communityAddress string,
	ecologicalAddress string,
	coin sdk.Coin,
	chatRewardLog []ChatReward,
	maxPhoneNumber uint64,
) Params {
	return Params{
		CommunityAddress:  communityAddress,
		EcologicalAddress: ecologicalAddress,
		MinMortgageCoin:   coin,
		ChatRewardLog:     chatRewardLog,
		MaxPhoneNumber:    maxPhoneNumber,
	}
}

func DefaultParams() Params {


	return Params{
		CommunityAddress:  "dex1vmx0e3r7v93axstfkqpxjpvgearcmxls2x287f",
		EcologicalAddress: "dex1wrx8tdyv5j3l5lst9n080aar3v0f5zywh303cy",
		MinMortgageCoin:   sdk.NewCoin(config.BaseDenom, sdk.NewInt(1000000000000000000)),
		ChatRewardLog: []ChatReward{
			{
				Height: 1,
				Value:  "0.01",
			},
		},
		MaxPhoneNumber: 10,
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

	if err := validateChatRewardLog(p.ChatRewardLog); err != nil {
		return err
	}

	if err := validateMaxPhoneNumber(p.MaxPhoneNumber); err != nil {

	}
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyCommunityAddress, &p.CommunityAddress, validateAddrString),
		paramtypes.NewParamSetPair(KeyEcologicalAddress, &p.EcologicalAddress, validateAddrString),
		paramtypes.NewParamSetPair(KeyCoin, &p.MinMortgageCoin, validateCoin),
		paramtypes.NewParamSetPair(KeyChatRewardLog, &p.ChatRewardLog, validateChatRewardLog),
		paramtypes.NewParamSetPair(KeyMaxPhoneNumber, &p.MaxPhoneNumber, validateMaxPhoneNumber),
	}
}

func validateAddrString(i interface{}) error {

	return nil
}

func validateChatRewardLog(i interface{}) error {
	v, ok := i.([]ChatReward)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) == 0 {
		return fmt.Errorf("invalid parameter len")
	}

	if len(v) > 1 {
		for i := 0; i < len(v)-1; i++ {
			if v[i].Height > v[i+1].Height {
				return fmt.Errorf("invalid parameter height")
			}

			curValue := v[i].Value
			curValueDec, err := sdk.NewDecFromStr(curValue)
			if err != nil {
				return fmt.Errorf("invalid value format")
			}
			if curValueDec.GT(sdk.NewDec(1)) {
				return fmt.Errorf("invalid value (GT1)")
			}
			if curValueDec.LT(sdk.MustNewDecFromStr("0.0001")) {
				return fmt.Errorf("invalid value (LT0.0001)")
			}
		}
	}

	return nil
}

func validateMaxPhoneNumber(i interface{}) error {

	return nil
}

func validateCoin(i interface{}) error {

	return nil
}

func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable(
		paramtypes.NewParamSetPair(KeyCommunityAddress, DefaultParams().CommunityAddress, validateAddrString),
		paramtypes.NewParamSetPair(KeyEcologicalAddress, DefaultParams().EcologicalAddress, validateAddrString),
		paramtypes.NewParamSetPair(KeyCoin, DefaultParams().MinMortgageCoin, validateCoin),
		paramtypes.NewParamSetPair(KeyChatRewardLog, DefaultParams().ChatRewardLog, validateChatRewardLog),
		paramtypes.NewParamSetPair(KeyMaxPhoneNumber, DefaultParams().MaxPhoneNumber, validateMaxPhoneNumber),
	)
}
