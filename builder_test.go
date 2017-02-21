package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func TestDeleteRuleFromSlice(t *testing.T) {
	cweRules := []*cloudwatchevents.Rule{
		&cloudwatchevents.Rule{
			Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-1"),
			Description:        aws.String("Test rule 1"),
			EventPattern:       nil,
			Name:               aws.String("test-1"),
			RoleArn:            nil,
			ScheduleExpression: aws.String("cron(0 20 * * ? *)"),
			State:              aws.String("ENABLED"),
		},
		&cloudwatchevents.Rule{
			Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-2"),
			Description:        aws.String("Test rule 2"),
			EventPattern:       nil,
			Name:               aws.String("test-2"),
			RoleArn:            nil,
			ScheduleExpression: aws.String("cron(0 20 2 * ? *)"),
			State:              aws.String("ENABLED"),
		},
	}

	target := DeleteRuleFromSlice(cweRules, 1)
	if cap(target) != 1 {
		t.Errorf("should return deleted slice")
	}
	if *target[0].Name != "test-1" {
		t.Errorf("should deleted second rule(index is 1)")
	}
}
