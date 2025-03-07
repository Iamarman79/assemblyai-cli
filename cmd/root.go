/*
Copyright © 2022 AssemblyAI support@assemblyai.com
*/
package cmd

import (
	"fmt"
	"os"

	S "github.com/AssemblyAI/assemblyai-cli/schemas"
	U "github.com/AssemblyAI/assemblyai-cli/utils"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var VERSION string

var rootCmd = &cobra.Command{
	Use:   "assemblyai",
	Short: "AssemblyAI CLI",
	Long: `Please authenticate to use the CLI.
assemblyai config [token]`,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Run: func(cmd *cobra.Command, args []string) {
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Printf("AssemblyAI CLI %s\n", VERSION)
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	if VERSION == "" {
		godotenv.Load()
		VERSION = os.Getenv("VERSION")
	}
	U.CheckForUpdates(VERSION)
	if err := rootCmd.Execute(); err != nil {
		printErrorProps := S.PrintErrorProps{
			Error:   err,
			Message: err.Error(),
		}
		U.PrintError(printErrorProps)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Check current installed version.")
	rootCmd.Flags().Bool("test", false, "Flag for test executing purpose")
	rootCmd.Flags().MarkHidden("test")
}
