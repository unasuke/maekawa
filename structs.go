package main

import (
	cwe "github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// Rules is for store unmarshalized configuration yaml
type Rules struct {
	Rules []Rule
}

// Rule is for expression CloudWatch Events Rule
type Rule struct {
	Description        string   `yaml:"description"`
	EventPattern       string   `yaml:"event_pattern"`
	Name               string   `yaml:"name"`
	RoleArn            string   `yaml:"role_arn"`
	ScheduleExpression string   `yaml:"schedule_expression"`
	State              string   `yaml:"state"`
	Targets            []Target `yaml:"targets"`
	ActualRule         cwe.Rule
	NeedUpdate         bool
	NeedDelete         bool
}

// Target is for expression CloudWatch Events Target
type Target struct {
	Arn               string            `yaml:"arn"`
	ID                string            `yaml:"id"`
	Input             string            `yaml:"input"`
	InputPath         string            `yaml:"input_path"`
	EcsParameters     EcsParameters     `yaml:"ecs_parameters"`
	KinesisParameters KinesisParameters `yaml:"kinesis_parameters"`
	RoleArn           string            `yaml:"role_arn"`
	ActualTarget      cwe.Target
	NeedUpdate        bool
	NeedDelete        bool
}

// EcsParameters is a part of the Target
type EcsParameters struct {
	TaskCount         int64  `yaml:"task_count"`
	TaskDefinitionArn string `yaml:"task_definition_arn"`
}

// KinesisParameters is a part of the Target
type KinesisParameters struct {
	PartitionKeyPath string `yaml:"partition_key_path"`
}

// LambdaPolicy is for JSON that return from Lambda.GetPolicy
type LambdaPolicy struct {
	Version   string             `json:"Version"`
	ID        string             `json:"Id"`
	Statement *[]PolicyStatement `json:"Statement"`
}

// PolicyStatement is a part of the LambdaPolicy
type PolicyStatement struct {
	Resource    string           `json:"Resource"`
	Condition   *PolicyCondition `json:"Condition"`
	StatementID string           `json:"Sid"`
	Effect      string           `json:"Effect"`
	Principal   *PolicyPrincipal `json:"Principal"`
	Action      string           `json:"Action"`
}

// PolicyCondition is a part of the LambdaPolicy
type PolicyCondition struct {
	ArnLike *PolicyArnLike `json:"ArnLike"`
}

// PolicyArnLike is a part of the LambdaPolicy
type PolicyArnLike struct {
	AwsSourceArn string `json:"AWS:SourceArn"`
}

// PolicyPrincipal is a part of the LambdaPolicy
type PolicyPrincipal struct {
	Service string `json:"Service"`
}
