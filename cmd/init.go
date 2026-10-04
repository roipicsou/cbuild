package cmd

import (
	"os"
	"path/filepath"

	persfile "github.com/roipicsou/cbuild/persFile"
	"github.com/spf13/cobra"
)

/*
	Detection du dossier
	Creation du fichier de configuration
	Creation du fichier c
*/

var nameProject string

func runInit(cmd *cobra.Command, arg []string) {
	if nameProject == "" {
		chemin, err := os.Getwd()

		if err != nil {
			return
		}
		nameProject = filepath.Base(chemin)
	}

	config := persfile.NewConfig("gcc", nameProject)

	err := persfile.WrtieFile("test.yaml", *config)
	if err != nil {
		return
	}
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialise le projet",
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Affiche en detaille l'initialisation du projet")
	initCmd.Flags().StringVarP(&nameProject, "name", "n", "", "Choix du nom a la place de cuilui du nom de dossier")
}
