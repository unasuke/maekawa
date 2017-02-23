package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	cwe "github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func TestDeleteRuleFromSlice(t *testing.T) {
	cweRules := []*cwe.Rule{
		&cwe.Rule{
			Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-1"),
			Description:        aws.String("Test rule 1"),
			EventPattern:       nil,
			Name:               aws.String("test-1"),
			RoleArn:            nil,
			ScheduleExpression: aws.String("cron(0 20 * * ? *)"),
			State:              aws.String("ENABLED"),
		},
		&cwe.Rule{
			Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-2"),
			Description:        aws.String("Test rule 2"),
			EventPattern:       nil,
			Name:               aws.String("test-2"),
			RoleArn:            nil,
			ScheduleExpression: aws.String("cron(0 20 2 * ? *)"),
			State:              aws.String("ENABLED"),
		},
	}

	result := DeleteRuleFromSlice(cweRules, 1)
	if len(result) != 1 {
		t.Errorf("should return deleted slice")
	}
	if *result[0].Name != "test-1" {
		t.Errorf("should deleted second rule(index is 1)")
	}
}

func TestDeleteTargetFromSlice(t *testing.T) {
	cweTargets := []*cwe.Target{
		&cwe.Target{
			Arn: aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-1"),
			Id:  aws.String("Id1"),
		},
		&cwe.Target{
			Arn: aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-2"),
			Id:  aws.String("Id2"),
		},
	}

	result := DeleteTargetFromSlice(cweTargets, 1)
	if len(result) != 1 {
		t.Errorf("should return deleted slice")
	}
	if *result[0].Id != "Id1" {
		t.Errorf("should deleted second target(index is 1)")
	}
}
