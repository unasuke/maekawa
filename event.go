package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// return match score ClowdWatchEvent rule and descripbed rule.
func MatchScoreForCWEventRuleAndDescribedRule(cweRule cloudwatchevents.Rule, describedRule Rule) float64 {
	const Elements = 5.0
	matchCount := 0.0

	if *cweRule.Name == describedRule.Name {
		matchCount++
	}
	if *cweRule.EventPattern == describedRule.EventPattern {
		matchCount++
	}
	if *cweRule.Description == describedRule.Description {
		matchCount++
	}
	if *cweRule.ScheduleExpression == describedRule.ScheduleExpression {
		matchCount++
	}
	if *cweRule.State == describedRule.State {
		matchCount++
	}
	return (matchCount / Elements)
}
