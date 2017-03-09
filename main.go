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
	"github.com/aws/aws-sdk-go/service/lambda"
)

var Version = "0.2.0"

func main() {
	var (
		apply, dryrun                           bool
		file, awsRegion                         string
		err                                     error
		sess                                    *session.Session
		cweRulesBeforeApply, cweRulesAfterApply *cloudwatchevents.ListRulesOutput
	)

	flag.BoolVar(&apply, "apply", false, "apply to CloudWatch Events")
	flag.BoolVar(&dryrun, "dry-run", false, "dry-run")
	flag.StringVar(&file, "file", "config.yml", "file path to setting yaml")
	flag.StringVar(&file, "f", "config.yml", "file path to setting yaml (shorthand)")
	flag.StringVar(&awsRegion, "region", os.Getenv("AWS_REGION"), "aws region")
	flag.Parse()

	sess, err = session.NewSession(
		&aws.Config{
			Region: aws.String(awsRegion),
		},
	)
	if err != nil {
		fmt.Printf("Session error %v\n", err)
		os.Exit(1)
	}

	cweClient := cloudwatchevents.New(sess)
	lambdaClient := lambda.New(sess)

	cweRulesBeforeApply, err = cweClient.ListRules(nil)
	if err != nil {
		fmt.Printf("API error %v\n", err)
		os.Exit(1)
	}

	describedRules := Rules{}
	err = loadYaml(file, &describedRules)
	if err != nil {
		fmt.Printf("File error %v\n", err)
		os.Exit(1)
	}

	describedRules.Rules = AssociateRules(cweRulesBeforeApply.Rules, describedRules.Rules)
	for i, rule := range describedRules.Rules {
		t, _ := fetchActualTargetsByRule(cweClient, rule)
		describedRules.Rules[i].Targets = AssociateTargets(t, describedRules.Rules[i].Targets)
	}
	CheckIsNeedUpdateOrDelete(describedRules.Rules)
	displayWhatWillChange(describedRules.Rules)

	if apply && !dryrun {
		err = updateCloudWatchEvents(cweClient, describedRules.Rules)
		if err != nil {
			fmt.Printf("API error %v\n", err)
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
