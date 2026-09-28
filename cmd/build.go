package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func runBuild(cmd *cobra.Command, arg []string) {
	if verbose {
		fmt.Println("A coder")
	} else {
		fmt.Println("A coder")
	}
}

var buildCmd = &cobra.Command{
	Use: "build",
	Short: "compile le projet",
	Run: runBuild,
}

func init(){
	rootCmd.AddCommand(buildCmd)
	buildCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Affiche en detaille de la compilation du projet du projet")
}