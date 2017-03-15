package main

import (
	cwe "github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// CompareString do compare strings
// nil and empty string are same value
// "" == "" => true
// nil == "" => true
// nil == nil => true
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

// DeleteRuleFromSlice return ClowdWatchEvent Rules that deleted specified index rule.
func DeleteRuleFromSlice(src []*cwe.Rule, deleteIndex int) []*cwe.Rule {
	dest := []*cwe.Rule{}
	for i, rule := range src {
		if i != deleteIndex {
			dest = append(dest, rule)
		}
	}
	return dest
}

// DeleteTargetFromSlice return ClowdWatchEvent Targets that deleted specified index target.
func DeleteTargetFromSlice(src []*cwe.Target, deleteIndex int) []*cwe.Target {
	dest := []*cwe.Target{}
	for i, target := range src {
		if i != deleteIndex {
			dest = append(dest, target)
		}
	}
	return dest
}

// NilOrStringPtr return string pointer if str is not empty
// if str is empty ("") retuen nil
func NilOrStringPtr(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

// NilSafeStr retuen empty string or passed string
func NilSafeStr(str *string) string {
	if str == nil {
		return ""
	} else {
		return *str
	}
}
