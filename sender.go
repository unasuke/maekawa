package main

import (
	cwe "github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func updateCloudWatchEvents(client *cwe.CloudWatchEvents, rules []Rule) error {
	for _, rule := range rules {
		err := updateCloudWatchEventRule(client, rule)

		if err != nil {
			return err
		}

		for _, target := range rule.Targets {
			erro := updateCloudWatchEventTarget(client, rule, target)

			if erro != nil {
				return erro
			}
		}
	}
	return nil
}

func updateCloudWatchEventRule(client *cwe.CloudWatchEvents, rule Rule) error {
	if rule.NeedDelete {
		for _, target := range rule.Targets {
			_, err2 := client.RemoveTargets(&cwe.RemoveTargetsInput{
				Ids:  []*string{target.ActualTarget.Id},
				Rule: rule.ActualRule.Name,
			})
			if err2 != nil {
				return err2
			}
		}
		_, err3 := client.DeleteRule(&cwe.DeleteRuleInput{
			Name: rule.ActualRule.Name,
		})
		if err3 != nil {
			return err3
		}
	} else if rule.NeedUpdate {
		_, err := client.PutRule(&cwe.PutRuleInput{
			Name:               NilOrStringPtr(rule.Name),
			Description:        NilOrStringPtr(rule.Description),
			EventPattern:       NilOrStringPtr(rule.EventPattern),
			RoleArn:            NilOrStringPtr(rule.RoleArn),
			ScheduleExpression: NilOrStringPtr(rule.ScheduleExpression),
			State:              NilOrStringPtr(rule.State),
		})
		return err
	}
	return nil
}

func updateCloudWatchEventTarget(client *cwe.CloudWatchEvents, rule Rule, target Target) error {
	if target.NeedDelete {
		// if rule.NeedDelete is true, this target was removed in updateCloudWatchEventRule
		if !rule.NeedDelete {
			_, err2 := client.RemoveTargets(&cwe.RemoveTargetsInput{
				Ids:  []*string{target.ActualTarget.Id},
				Rule: rule.ActualRule.Name,
			})
			if err2 != nil {
				return err2
			}
		}
	} else if target.NeedUpdate {
		_, err := client.PutTargets(&cwe.PutTargetsInput{
			Rule: NilOrStringPtr(rule.Name),
			Targets: []*cwe.Target{
				{
					Arn:       NilOrStringPtr(target.Arn),
					Id:        NilOrStringPtr(target.Id),
					Input:     NilOrStringPtr(target.Input),
					InputPath: NilOrStringPtr(target.InputPath),
				},
			},
		})
		return err
	}

	return nil
}
