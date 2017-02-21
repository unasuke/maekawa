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
