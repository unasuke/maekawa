package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func TestMatchScoreForCWEventRuleAndDescribedRule(t *testing.T) {
	cweventRule := cloudwatchevents.Rule{
		Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-cwevent"),
		Description:        aws.String("Test rule"),
		EventPattern:       aws.String(""),
		Name:               aws.String("test-cwevent"),
		RoleArn:            aws.String(""),
		ScheduleExpression: aws.String("cron(0 20 * * ? *)"),
		State:              aws.String("ENABLED"),
	}

	describedRule := Rule{
		Description:        "Test rule",
		EventPattern:       "",
		Name:               "test-cwevent",
		ScheduleExpression: "cron(0 20 * * ? *)",
		State:              "ENABLED",
		LambdaFunctions:    nil,
	}

	result1 := MatchScoreForCWEventRuleAndDescribedRule(cweventRule, describedRule)
	if result1 != 1.0 {
		t.Error("match score should be eq 1.0 but got", result1)
	}

	describedRule.Description = "another test rule"
	result2 := MatchScoreForCWEventRuleAndDescribedRule(cweventRule, describedRule)
	if result2 != 0.8 {
		t.Error("match score should be eq 0.8 but got", result2)
	}
}
