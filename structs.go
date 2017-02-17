package main

type Rules struct {
	Rules []Rule
}

type Rule struct {
	Description        string           `yaml:"description"`
	EventPattern       string           `yaml:"event_pattern"`
	Name               string           `yaml:"name"`
	ScheduleExpression string           `yaml:"schedule_expression"`
	State              string           `yaml:"state"`
	LambdaFunctions    []LambdaFunction `yaml:"lambda_functions"`
}

type LambdaFunction struct {
	Name      string `yaml:"name"`
	Input     string `yaml:"input"`
	InputPath string `yaml:"input_path"`
}
