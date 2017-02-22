package main

import (
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
func fetchActualTargetsByRules(client *cloudwatchevents.CloudWatchEvents, rules *Rules) error {
	for i, rule := range rules.Rules {
		if rule.ActualRule.Name == nil {
			continue
		}
		targets, err := client.ListTargetsByRule(&cloudwatchevents.ListTargetsByRuleInput{
			Rule: rule.ActualRule.Name,
		})
		if err != nil {
			return err
		}
		rules.Rules[i].ActualTargets = targets.Targets
	}
	return nil
}

// return match score ClowdWatchEvent rule and descripbed rule.
func MatchScoreForCWEventRuleAndDescribedRule(cweRule cloudwatchevents.Rule, describedRule Rule) float64 {
	const Elements = 5.0
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
	if CompareString(cweRule.State, &describedRule.State) {
		matchCount++
	}
	return (matchCount / Elements)
}

// return true when rule is new defined in yaml configration file
// judgemant by name(arn)
func IsNewDefinedRule(cweRule cloudwatchevents.Rule, describedRule Rule) bool {
	return *cweRule.Name != describedRule.Name
}

// compare strings
// nil and empty string are same value
func CompareString(a, b *string) bool {
	if a == nil || *a == "" {
		if b == nil || *b == "" {
			return true
		} else {
			return false
		}
	} else if b == nil || *b == "" {
		if a == nil || *a == "" {
			return true
		} else {
			return false
		}
	}
	return *a == *b
}
