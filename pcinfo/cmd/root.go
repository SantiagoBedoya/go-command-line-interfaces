/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
	"runtime"

	"github.com/SantiagoBedoya/pcinfo/pkg"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "pcinfo",
	// Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		t := table.NewWriter()
		osName := runtime.GOOS
		t.SetTitle("Computer information - " + osName)
		t.AppendHeader(table.Row{"Component", "Property", "Value"})
		t.SetOutputMirror(os.Stdout)
		t.AppendRows(pkg.GetCpuInfo())
		t.AppendSeparator()
		t.AppendRows(pkg.GetMemoryInfo())
		t.AppendSeparator()
		t.AppendRows(pkg.GetHostInfo())
		t.AppendSeparator()
		t.AppendRows(pkg.GetMacInfo())
		t.Render()
	},
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pcinfo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
