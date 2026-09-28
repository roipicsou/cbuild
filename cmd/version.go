package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func runVersion(cmd *cobra.Command, arg []string){
	if verbose {
		fmt.Println("MonCLI version 0.1")
	} else {
		fmt.Println("v.0.1")
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Affiche la version de l'outil",
	Run: runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Affiche les détails de version")
}