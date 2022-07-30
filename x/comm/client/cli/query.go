package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"freemasonry.cc/blockchain/x/comm/types"
)


func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the erc20 module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(



	)
	return cmd
}

//












//

//




//



//




//



//



//












//

//



//




//



//



//












//

//

//




//



//



