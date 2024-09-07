package main

import (
	"fmt"
	"strings"
)

func getCredentialsFormatted(accessKey, secretAccessKey, sessionToken *string) string {
	return fmt.Sprintf(`export AWS_ACCESS_KEY_ID="` + *accessKey + `"
export AWS_SECRET_ACCESS_KEY="` + *secretAccessKey + `"
export AWS_SESSION_TOKEN="` + *sessionToken + `"
`)
}

func printCredentials(accessKey, secretAccessKey, sessionToken *string) {
	fmt.Println("\nCredentials:")
	fmt.Println(getCredentialsFormatted(accessKey, secretAccessKey, sessionToken))
}

func printAssumeOutputMessage(roleToBeAssumed, sessionName string, duration int32) {
	dashesString := strings.Repeat("-", len(fmt.Sprintf("Role ARN:    	 %-s", roleToBeAssumed))+10)

	strMessage := fmt.Sprintf(`Role ARN:    	 %-s
Session Name:	 %-s
Duration:    	 %-d seconds`, roleToBeAssumed, sessionName, duration)

	output := fmt.Sprintf(dashesString + "\n" + strMessage + "\n" + dashesString)

	fmt.Println(output)
}

func printCallerIdentity(account, arn, userId *string) {
	output := fmt.Sprintf(`Account: 	%s
Arn: 		%s
UserId: 	%s`, *account, *arn, *userId)

	fmt.Println(output)
}
