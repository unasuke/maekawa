package main

import (
	"strings"
)

func IsLambdaFunction(arn string) bool {
	return strings.HasPrefix(arn, "arn:aws:lambda")
}
