/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// DescFile the required yml file name and path
var DescFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ddlgen",
	Short: "A database DDL generation tool",
	Long:  `A tool which takes JSON files describing tables and relationships and generates PostgreSQL compatible DDL.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing the %v tool:  '%s'", "ddlgen", err)
		os.Exit(1)
	}
}

func init() {
	// require flag
	rootCmd.PersistentFlags().StringVarP(&DescFile, "descriptor", "d", "", "The file path of the required YAML descriptor file used to direct tool generation.")
	rootCmd.MarkPersistentFlagRequired("descriptor")
}
