package main

import "fmt"


// return will be updated rules and targets
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

// return will be deleted rules and targets
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
