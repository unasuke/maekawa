package main

import (
	"testing"
)

func TestIsLambdaFunction(t *testing.T) {
	lambdaArn := "arn:aws:lambda:ap-northeast-1:000000000000:function:lambda-function-name:3"
	if IsLambdaFunction(lambdaArn) == false {
		t.Errorf("%s is lambda function", lambdaArn)
	}

	ec2Arn := "arn:aws:ec2:us-east-1:123456789012:instance/*"
	if IsLambdaFunction(ec2Arn) == true {
		t.Errorf("%s is not lambda function", ec2Arn)
	}
}
