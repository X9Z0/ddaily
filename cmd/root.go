/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ddaily",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ddaily.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func getThePath() string {
    var dataDir string

    switch osType := runtime.GOOS; osType {
    case "linux":
        dataDir = filepath.Join(os.Getenv("HOME"), ".local", "share", "ddaily")
    case "darwin":
        dataDir = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "ddaily")
    default:
        dataDir = filepath.Join(os.Getenv("HOME"), ".ddaily")
    }

    err := os.MkdirAll(dataDir, os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }

    return dataDir
}