package main

import (
	"github.com/aws/aws-sdk-go/aws"
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
				buildTarget(target),
			},
		})
		return err
	}

	return nil
}

func buildTarget(t Target) *cwe.Target {
	var target cwe.Target
	target.Arn = NilOrStringPtr(t.Arn)
	target.Id = NilOrStringPtr(t.Id)
	target.Input = NilOrStringPtr(t.Input)
	target.InputPath = NilOrStringPtr(t.InputPath)
	target.RoleArn = NilOrStringPtr(t.RoleArn)

	if t.EcsParameters.TaskDefinitionArn != "" && t.EcsParameters.TaskCount != 0 {
		target.EcsParameters = &cwe.EcsParameters{
			TaskDefinitionArn: NilOrStringPtr(t.EcsParameters.TaskDefinitionArn),
			TaskCount:         aws.Int64(t.EcsParameters.TaskCount),
		}
	}

	if t.KinesisParameters.PartitionKeyPath != "" {
		target.KinesisParameters = &cwe.KinesisParameters{
			PartitionKeyPath: NilOrStringPtr(t.KinesisParameters.PartitionKeyPath),
		}
	}

	return &target
}
