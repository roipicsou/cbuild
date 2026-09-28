package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:   "moncli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bienvenue sur moncli ! Utilise --help ou -h pour voir les options.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur : %v\n", err)
		os.Exit(1)
	}
}