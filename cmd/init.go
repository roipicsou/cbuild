package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/*
	Detection du dossier
	Creation du fichier de configuration
	Creation du fichier c
*/

func runInit(cmd *cobra.Command, arg []string){
	if verbose {
		fmt.Println("A coder")
	} else {
		fmt.Println("A coder")
	}
}

var nameProject string

var initCmd = &cobra.Command{
	Use: "init",
	Short: "initialise le projet",
	Run: runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Affiche en detaille l'initialisation du projet")
	initCmd.Flags().StringVarP(&nameProject, "name", "n", "Non", "Choix du nom a la place de cuilui du nom de dossier")
}