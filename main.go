/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"os"

	"github.com/k-makino-jp/cli-built-with-golang-following-clean-architecture/interface_adapter/controller"
)

func main() {
	const successfulExitCode, failureExitCode int = 0, 1
	controller := controller.GetController()
	if err := controller.Execute(); err != nil {
		os.Exit(failureExitCode)
	}
	os.Exit(successfulExitCode)
}
