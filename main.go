package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

var Version = "0.2.0"

func main() {
	var (
		apply, dryrun   bool
		file, awsRegion string
	)

	flag.BoolVar(&apply, "apply", false, "apply to CloudWatch Events")
	flag.BoolVar(&dryrun, "dry-run", false, "dry-run")
	flag.StringVar(&file, "file", "config.yml", "file path to setting yaml")
	flag.StringVar(&file, "f", "config.yml", "file path to setting yaml (shorthand)")
	flag.StringVar(&awsRegion, "region", os.Getenv("AWS_REGION"), "aws region")
	flag.Parse()

	sess, errS := session.NewSession(
		&aws.Config{
			Region: aws.String(awsRegion),
		},
	)
	if errS != nil {
		fmt.Printf("Session error %v\n", errS)
		os.Exit(1)
	}

	cweRulesOutput, errR := cloudwatchevents.New(sess).ListRules(nil)
	if errR != nil {
		fmt.Printf("API error %v\n", errR)
		os.Exit(1)
	}

	describedRules := Rules{}
	errY := loadYaml(file, &describedRules)
	if errY != nil {
		fmt.Printf("File error %v\n", errY)
		os.Exit(1)
	}

	describedRules.Rules = AssociateRules(cweRulesOutput.Rules, describedRules.Rules)
	for i, rule := range describedRules.Rules {
		t, _ := fetchActualTargetsByRule(cloudwatchevents.New(sess), rule)
		describedRules.Rules[i].Targets = AssociateTargets(t, describedRules.Rules[i].Targets)
	}
	CheckIsNeedUpdateOrDelete(describedRules.Rules)
	displayWhatWillChange(describedRules.Rules)

	if apply && !dryrun {
		errU := updateCloudWatchEvents(cloudwatchevents.New(sess), describedRules.Rules)
		if errU != nil {
			fmt.Printf("API error %v\n", errU)
			os.Exit(1)
		}
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
