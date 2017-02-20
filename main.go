package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func main() {
	var (
		apply, dryrun   bool
		file, awsRegion string
	)

	flag.BoolVar(&apply, "apply", false, "apply to CloudWatch Events")
	flag.BoolVar(&dryrun, "dry-run", false, "dry-run")
	flag.StringVar(&file, "file", "config.yml", "file path to setting yaml")
	flag.StringVar(&file, "f", "config.yml", "file path to setting yaml (shorthand)")
	flag.StringVar(&awsRegion, "region", "", "aws region")
	flag.Parse()

	sess, errS := session.NewSession(
		&aws.Config{
			Region: aws.String(awsRegion),
		},
	)
	if errS != nil {
		fmt.Errorf("Session error %v", errS)
	}

	cweRulesOutput, errR := cloudwatchevents.New(sess).ListRules(nil)
	if errR != nil {
		fmt.Errorf("API error %v", errR)
	}

	describedRules := Rules{}
	errY := loadYaml(file, &describedRules)
	if errY != nil {
		fmt.Errorf("File error %v", errY)
	}
}

func loadYaml(file string, r *Rules) error {

	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &r)
	if err != nil {
		return err
	}

	return nil
}
