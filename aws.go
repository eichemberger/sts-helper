package main

import (
	"context"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/spf13/cobra"
)

func assumeRole(cmd *cobra.Command, args []string) {
	roleToBeAssumed := getStringFlag(cmd, "role")
	sessionName := getStringFlag(cmd, "session-name")
	duration, err := cmd.Flags().GetInt32("duration")

	if err != nil {
		panic(err)
	}

	shouldCopyIntoClipboard, err := cmd.Flags().GetBool("copy")

	if err != nil {
		panic(err)
	}

	if roleToBeAssumed == "" {
		fmt.Println("[ERROR]: Role ARN is required")
		return
	}

	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	stsClient := sts.NewFromConfig(cfg)
	assumeRoleOutput, err := stsClient.AssumeRole(ctx, &sts.AssumeRoleInput{
		RoleArn:         aws.String(roleToBeAssumed),
		RoleSessionName: aws.String(sessionName),
		DurationSeconds: &duration,
	})

	if err != nil {
		panic(err)
	}

	printAssumeOutputMessage(roleToBeAssumed, sessionName, duration)
	printCredentials(
		assumeRoleOutput.Credentials.AccessKeyId,
		assumeRoleOutput.Credentials.SecretAccessKey,
		assumeRoleOutput.Credentials.SessionToken,
	)

	if shouldCopyIntoClipboard {
		err := clipboard.WriteAll(getCredentialsFormatted(
			assumeRoleOutput.Credentials.AccessKeyId,
			assumeRoleOutput.Credentials.SecretAccessKey,
			assumeRoleOutput.Credentials.SessionToken,
		))

		if err != nil {
			panic(err)
		}

		fmt.Println("Credentials copied to clipboard")
	}

}

func getCallerIdentity(cmd *cobra.Command, args []string) {
	ctx := context.TODO()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic(err)
	}

	stsClient := sts.NewFromConfig(cfg)
	callerIdentityOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})

	if err != nil {
		panic(err)
	}

	printCallerIdentity(callerIdentityOutput.Account, callerIdentityOutput.Arn, callerIdentityOutput.UserId)
}
