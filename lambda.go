package main

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func addPermissionToLambdaFromCloudWatchEvents(lc *lambda.Lambda, rules []Rule) error {
	for _, rule := range rules {
		for _, target := range rule.Targets {
			if !IsLambdaFunction(target.Arn) {
				continue
			}

			if result, err := isAlreadyAddPermission(lc, rule, target); err != nil {
				return err
			} else if result {
				// do nothing (already granted permission)
				continue
			} else {
				_, errL := lc.AddPermission(&lambda.AddPermissionInput{
					Action:       aws.String("lambda:InvokeFunction"),
					FunctionName: aws.String(LambdaFunctionNameFromArn(target.Arn)),
					Principal:    aws.String("events.amazonaws.com"),
					SourceArn:    rule.ActualRule.Arn,
					StatementId:  aws.String(target.ID),
				})

				if errL != nil {
					return errL
				}
			}
		}
	}
	return nil
}

func removePermissonFromLambda(lc *lambda.Lambda, rules []Rule) error {
	for _, rule := range rules {
		for _, target := range rule.Targets {
			if target.NeedDelete && IsLambdaFunction(*target.ActualTarget.Arn) {
				_, err := lc.RemovePermission(&lambda.RemovePermissionInput{
					FunctionName: target.ActualTarget.Arn,
					StatementId:  target.ActualTarget.Id,
				})

				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

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
			statement.StatementID == target.ID {
			return true, nil
		}
	}

	return false, nil
}

// IsLambdaFunction return true if passed arn is lambda
func IsLambdaFunction(arn string) bool {
	return strings.HasPrefix(arn, "arn:aws:lambda")
}

// LambdaFunctionNameFromArn return lambda function name from arn string
func LambdaFunctionNameFromArn(arn string) string {
	return strings.Split(arn, ":")[6]
}
