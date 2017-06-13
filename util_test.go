package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	cwe "github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func TestCompareString(t *testing.T) {
	string1 := "string"
	string2 := "sutoringu"
	string3 := ""

	if !CompareString(&string1, &string1) {
		t.Errorf("should true 'string' == 'string'")
	}

	if CompareString(&string1, &string2) {
		t.Errorf("should false 'string' == 'sutoringu'")
	}

	if CompareString(&string1, nil) {
		t.Errorf("should false 'string' == nil")
	}

	if !CompareString(nil, &string3) {
		t.Errorf("should true nil == ''")
	}

	if !CompareString(nil, nil) {
		t.Errorf("should true nil == nil")
	}
}

func TestCompareInt64(t *testing.T) {
	var integer1, integer2 int64 = 1, 2

	if !CompareInt64(&integer1, &integer1) {
		t.Errorf("should true 1 == 1")
	}

	if CompareInt64(&integer1, &integer2) {
		t.Errorf("should false 1 == 2")
	}

	if CompareInt64(&integer1, nil) {
		t.Errorf("should false 1 == nil")
	}

	if !CompareInt64(nil, nil) {
		t.Errorf("should true nil == nil")
	}
}

func TestCompareEcsParameters(t *testing.T) {
	ownEcs1 := EcsParameters{
		TaskDefinitionArn: "arn:aws:ecs:ap-northeast-1:000000000:task-definition/some-task:3",
		TaskCount:         2,
	}

	ownEcs2 := EcsParameters{
		TaskDefinitionArn: "arn:aws:ecs:ap-northeast-1:000000000:task-definition/some-task:4",
		TaskCount:         1,
	}

	ownEcsEmpty := EcsParameters{
		TaskDefinitionArn: "",
		TaskCount:         0,
	}

	theirsEcs1 := cwe.EcsParameters{
		TaskDefinitionArn: aws.String("arn:aws:ecs:ap-northeast-1:000000000:task-definition/some-task:3"),
		TaskCount:         aws.Int64(2),
	}

	if !CompareEcsParameters(&ownEcsEmpty, nil) {
		t.Errorf("should return true when compare ownEcsEmpty and nil")
	}

	if !CompareEcsParameters(&ownEcs1, &theirsEcs1) {
		t.Errorf("should return true when compare ownEcs1 and theirsEcs1")
	}

	if CompareEcsParameters(&ownEcs1, nil) {
		t.Errorf("should return false when compare ownEcs1 and nil")
	}

	if CompareEcsParameters(&ownEcs2, &theirsEcs1) {
		t.Errorf("should return false when compare ownEcs2 and theirsEcs1")
	}
}

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

func TestNilOrStringPtr(t *testing.T) {
	empty := ""
	if NilOrStringPtr(empty) != nil {
		t.Errorf("should return nil")
	}

	str := "some string"
	if r := NilOrStringPtr(str); *r != str {
		t.Errorf("shoul return string pointer")
	}
}

func TestNilSafeStr(t *testing.T) {
	str := "test"
	if NilSafeStr(&str) != "test" {
		t.Errorf("should return 'test'")
	}

	if NilSafeStr(nil) != "" {
		t.Errorf("should return empty string")
	}
}
