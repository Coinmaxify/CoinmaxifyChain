package version

import (
	"github.com/Coinmaxify/Coinmaxify/command"
	"github.com/Coinmaxify/Coinmaxify/versioning"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Returns the current Coinmaxify version",
		Args:  cobra.NoArgs,
		Run:   runCommand,
	}
}

func runCommand(cmd *cobra.Command, _ []string) {
	outputter := command.InitializeOutputter(cmd)
	defer outputter.WriteOutput()

	outputter.SetCommandResult(
		&VersionResult{
			Version: versioning.Version,
		},
	)
}
