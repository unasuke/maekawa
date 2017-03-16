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
	describedRules1 := []Rule{
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
	result1 := AssociateRules(cweRules, describedRules1)
	for i, r := range result1 {
		if r.Name != *r.ActualRule.Name {
			t.Errorf("result1[%d] should associate 'test-1' but associated %s", i, *r.ActualRule.Name)
		}
	}

	describedRules2 := []Rule{
		Rule{
			Description:        "Test rule 2",
			EventPattern:       "",
			Name:               "test-2",
			ScheduleExpression: "cron(0 20 2 * ? *)",
			State:              "ENABLED",
		},
	}
	result2 := AssociateRules(cweRules, describedRules2)
	if l := len(result2); l != 2 {
		t.Errorf("result2 length should 2 but length is %d", l)
	}
	for i, r := range result2 {
		if r.Name == "test-2" && r.Name != *r.ActualRule.Name {
			t.Errorf("result2[%d] should associate 'test-2' but associated %s", i, *r.ActualRule.Name)
		}
		if r.Name == "" && *r.ActualRule.Name != "test-1" {
			t.Errorf("result2[%d] (empty) should associate 'test-1' but associated %s", i, *r.ActualRule.Name)
		}
	}

	cweRules2 := []*cloudwatchevents.Rule{
		&cloudwatchevents.Rule{
			Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-1"),
			Description:        aws.String("Test rule 1"),
			EventPattern:       nil,
			Name:               aws.String("test-1"),
			RoleArn:            nil,
			ScheduleExpression: aws.String("cron(0 20 * * ? *)"),
			State:              aws.String("ENABLED"),
		},
	}

	describedRules3 := []Rule{
		Rule{
			Description:        "Test rule 2",
			EventPattern:       "",
			Name:               "test-2",
			ScheduleExpression: "cron(0 20 2 * ? *)",
			State:              "ENABLED",
		},
	}
	result3 := AssociateRules(cweRules2, describedRules3)
	if l := len(result3); l != 2 {
		t.Errorf("result3 length should 2 but length is %d", l)
	}
	for i, r := range result3 {
		if r.Name == "test-2" && r.ActualRule.Name != nil {
			t.Errorf("result3[%d] should associate empty but associated %s", i, *r.ActualRule.Name)
		}
		if r.Name == "" && *r.ActualRule.Name != "test-1" {
			t.Errorf("result3[%d] (empty) should associate 'test-1' but associated %s", i, *r.ActualRule.Name)
		}
	}
}

func TestAssociateTargets(t *testing.T) {
	cweTargets := []*cloudwatchevents.Target{
		&cloudwatchevents.Target{
			Arn: aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-1"),
			Id:  aws.String("Id1"),
		},
		&cloudwatchevents.Target{
			Arn: aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-2"),
			Id:  aws.String("Id2"),
		},
	}

	testTargets1 := []Target{
		Target{
			Arn: "arn:aws:lambda:ap-northeast-1:000000000000:function:test-2",
			Id:  "Id2",
		},
		Target{
			Arn: "arn:aws:lambda:ap-northeast-1:000000000000:function:test-1",
			Id:  "Id1",
		},
	}

	result1 := AssociateTargets(cweTargets, testTargets1)
	if result1[0].Arn != *result1[0].ActualTarget.Arn &&
		*result1[0].ActualTarget.Arn != *cweTargets[1].Arn {
		t.Errorf("result1[0] should associate cweTargets[1]")
	}
	if result1[1].Arn != *result1[1].ActualTarget.Arn &&
		*result1[1].ActualTarget.Arn != *cweTargets[0].Arn {
		t.Errorf("result1[1] should associate cweTargets[0]")
	}

	testTargets2 := []Target{
		Target{
			Arn: "arn:aws:lambda:ap-northeast-1:000000000000:function:test-2",
			Id:  "Id2",
		},
	}

	result2 := AssociateTargets(cweTargets, testTargets2)
	if l := len(result2); l != 2 {
		t.Errorf("testTargets2 length should be 2 but length is %d", l)
	}
	for i, r := range result2 {
		if r.Id == "Id2" && r.Arn != *r.ActualTarget.Arn {
			t.Errorf("result2[%d] should associate 'Id2' but associated %s", i, *r.ActualTarget.Id)
		}
		if r.Id == "" && *r.ActualTarget.Id != "Id1" {
			t.Errorf("result2[%d] (empty) should associate 'Id1' but associated %s", i, *r.ActualTarget.Id)
		}
	}

	testTargets3 := []Target{
		Target{
			Arn: "arn:aws:lambda:ap-northeast-1:000000000000:function:test-2",
			Id:  "Id2",
		},
		Target{
			Arn: "arn:aws:lambda:ap-northeast-1:000000000000:function:test-3",
			Id:  "Id3",
		},
	}
	result3 := AssociateTargets(cweTargets, testTargets3)
	if l := len(result3); l != 3 {
		t.Errorf("testTargets3 length should be 3 but length is %d", l)
	}
	for i, r := range result3 {
		if r.Id == "" && *r.ActualTarget.Id != "Id1" {
			t.Errorf("ActualTarget 'Id1' should associate empty target but associated %s", r.Id)
		}
		if r.Id == "Id2" && r.Arn != *r.ActualTarget.Arn {
			t.Errorf("result3[%d] should associate 'Id2' but associated %s", i, *r.ActualTarget.Id)
		}
		if r.Id == "Id3" && r.ActualTarget.Id != nil {
			t.Errorf("result3[%d] (empty) should associate empty ActualTarget but associated %s", i, *r.ActualTarget.Id)
		}
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

func TestJudgeRuleNeedDelete(t *testing.T) {
	rule1 := Rule{
		Description:        "Test rule 1",
		EventPattern:       "",
		Name:               "test-1",
		RoleArn:            "",
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
	JudgeRuleNeedDelete(&rule1)
	if rule1.NeedDelete == true {
		t.Errorf("rule1 shouldn't need delete")
	}

	rule2 := Rule{
		Description:        "",
		EventPattern:       "",
		Name:               "",
		ScheduleExpression: "",
		RoleArn:            "",
		State:              "",
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
	JudgeRuleNeedDelete(&rule2)
	if rule2.NeedDelete == false {
		t.Errorf("rule2 should need delete")
	}
}

func TestJudgeTargetNeedUpdate(t *testing.T) {
	target1 := Target{
		Arn:   "arn:aws:lambda:ap-northeast-1:000000000000:function:test-1",
		Id:    "Id1",
		Input: "input",
		ActualTarget: cloudwatchevents.Target{
			Arn:       aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-1"),
			Id:        aws.String("Id1"),
			Input:     aws.String("input"),
			InputPath: nil,
		},
	}
	JudgeTargetNeedUpdate(&target1)
	if target1.NeedUpdate == true {
		t.Errorf("target1 shouldn't need update")
	}

	target2 := Target{
		Arn:   "arn:aws:lambda:ap-northeast-1:000000000000:function:test-2",
		Id:    "Id2",
		Input: "input",
		ActualTarget: cloudwatchevents.Target{
			Arn:       aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-2"),
			Id:        aws.String("Id2"),
			Input:     aws.String("another input"),
			InputPath: nil,
		},
	}
	JudgeTargetNeedUpdate(&target2)
	if target2.NeedUpdate == false {
		t.Errorf("target2 should need update")
	}
}

func TestJudgeTargetNeedDelete(t *testing.T) {
	target1 := Target{
		Arn:   "arn:aws:lambda:ap-northeast-1:000000000000:function:test-1",
		Id:    "Id1",
		Input: "input",
		ActualTarget: cloudwatchevents.Target{
			Arn:       aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-1"),
			Id:        aws.String("Id1"),
			Input:     aws.String("input"),
			InputPath: nil,
		},
	}
	JudgeTargetNeedDelete(&target1)
	if target1.NeedDelete == true {
		t.Errorf("target1 shouldn't need delete")
	}

	target2 := Target{
		Arn:   "",
		Id:    "",
		Input: "",
		ActualTarget: cloudwatchevents.Target{
			Arn:       aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-2"),
			Id:        aws.String("Id2"),
			Input:     aws.String("another input"),
			InputPath: nil,
		},
	}
	JudgeTargetNeedDelete(&target2)
	if target2.NeedDelete == false {
		t.Errorf("target2 should need delete")
	}
}

func TestCheckIsNeedUpdateOrDelete(t *testing.T) {
	rules := []Rule{
		Rule{
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
			Targets: []Target{
				Target{
					Arn:   "arn:aws:lambda:ap-northeast-1:000000000000:function:test-1",
					Id:    "Id1",
					Input: "input",
					ActualTarget: cloudwatchevents.Target{
						Arn:       aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-1"),
						Id:        aws.String("Id1"),
						Input:     aws.String("input"),
						InputPath: nil,
					},
				},
			},
		},
		Rule{
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
			Targets: []Target{
				Target{
					Arn:   "arn:aws:lambda:ap-northeast-1:000000000000:function:test-2",
					Id:    "Id2",
					Input: "input",
					ActualTarget: cloudwatchevents.Target{
						Arn:       aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-2"),
						Id:        aws.String("Id2"),
						Input:     aws.String("another input"),
						InputPath: nil,
					},
				},
			},
		},
		Rule{
			Description:        "",
			EventPattern:       "",
			Name:               "",
			ScheduleExpression: "",
			State:              "",
			ActualRule: cloudwatchevents.Rule{
				Arn:                aws.String("arn:aws:events:ap-northeast-1:000000000000:rule/test-3"),
				Description:        aws.String("Test rule 3"),
				EventPattern:       nil,
				Name:               aws.String("test-3"),
				RoleArn:            nil,
				ScheduleExpression: aws.String("rate(2 day)"),
				State:              aws.String("ENABLED"),
			},
			Targets: []Target{
				Target{
					Arn:   "",
					Id:    "",
					Input: "",
					ActualTarget: cloudwatchevents.Target{
						Arn:       aws.String("arn:aws:lambda:ap-northeast-1:000000000000:function:test-3"),
						Id:        aws.String("Id3"),
						Input:     aws.String("will removed input"),
						InputPath: nil,
					},
				},
			},
		},
	}
	CheckIsNeedUpdateOrDelete(rules)
	if rules[0].NeedUpdate == true {
		t.Errorf("rule[0] shouldn't need update")
	}
	if rules[0].Targets[0].NeedUpdate == true {
		t.Errorf("rule[0].Targets[0] shouldn't need update")
	}
	if rules[1].NeedUpdate == false {
		t.Errorf("rule[1] should need update")
	}
	if rules[1].Targets[0].NeedUpdate == false {
		t.Errorf("rule[1].Targets[0] should need update")
	}
	if rules[2].NeedDelete == false {
		t.Errorf("rule[2] should need delete")
	}
	if rules[2].Targets[0].NeedDelete == false {
		t.Errorf("rule[2].Targets[0] should need delete")
	}
}
