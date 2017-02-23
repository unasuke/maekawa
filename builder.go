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

// return ClowdWatchEvent Targets that deleted specified index target.
func DeleteTargetFromSlice(src []*cloudwatchevents.Target, deleteIndex int) []*cloudwatchevents.Target {
	dest := []*cloudwatchevents.Target{}
	for i, target := range src {
		if i != deleteIndex {
			dest = append(dest, target)
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
