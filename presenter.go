package main

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func displayWhatWillChange(rules []Rule) {
	updates := WillUpdateRulesAndTargets(rules)
	deletes := WillDeleteRulesAndTargets(rules)
	if len(updates) == 0 && len(deletes) == 0 {
		fmt.Println("No Changes")
	}
	if len(updates) > 0 {
		fmt.Println("Updates")
		for _, r := range updates {
			ShowWillUpdateFieldInRule(r)
			for _, t := range r.Targets {
				if t.NeedUpdate && !t.NeedDelete {
					ShowWillUpdateFieldInTarget(t)
				}
			}
		}
	}
	if len(deletes) > 0 {
		fmt.Println("Deletes")
		for _, r := range deletes {
			ShowWillDeleteRule(r)
			for _, t := range r.Targets {
				if t.NeedDelete {
					ShowWillDeleteTarget(t)
				}
			}
		}
	}
}

// WillUpdateRulesAndTargets return will be updated rules and targets
func WillUpdateRulesAndTargets(rules []Rule) []Rule {
	u := make([]Rule, 0)
	for _, rule := range rules {
		if rule.NeedUpdate && !rule.NeedDelete {
			u = append(u, rule)
		} else {
			for _, target := range rule.Targets {
				if target.NeedUpdate && !target.NeedDelete {
					u = append(u, rule)
					break
				}
			}
		}
	}
	return u
}

// WillDeleteRulesAndTargets return will be deleted rules and targets
func WillDeleteRulesAndTargets(rules []Rule) []Rule {
	d := make([]Rule, 0)
	for _, rule := range rules {
		if rule.NeedDelete {
			d = append(d, rule)
		} else {
			for _, target := range rule.Targets {
				if target.NeedDelete {
					d = append(d, rule)
					break
				}
			}
		}
	}
	return d
}

// ShowWillUpdateFieldInRule print what will rule changes to stdout
func ShowWillUpdateFieldInRule(rule Rule) {
	fmt.Printf("Rule: %s\n", rule.Name)
	if !CompareString(&rule.Name, rule.ActualRule.Name) {
		fmt.Printf("  Name: %s  ->  %s\n", NilSafeStr(rule.ActualRule.Name), rule.Name)
	}
	if !CompareString(&rule.Description, rule.ActualRule.Description) {
		fmt.Printf("  Description: %s  ->  %s\n", NilSafeStr(rule.ActualRule.Description), rule.Description)
	}
	if !CompareString(&rule.EventPattern, rule.ActualRule.EventPattern) {
		fmt.Printf("  EventPattern: %s  ->  %s\n", NilSafeStr(rule.ActualRule.EventPattern), rule.EventPattern)
	}
	if !CompareString(&rule.RoleArn, rule.ActualRule.RoleArn) {
		fmt.Printf("  RoleArn: %s  ->  %s\n", NilSafeStr(rule.ActualRule.RoleArn), rule.RoleArn)
	}
	if !CompareString(&rule.ScheduleExpression, rule.ActualRule.ScheduleExpression) {
		fmt.Printf("  ScheduleExpression: %s  ->  %s\n", NilSafeStr(rule.ActualRule.ScheduleExpression), rule.ScheduleExpression)
	}
	if !CompareString(&rule.State, rule.ActualRule.State) {
		fmt.Printf("  State: %s  ->  %s\n", NilSafeStr(rule.ActualRule.State), rule.State)
	}
}

// ShowWillUpdateFieldInTarget print what will target changes to stdout
func ShowWillUpdateFieldInTarget(target Target) {
	fmt.Printf("  Target: %s\n", target.Arn)
	if !CompareString(&target.Arn, target.ActualTarget.Arn) {
		fmt.Printf("    Arn: %s  ->  %s\n", NilSafeStr(target.ActualTarget.Arn), target.Arn)
	}
	if !CompareString(&target.ID, target.ActualTarget.Id) {
		fmt.Printf("    ID: %s  ->  %s\n", NilSafeStr(target.ActualTarget.Id), target.ID)
	}
	if !CompareString(&target.Input, target.ActualTarget.Input) {
		fmt.Printf("    Input: %s  ->  %s\n", NilSafeStr(target.ActualTarget.Input), target.Input)
	}
	if !CompareString(&target.InputPath, target.ActualTarget.InputPath) {
		fmt.Printf("    InputPath: %s  ->  %s\n", NilSafeStr(target.ActualTarget.InputPath), target.InputPath)
	}
	if !CompareString(&target.RoleArn, target.ActualTarget.RoleArn) {
		fmt.Printf("    RoleArn: %s  ->  %s\n", NilSafeStr(target.ActualTarget.RoleArn), target.RoleArn)
	}
	if !CompareEcsParameters(&target.EcsParameters, target.ActualTarget.EcsParameters) {
		fmt.Println("    EcsParameters:")
		ShowDiffOfTheEcsParameters(target.ActualTarget.EcsParameters, target.EcsParameters)
	}
	if !CompareKinesisParameters(&target.KinesisParameters, target.ActualTarget.KinesisParameters) {
		fmt.Println("    KinesisParameters:")
		ShowDiffOfTheKinesisParameters(target.ActualTarget.KinesisParameters, target.KinesisParameters)
	}
}

// ShowDiffOfTheEcsParameters print what will EcsParameters in target changes to stdout
func ShowDiffOfTheEcsParameters(current *cloudwatchevents.EcsParameters, expect EcsParameters) {
	var currTaskDefinitionArn, currTaskCount string
	if current != nil {
		currTaskDefinitionArn = *current.TaskDefinitionArn
		currTaskCount = strconv.FormatInt(*current.TaskCount, 10)
	}

	if !CompareString(&currTaskDefinitionArn, &expect.TaskDefinitionArn) {
		fmt.Printf("      TaskDefinitionArn: %s  ->  %s\n", currTaskDefinitionArn, expect.TaskDefinitionArn)
	}

	if current != nil && !CompareInt64(current.TaskCount, &expect.TaskCount) {
		fmt.Printf("      TaskCount: %s  ->  %d\n", currTaskCount, expect.TaskCount)
	}
}

// ShowDiffOfTheKinesisParameters print what will KinesisParameters in target changes to stdout
func ShowDiffOfTheKinesisParameters(current *cloudwatchevents.KinesisParameters, expect KinesisParameters) {
	var currPart string
	if current != nil {
		currPart = *current.PartitionKeyPath
	}
	fmt.Printf("      PartitionKeyPath: %s  ->  %s\n", currPart, expect.PartitionKeyPath)
}

// ShowWillDeleteRule print the rule will delete to stdout
func ShowWillDeleteRule(rule Rule) {
	if rule.NeedDelete {
		fmt.Printf("Rule: %s this will delete\n", *rule.ActualRule.Name)
	} else {
		fmt.Printf("Rule: %s\n", rule.Name)
	}
}

// ShowWillDeleteTarget print the target will delete to stdout
func ShowWillDeleteTarget(target Target) {
	fmt.Printf("  Target: %s this will delete\n", *target.ActualTarget.Id)
}
