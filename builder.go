package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// AssociateRules is asociate ClowdWatchEvent Rule and descripbed Rule(name based)
func AssociateRules(cweRules []*cloudwatchevents.Rule, describedRules []Rule) []Rule {
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
			describedRules = append(describedRules, Rule{
				ActualRule: *dupRule,
			})
		}
	}
	return describedRules
}

// AssociateTargets is asociate ClowdWatchEvent targets and descripbed targets(Id based)
func AssociateTargets(cweTargets []*cloudwatchevents.Target, describedTargets []Target) []Target {
	// if ClowdWatchEvents Targets is more than declareted targets, append number of lack target{}
	dupCWETargets := make([]*cloudwatchevents.Target, len(cweTargets))
	copy(dupCWETargets, cweTargets)

	for i, target := range describedTargets {
		for j, cweTarget := range dupCWETargets {
			if CompareString(&target.Arn, cweTarget.Arn) && CompareString(&target.Id, cweTarget.Id) {
				describedTargets[i].ActualTarget = *cweTarget
				dupCWETargets = DeleteTargetFromSlice(dupCWETargets, j)
				break
			}
		}
	}
	if len(dupCWETargets) > 0 {
		for _, dupTarget := range dupCWETargets {
			describedTargets = append(describedTargets, Target{
				ActualTarget: *dupTarget,
			})
		}
	}
	return describedTargets
}

// JudgeRuleNeedUpdate is judge the rule need update
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

// JudgeRuleNeedDelete is judge the rule need delete
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

// JudgeTargetNeedUpdate is judge the target need update
// compare target and ActualTarget
func JudgeTargetNeedUpdate(t *Target) {
	if !CompareString(&t.Arn, t.ActualTarget.Arn) ||
		!CompareString(&t.Id, t.ActualTarget.Id) ||
		!CompareString(&t.Input, t.ActualTarget.Input) ||
		!CompareString(&t.InputPath, t.ActualTarget.InputPath) {
		t.NeedUpdate = true
	}
}

// JudgeTargetNeedDelete is judge the target need delete
func JudgeTargetNeedDelete(t *Target) {
	if t.Arn == "" &&
		t.Id == "" &&
		t.ActualTarget.Arn != nil {
		t.NeedDelete = true
	}
}

// CheckIsNeedUpdateOrDelete is check all rules and targets,
// update "NeedUpdate" and "NeedDelete" field
func CheckIsNeedUpdateOrDelete(rules []Rule) {
	for i := range rules {
		JudgeRuleNeedUpdate(&rules[i])
		JudgeRuleNeedDelete(&rules[i])
		for j := range rules[i].Targets {
			JudgeTargetNeedUpdate(&rules[i].Targets[j])
			JudgeTargetNeedDelete(&rules[i].Targets[j])
		}
	}
}
