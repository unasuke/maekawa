package main

import "testing"

func TestWillUpdateRulesAndTargets(t *testing.T) {
	rules := []Rule{
		{
			Name:       "Rule1",
			NeedUpdate: true,
		},
		{
			Name:       "Rule2",
			NeedUpdate: false,
			Targets: []Target{
				{
					Id:         "Target1",
					NeedUpdate: true,
				},
			},
		},
		{
			Name:       "Rule3",
			NeedUpdate: false,
		},
	}

	result := WillUpdateRulesAndTargets(rules)
	if len(result) != 2 {
		t.Errorf("result should be including 2 rules")
	}
	for _, rule := range result {
		if rule.Name == "Rule3" {
			t.Errorf("result shouldn't including Rule3")
		}
	}
}

func TestWillDeleteRulesAndTargets(t *testing.T) {
	rules := []Rule{
		{
			Name:       "Rule1",
			NeedDelete: true,
		},
		{
			Name:       "Rule2",
			NeedDelete: false,
			Targets: []Target{
				{
					Id:         "Target1",
					NeedDelete: true,
				},
			},
		},
		{
			Name:       "Rule3",
			NeedUpdate: false,
		},
	}

	result := WillDeleteRulesAndTargets(rules)
	if len(result) != 2 {
		t.Errorf("result should be including 2 rules")
	}
	for _, rule := range result {
		if rule.Name == "Rule3" {
			t.Errorf("result shouldn't including Rule3")
		}
	}
}
