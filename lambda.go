package main

import (
	"strings"
)

func IsLambdaFunction(arn string) bool {
	return strings.HasPrefix(arn, "arn:aws:lambda")
}

func LambdaFunctionNameFromArn(arn string) string {
	return strings.Split(arn, ":")[6]
}
