package main

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/service/lambda"
)

func isAlreadyAddPermission(lc *lambda.Lambda, rule Rule, target Target) (bool, error) {
	var policy LambdaPolicy

	policyOutput, err := lc.GetPolicy(&lambda.GetPolicyInput{
		FunctionName: &target.Arn,
	})

	if err != nil {
		return false, err
	}

	errJ := json.Unmarshal([]byte(*policyOutput.Policy), &policy)
	if errJ != nil {
		return false, errJ
	}

	for _, statement := range *policy.Statement {
		if (statement.Resource == target.Arn &&
			strings.HasSuffix(statement.Condition.ArnLike.AwsSourceArn, rule.Name) &&
			statement.Effect == "Allow" &&
			statement.Principal.Service == "events.amazonaws.com" &&
			statement.Action == "lambda:InvokeFunction") ||
			statement.StatementId == target.Id {
			return true, nil
		}
	}

	return false, nil
}

func IsLambdaFunction(arn string) bool {
	return strings.HasPrefix(arn, "arn:aws:lambda")
}

func LambdaFunctionNameFromArn(arn string) string {
	return strings.Split(arn, ":")[6]
}
