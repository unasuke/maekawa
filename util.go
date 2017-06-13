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
		}
		return false
	} else if b == nil || *b == "" {
		if a == nil || *a == "" {
			return true
		}
		return false
	}
	return *a == *b
}

// CompareInt64 do compare Int64
// 1 == 1 => true
// 1 == 2 => false
// 1 == nil => false
// nil == nil => true
func CompareInt64(a, b *int64) bool {
	if a == nil {
		if b == nil {
			return true
		}
		return false
	} else if b == nil {
		return false
	}
	return *a == *b
}

// CompareKinesisParameters do compare KinesisParameters and cloudwatchevents.KinesisParameters
func CompareKinesisParameters(own *KinesisParameters, theirs *cwe.KinesisParameters) bool {
	if theirs == nil {
		if own.PartitionKeyPath == "" {
			return true
		}
	} else {
		if CompareString(&own.PartitionKeyPath, theirs.PartitionKeyPath) {
			return true
		}
	}
	return false
}

// CompareEcsParameters do compare EcsParameters and cloudwatchevents.EcsParameters
func CompareEcsParameters(own *EcsParameters, theirs *cwe.EcsParameters) bool {
	if theirs == nil {
		if own.TaskDefinitionArn == "" && own.TaskCount == 0 {
			return true
		}
	} else {
		if CompareString(&own.TaskDefinitionArn, theirs.TaskDefinitionArn) &&
			CompareInt64(&own.TaskCount, theirs.TaskCount) {
			return true
		}
	}
	return false
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
	}
	return *str
}
