package cmd

import (
	"example.com/goddlgen/pkg/descriptor"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates PostgreSQL DDL for data specified in the descriptor file.",

	Run: func(cmd *cobra.Command, args []string) {
		// read the descriptor file
		var desc = descriptor.Descriptor{}
		err := desc.ReadFromYml(DescFile)
		if err != nil {
			panic(err)
		}

		// TODO Add Implementation

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
