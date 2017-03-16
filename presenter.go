package main

import "fmt"

func displayWhatWillChange(rules []Rule) {
	updates := WillUpdateRulesAndTargets(rules)
	deletes := WillDeleteRulesAndTargets(rules)
	if len(updates) == 0 && len(deletes) == 0 {
		fmt.Println("No Changes")
	}
	if len(updates) > 0 {
		fmt.Println("Updates")
		for _, r := range updates {
			ShowWillUpdateFieldInRule(r)
			for _, t := range r.Targets {
				if t.NeedUpdate && !t.NeedDelete {
					ShowWillUpdateFieldInTarget(t)
				}
			}
		}
	}
	if len(deletes) > 0 {
		fmt.Println("Deletes")
		for _, r := range deletes {
			ShowWillDeleteRule(r)
			for _, t := range r.Targets {
				if t.NeedDelete {
					ShowWillDeleteTarget(t)
				}
			}
		}
	}
}

// WillUpdateRulesAndTargets return will be updated rules and targets
func WillUpdateRulesAndTargets(rules []Rule) []Rule {
	u := make([]Rule, 0)
	for _, rule := range rules {
		if rule.NeedUpdate && !rule.NeedDelete {
			u = append(u, rule)
		} else {
			for _, target := range rule.Targets {
				if target.NeedUpdate && !target.NeedDelete {
					u = append(u, rule)
					break
				}
			}
		}
	}
	return u
}

// WillDeleteRulesAndTargets return will be deleted rules and targets
func WillDeleteRulesAndTargets(rules []Rule) []Rule {
	d := make([]Rule, 0)
	for _, rule := range rules {
		if rule.NeedDelete {
			d = append(d, rule)
		} else {
			for _, target := range rule.Targets {
				if target.NeedDelete {
					d = append(d, rule)
					break
				}
			}
		}
	}
	return d
}

// ShowWillUpdateFieldInRule print what will rule changes to stdout
func ShowWillUpdateFieldInRule(rule Rule) {
	fmt.Printf("Rule: %s\n", rule.Name)
	if !CompareString(&rule.Name, rule.ActualRule.Name) {
		fmt.Printf("Name: %s \t->\t %s\n", NilSafeStr(rule.ActualRule.Name), rule.Name)
	}
	if !CompareString(&rule.Description, rule.ActualRule.Description) {
		fmt.Printf("Description: %s \t->\t %s\n", NilSafeStr(rule.ActualRule.Description), rule.Description)
	}
	if !CompareString(&rule.EventPattern, rule.ActualRule.EventPattern) {
		fmt.Printf("EventPattern: %s \t->\t %s\n", NilSafeStr(rule.ActualRule.EventPattern), rule.EventPattern)
	}
	if !CompareString(&rule.RoleArn, rule.ActualRule.RoleArn) {
		fmt.Printf("RoleArn: %s \t->\t %s\n", NilSafeStr(rule.ActualRule.RoleArn), rule.RoleArn)
	}
	if !CompareString(&rule.ScheduleExpression, rule.ActualRule.ScheduleExpression) {
		fmt.Printf("ScheduleExpression: %s \t->\t %s\n", NilSafeStr(rule.ActualRule.ScheduleExpression), rule.ScheduleExpression)
	}
	if !CompareString(&rule.State, rule.ActualRule.State) {
		fmt.Printf("State: %s \t->\t %s\n", NilSafeStr(rule.ActualRule.State), rule.State)
	}
}

// ShowWillUpdateFieldInTarget print what will target changes to stdout
func ShowWillUpdateFieldInTarget(target Target) {
	fmt.Printf("Target: %s\n", target.Arn)
	if !CompareString(&target.Arn, target.ActualTarget.Arn) {
		fmt.Printf("Arn: %s \t->\t %s\n", NilSafeStr(target.ActualTarget.Arn), target.Arn)
	}
	if !CompareString(&target.Id, target.ActualTarget.Id) {
		fmt.Printf("Id: %s \t->\t %s\n", NilSafeStr(target.ActualTarget.Id), target.Id)
	}
	if !CompareString(&target.Input, target.ActualTarget.Input) {
		fmt.Printf("Input: %s \t->\t %s\n", NilSafeStr(target.ActualTarget.Input), target.Input)
	}
	if !CompareString(&target.InputPath, target.ActualTarget.InputPath) {
		fmt.Printf("InputPath: %s \t->\t %s\n", NilSafeStr(target.ActualTarget.InputPath), target.InputPath)
	}
}

// ShowWillDeleteRule print the rule will delete to stdout
func ShowWillDeleteRule(rule Rule) {
	if rule.NeedDelete {
		fmt.Printf("Rule: %s this will delete\n", *rule.ActualRule.Name)
	} else {
		fmt.Printf("Rule: %s\n", rule.Name)
	}
}

// ShowWillDeleteTarget print the target will delete to stdout
func ShowWillDeleteTarget(target Target) {
	fmt.Printf("Target: %s this will delete\n", *target.ActualTarget.Id)
}
