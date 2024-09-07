package main

import "github.com/spf13/cobra"

func getStringFlag(cmd *cobra.Command, flagName string) string {
	flag, err := cmd.Flags().GetString(flagName)
	if err != nil {
		panic(err)
	}
	return flag
}
