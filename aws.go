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

func getAWSConfig(profile, region string) (aws.Config, context.Context) {

	ctx := context.TODO()

	var cfg aws.Config
	var err error

	if profile == "" {
		cfg, err = config.LoadDefaultConfig(ctx, config.WithRegion(region))
	} else {
		cfg, err = config.LoadDefaultConfig(ctx,
			config.WithSharedConfigProfile(profile),
			config.WithRegion(region),
		)
	}

	if err != nil {
		panic(err)
	}

	return cfg, ctx
}

func assumeRole(cmd *cobra.Command, args []string) {
	roleToBeAssumed := getStringFlag(cmd, "role")
	sessionName := getStringFlag(cmd, "session-name")
	profile := getStringFlag(cmd, "profile")
	region := getStringFlag(cmd, "region")
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

	cfg, ctx := getAWSConfig(profile, region)

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
	cfg, ctx := getAWSConfig(getStringFlag(cmd, "profile"), getStringFlag(cmd, "region"))

	stsClient := sts.NewFromConfig(cfg)
	callerIdentityOutput, err := stsClient.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})

	if err != nil {
		panic(err)
	}

	printCallerIdentity(callerIdentityOutput.Account, callerIdentityOutput.Arn, callerIdentityOutput.UserId)
}
