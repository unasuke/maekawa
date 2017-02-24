package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// Associate ClowdWatchEvent Rule and descripbed Rule(name based)
func AssociateRules(cweRules []*cloudwatchevents.Rule, describedRules []Rule) []Rule {
	if l := len(cweRules) - len(describedRules); l > 0 {
		r := make([]Rule, l)
		describedRules = append(describedRules, r...)
	}
	dupCWERules := make([]*cloudwatchevents.Rule, len(cweRules))
	copy(dupCWERules, cweRules)

	for i, rule := range describedRules {
		for j, cweRule := range dupCWERules {
			if *cweRule.Name == rule.Name {
				describedRules[i].ActualRule = *cweRule
				dupCWERules = DeleteRuleFromSlice(dupCWERules, j)
				break
			}
		}
	}
	if len(dupCWERules) > 0 {
		for _, dupRule := range dupCWERules {
			for j, rule := range describedRules {
				if rule.ActualRule.Arn == nil {
					describedRules[j].ActualRule = *dupRule
				}
			}
		}
	}
	return describedRules
}

func AssociateTargets(cweTargets []*cloudwatchevents.Target, describedTargets []Target) []Target {
	// if ClowdWatchEvents Targets is more than declareted targets, append number of lack target{}
	if l := len(cweTargets) - len(describedTargets); l > 0 {
		t := make([]Target, l)
		describedTargets = append(describedTargets, t...)
	}
	dupCWETargets := make([]*cloudwatchevents.Target, len(cweTargets))
	copy(dupCWETargets, cweTargets)

	for i, target := range describedTargets {
		for j, cweTarget := range cweTargets {
			if CompareString(&target.Arn, cweTarget.Arn) && CompareString(&target.Id, cweTarget.Id) {
				describedTargets[i].ActualTarget = *cweTarget
				dupCWETargets = DeleteTargetFromSlice(dupCWETargets, j)
				break
			}
		}
	}
	if len(dupCWETargets) > 0 {
		for _, dupTarget := range dupCWETargets {
			for j, target := range describedTargets {
				if target.ActualTarget.Arn == nil {
					describedTargets[j].ActualTarget = *dupTarget
				}
			}
		}
	}

	return describedTargets
}

// judge is rule need update
// compare rule and ActualRule
func JudgeRuleNeedUpdate(r *Rule) {
	if !CompareString(&r.Name, r.ActualRule.Name) ||
		!CompareString(&r.Description, r.ActualRule.Description) ||
		!CompareString(&r.EventPattern, r.ActualRule.EventPattern) ||
		!CompareString(&r.ScheduleExpression, r.ActualRule.ScheduleExpression) ||
		!CompareString(&r.State, r.ActualRule.State) {
		r.NeedUpdate = true
	}
}

func JudgeRuleNeedDelete(r *Rule) {
	if r.Name == "" &&
		r.Description == "" &&
		r.EventPattern == "" &&
		r.ScheduleExpression == "" &&
		r.RoleArn == "" &&
		r.State == "" &&
		r.ActualRule.Name != nil {
		r.NeedDelete = true
	}
}

// judge is target need update
// compare target and ActualTarget
func JudgeTargetNeedUpdate(t *Target) {
	if !CompareString(&t.Input, t.ActualTarget.Input) ||
		!CompareString(&t.InputPath, t.ActualTarget.InputPath) {
		t.NeedUpdate = true
	}
}

func JudgeTargetNeedDelete(t *Target) {
	if t.Arn == "" &&
		t.Id == "" &&
		t.ActualTarget.Arn != nil {
		t.NeedDelete = true
	}
}

// check all rules and targets, update "NeedUpdate" and "NeedDelete" field
func CheckIsNeedUpdateOrDelete(rules []Rule) {
	for i, _ := range rules {
		JudgeRuleNeedUpdate(&rules[i])
		JudgeRuleNeedDelete(&rules[i])
		for j, _ := range rules[i].Targets {
			JudgeTargetNeedUpdate(&rules[i].Targets[j])
			JudgeTargetNeedDelete(&rules[i].Targets[j])
		}
	}
}
