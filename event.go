package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// return match score ClowdWatchEvent rule and descripbed rule.
func MatchScoreForCWEventRuleAndDescribedRule(cweRule cloudwatchevents.Rule, describedRule Rule) float64 {
	const Elements = 5.0
	matchCount := 0.0

	if cweRule.Name != nil && *cweRule.Name == describedRule.Name {
		matchCount++
	}
	if cweRule.EventPattern != nil && *cweRule.EventPattern == describedRule.EventPattern {
		matchCount++
	}
	if cweRule.Description != nil && *cweRule.Description == describedRule.Description {
		matchCount++
	}
	if cweRule.ScheduleExpression != nil && *cweRule.ScheduleExpression == describedRule.ScheduleExpression {
		matchCount++
	}
	if cweRule.State != nil && *cweRule.State == describedRule.State {
		matchCount++
	}
	return (matchCount / Elements)
}
