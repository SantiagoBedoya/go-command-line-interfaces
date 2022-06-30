/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/SantiagoBedoya/spotify-cli/internal/albums"
	"github.com/spf13/cobra"
)

// albumsCmd represents the albums command
var albumsCmd = &cobra.Command{
	Use:   "albums",
	Short: "Lookup albums",
	Run: func(cmd *cobra.Command, args []string) {
		service := albums.NewService()
		service.Get()
	},
}

func init() {
	rootCmd.AddCommand(albumsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// albumsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// albumsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
