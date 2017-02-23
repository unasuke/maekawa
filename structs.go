package main

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

type Rules struct {
	Rules []Rule
}

type Rule struct {
	Description        string   `yaml:"description"`
	EventPattern       string   `yaml:"event_pattern"`
	Name               string   `yaml:"name"`
	ScheduleExpression string   `yaml:"schedule_expression"`
	State              string   `yaml:"state"`
	Targets            []Target `yaml:"targets"`
	ActualRule         cloudwatchevents.Rule
	NeedUpdate         bool
}

type Target struct {
	Arn          string `yaml:"arn"`
	Id           string `yaml:"id"`
	Input        string `yaml:"input"`
	InputPath    string `yaml:"input_path"`
	ActualTarget *cloudwatchevents.Target
	NeedUpdate   bool
}
