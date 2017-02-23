package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

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
