package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func runClean(cmd *cobra.Command, arg []string) {
	if verbose {
		fmt.Println("A coder")
	} else {
		fmt.Println("A coder")
	}
}

var cleanCmd = &cobra.Command{
	Use: "clean",
	Short: "netoi le projet",
	Run: runClean,
}

func init(){
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Affiche en detaille du netoiague du projet du projet")
}