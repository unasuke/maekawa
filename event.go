package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// return ClowdWatchEvent rule best matched one by yaml descripbed rule.
func fetchCWEventRuleFromDescribedRule(client *cloudwatchevents.CloudWatchEvents, describedRule Rule) (cloudwatchevents.Rule, error) {
	var bestMatchedRule cloudwatchevents.Rule
	var score float64

	resp, err := client.ListRules(nil)
	if err != nil {
		return bestMatchedRule, err
	}

	for _, rule := range resp.Rules {
		var s = MatchScoreForCWEventRuleAndDescribedRule(*rule, describedRule)

		if score < s {
			bestMatchedRule = *rule
			score = s
		}
	}
	return bestMatchedRule, nil
}

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

// return true when rule is new defined in yaml configration file
// judgemant by name(arn)
func IsNewDefinedRule(cweRule cloudwatchevents.Rule, describedRule Rule) bool {
	return *cweRule.Name != describedRule.Name
}
