package root

import (
	"fmt"
	"github.com/Coinmaxify/Coinmaxify/command/backup"
	"github.com/Coinmaxify/Coinmaxify/command/genesis"
	"github.com/Coinmaxify/Coinmaxify/command/helper"
	"github.com/Coinmaxify/Coinmaxify/command/ibft"
	"github.com/Coinmaxify/Coinmaxify/command/license"
	"github.com/Coinmaxify/Coinmaxify/command/loadbot"
	"github.com/Coinmaxify/Coinmaxify/command/monitor"
	"github.com/Coinmaxify/Coinmaxify/command/peers"
	"github.com/Coinmaxify/Coinmaxify/command/secrets"
	"github.com/Coinmaxify/Coinmaxify/command/server"
	"github.com/Coinmaxify/Coinmaxify/command/status"
	"github.com/Coinmaxify/Coinmaxify/command/txpool"
	"github.com/Coinmaxify/Coinmaxify/command/version"
	"github.com/spf13/cobra"
	"os"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Copyright 2022 | The Coinmaxify Developers",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		loadbot.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
