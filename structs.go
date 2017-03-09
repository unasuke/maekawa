package main

import (
	cwe "github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// struct for store unmarshalized configuration yaml
type Rules struct {
	Rules []Rule
}

// struct for expression CloudWatch Events Rule
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

// struct for expression CloudWatch Events Target
type Target struct {
	Arn          string `yaml:"arn"`
	Id           string `yaml:"id"`
	Input        string `yaml:"input"`
	InputPath    string `yaml:"input_path"`
	ActualTarget cwe.Target
	NeedUpdate   bool
	NeedDelete   bool
}

// struct for JSON that return from Lambda.GetPolicy
type LambdaPolicy struct {
	Version   string             `json:"Version"`
	Id        string             `json:"Id"`
	Statement *[]PolicyStatement `json:"Statement"`
}

// part of the LambdaPolicy
type PolicyStatement struct {
	Resource    string           `json:"Resource"`
	Condition   *PolicyCondition `json:"Condition"`
	StatementId string           `json:"Sid"`
	Effect      string           `json:"Effect"`
	Principal   *PolicyPrincipal `json:"Principal"`
	Action      string           `json:"Action"`
}

// part of the LambdaPolicy
type PolicyCondition struct {
	ArnLike *PolicyArnLike `json:"ArnLike"`
}

// part of the LambdaPolicy
type PolicyArnLike struct {
	AwsSourceArn string `json:"AWS:SourceArn"`
}

// part of the LambdaPolicy
type PolicyPrincipal struct {
	Service string `json:"Service"`
}
