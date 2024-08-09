/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"ddaily/cmd"

	gap "github.com/muesli/go-app-paths"
)

func main() {
	cmd.Execute()
}

func setupPath() string {
	scope := gap.NewScope
}