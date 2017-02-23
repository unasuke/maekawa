package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func TestAssociateRules(t *testing.T) {
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
	describedRules := []Rule{
		Rule{
			Description:        "Test rule 1",
			EventPattern:       "",
			Name:               "test-1",
			ScheduleExpression: "cron(0 20 * * ? *)",
			State:              "ENABLED",
		},
		Rule{
			Description:        "Test rule 2",
			EventPattern:       "",
			Name:               "test-2",
			ScheduleExpression: "cron(0 20 2 * ? *)",
			State:              "ENABLED",
		},
	}
	AssociateRules(cweRules, describedRules)
	if *describedRules[0].ActualRule.Name != "test-1" {
		t.Errorf("should associate 'test-1'")
	}
	if *describedRules[1].ActualRule.Name != "test-2" {
		t.Errorf("should associate 'test-2'")
	}
}

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
	if len(target) != 1 {
		t.Errorf("should return deleted slice")
	}
	if *target[0].Name != "test-1" {
		t.Errorf("should deleted second rule(index is 1)")
	}
}

func TestJudgeRuleNeedUpdate(t *testing.T) {
	rule1 := Rule{
		Description:        "Test rule 1",
		EventPattern:       "",
		Name:               "test-1",
		ScheduleExpression: "cron(0 20 * * ? *)",
		State:              "ENABLED",
		ActualRule: cloudwatchevents.Rule{
			Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-1"),
			Description:        aws.String("Test rule 1"),
			EventPattern:       nil,
			Name:               aws.String("test-1"),
			RoleArn:            nil,
			ScheduleExpression: aws.String("cron(0 20 * * ? *)"),
			State:              aws.String("ENABLED"),
		},
	}
	JudgeRuleNeedUpdate(&rule1)
	if rule1.NeedUpdate != false {
		t.Errorf("rule1 shouldn't need update")
	}

	rule2 := Rule{
		Description:        "Test rule 2",
		EventPattern:       "",
		Name:               "test-2",
		ScheduleExpression: "cron(0 20 * * ? *)",
		State:              "ENABLED",
		ActualRule: cloudwatchevents.Rule{
			Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-2"),
			Description:        aws.String("Test rule 2"),
			EventPattern:       nil,
			Name:               aws.String("test-2"),
			RoleArn:            nil,
			ScheduleExpression: aws.String("rate(1 day)"),
			State:              aws.String("ENABLED"),
		},
	}
	JudgeRuleNeedUpdate(&rule2)
	if rule2.NeedUpdate != true {
		t.Errorf("rule2 should need update")
	}
}
