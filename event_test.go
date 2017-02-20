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
		EventPattern:       nil,
		Name:               aws.String("test-cwevent"),
		RoleArn:            nil,
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
	if result1 != 0.8 {
		t.Error("match score should be eq 0.8 but got", result1)
	}

	describedRule.Description = "another test rule"
	result2 := MatchScoreForCWEventRuleAndDescribedRule(cweventRule, describedRule)
	if result2 != 0.6 {
		t.Error("match score should be eq 0.6 but got", result2)
	}
}

func TestIsNewDefinedRule(t *testing.T) {
	cweventRule := cloudwatchevents.Rule{
		Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-cwevent"),
		Description:        aws.String("Test rule"),
		EventPattern:       nil,
		Name:               aws.String("test-cwevent"),
		RoleArn:            nil,
		ScheduleExpression: aws.String("cron(0 20 * * ? *)"),
		State:              aws.String("ENABLED"),
	}

	rule1 := Rule{
		Description:        "Test rule",
		EventPattern:       "",
		Name:               "test-cwevent",
		ScheduleExpression: "cron(0 20 * * ? *)",
		State:              "ENABLED",
		LambdaFunctions:    nil,
	}

	result1 := IsNewDefinedRule(cweventRule, rule1)
	if result1 != false {
		t.Error("return value should be false")
	}

	rule2 := Rule{
		Description:        "New defined test rule",
		EventPattern:       "",
		Name:               "another-test-cwevent",
		ScheduleExpression: "cron(0 20 * * ? *)",
		State:              "ENABLED",
		LambdaFunctions:    nil,
	}

	result2 := IsNewDefinedRule(cweventRule, rule2)
	if result2 != true {
		t.Error("return value should be true")
	}
}
