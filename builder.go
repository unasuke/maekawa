package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// Associate ClowdWatchEvent Rule and descripbed Rule(name based)
func AssociateRules(cweRules []*cloudwatchevents.Rule, describedRules []Rule) {
	dupCWERules := make([]*cloudwatchevents.Rule, cap(cweRules))
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
}

// return ClowdWatchEvent Rules that deleted specified index rule.
func DeleteRuleFromSlice(src []*cloudwatchevents.Rule, deleteIndex int) []*cloudwatchevents.Rule {
	dest := []*cloudwatchevents.Rule{}
	for i, rule := range src {
		if i != deleteIndex {
			dest = append(dest, rule)
		}
	}
	return dest
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
