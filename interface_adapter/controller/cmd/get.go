/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/k-makino-jp/cli-built-with-golang-following-clean-architecture/interface_adapter/repository"
	"github.com/k-makino-jp/cli-built-with-golang-following-clean-architecture/usecase"
	"github.com/spf13/cobra"
)

var (
	useCaseGetCmd usecase.Cmd

	// getCmd represents the get command
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "get executes HTTP GET method.",
		Long:  "get executes HTTP GET method.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return useCaseGetCmd.Execute()
		},
	}
)

func init() {
	rootCmdController.AddCommand(getCmd)
	useCaseGetCmd = usecase.GetCmdInteractor{
		UserCollector: repository.NewUserRepository(),
	}
}
