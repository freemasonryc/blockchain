package core

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ChainID = "sc_8888-1"

	CommandName = "smartchain"

	Version = "22.06.15"


	MainToken = "tt"


	GatewayToken = "fmc"

	ChainDefaultFeeStr string = "0.001"


	MinRealAmountFloat64 float64 = 0.0000000001

	CoinPlaces = 18


	RealToLedgerRate float64 = float64(RealToLedgerRateInt64)


	RealToLedgerRateInt64 int64 = 1000000000000000000


	LedgerToRealRate = "0.000000000000000001"


	CommitTime = 6

	MinimumGasPrices = "0.00005"


	GatewayBonusAddress = "gatewayBonus"
)

var (
	EvmRpcURL = "http://localhost:8545"

	ServerURL = "http://127.0.0.1:" + CosmosApiPort

	RpcURL = "tcp://127.0.0.1:" + RpcPort


	DefaultChainSeed = []string{}

	CosmosApiPort = "1317"

	RpcPort = "26657"

	P2pPort = "26656"
)

var (
	MinRealAmountDec = sdk.NewDecWithPrec(1, 10)

	RealToLedgerRateDec = sdk.MustNewDecFromStr("1000000000000000000")


	MortgageRatioDecNode = sdk.NewDec(5).Quo(sdk.NewDec(100))


	MortgageRatioDecBurn = sdk.NewDec(5).Quo(sdk.NewDec(100))


	MortgageRatioDecCommunity = sdk.NewDec(1).Quo(sdk.NewDec(100))


	MortgageRatioDecEcological = sdk.NewDec(1).Quo(sdk.NewDec(100))


	MortgageRatioDecPos = sdk.NewDec(3).Quo(sdk.NewDec(100))


	MortgageRatioDecRemain = sdk.NewDec(85).Quo(sdk.NewDec(100))
)
