package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var cmdAssumeRole = &cobra.Command{
		Use:   "assume",
		Short: "Assume role",
		Long:  `Print the credentials of the assumed role. By default, the credentials are copied to the clipboard. To avoid this, use the --copy=false flag.`,
		Run:   assumeRole,
	}

	var cmdGetCallerId = &cobra.Command{
		Use:   "whoami",
		Short: "Get caller identity",
		Long:  `Get the caller identity of the current user.`,
		Run:   getCallerIdentity,
	}

	var rootCmd = &cobra.Command{Use: "sts"}
	cmdAssumeRole.PersistentFlags().String("role", "", "Role ARN to assume")
	cmdAssumeRole.PersistentFlags().String("session-name", "assumed-role", "Role session name")
	cmdAssumeRole.PersistentFlags().Int32("duration", 3600, "Duration in seconds")
	cmdAssumeRole.PersistentFlags().Bool("copy", true, "Copy the credentials to the clipboard")
	rootCmd.AddCommand(cmdAssumeRole)
	rootCmd.AddCommand(cmdGetCallerId)
	rootCmd.Execute()
}
