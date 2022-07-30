package cmd

import (
	"freemasonry.cc/blockchain/app"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"os"
)

func ChainRun(arg ...string) error {
	os.Args = arg
	rootCmd, _ := NewRootCmd()


	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {


		return err
	}
	return nil
}
