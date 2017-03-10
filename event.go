package main

import (
	"fmt"

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

// fetch ClowdWatchEvent target by Rule.ActualRule
func fetchActualTargetsByRule(client *cloudwatchevents.CloudWatchEvents, r Rule) ([]*cloudwatchevents.Target, error) {
	if r.ActualRule.Name == nil {
		return nil, fmt.Errorf("Rule.ActualRule.Name is must be present")
	}

	targets, err := client.ListTargetsByRule(&cloudwatchevents.ListTargetsByRuleInput{
		Rule: r.ActualRule.Name,
	})
	if err != nil {
		return nil, err
	}
	return targets.Targets, nil
}

// MatchScoreForCWEventRuleAndDescribedRule return match score ClowdWatchEvent rule and descripbed rule.
func MatchScoreForCWEventRuleAndDescribedRule(cweRule cloudwatchevents.Rule, describedRule Rule) float64 {
	const Elements = 6.0
	matchCount := 0.0

	if CompareString(cweRule.Name, &describedRule.Name) {
		matchCount++
	}
	if CompareString(cweRule.EventPattern, &describedRule.EventPattern) {
		matchCount++
	}
	if CompareString(cweRule.Description, &describedRule.Description) {
		matchCount++
	}
	if CompareString(cweRule.ScheduleExpression, &describedRule.ScheduleExpression) {
		matchCount++
	}
	if CompareString(cweRule.RoleArn, &describedRule.RoleArn) {
		matchCount++
	}
	if CompareString(cweRule.State, &describedRule.State) {
		matchCount++
	}
	return (matchCount / Elements)
}

// IsNewDefinedRule return true when rule is new defined in yaml configration file
// judgemant by name(arn)
func IsNewDefinedRule(cweRule cloudwatchevents.Rule, describedRule Rule) bool {
	return *cweRule.Name != describedRule.Name
}
