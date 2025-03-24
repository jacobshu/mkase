package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mkase",
		Short: "mkase creates .ase files",
		Long:  `A utility to creat .ase files that supports multiple input formats`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
