package cmd

import (
	"example.com/goddlgen/pkg/descriptor"
	"example.com/goddlgen/pkg/logger"
	"example.com/goddlgen/pkg/model"
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
		jsonData, err := model.ReadAll(desc.Input.JsonInput.FolderName)
		if err != nil {
			return
		}

		log := logger.Get()
		log.Info().Msgf("Successfully parsed %d JSON files", len(jsonData))

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
