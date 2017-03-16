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

// Version of the maekawa
var Version = "0.3.0"

func main() {
	var (
		apply, dryrun, version                  bool
		file, awsRegion                         string
		err                                     error
		sess                                    *session.Session
		cweRulesBeforeApply, cweRulesAfterApply *cloudwatchevents.ListRulesOutput
	)

	flag.BoolVar(&apply, "apply", false, "apply to CloudWatch Events")
	flag.BoolVar(&dryrun, "dry-run", false, "dry-run")
	flag.BoolVar(&version, "version", false, "show maekawa version")
	flag.BoolVar(&version, "v", false, "show maekawa version (shorthand)")
	flag.StringVar(&file, "file", "config.yml", "file path to setting yaml")
	flag.StringVar(&file, "f", "config.yml", "file path to setting yaml (shorthand)")
	flag.StringVar(&awsRegion, "region", os.Getenv("AWS_REGION"), "aws region")
	flag.Parse()

	if version {
		fmt.Printf("maekawa version %s\n", Version)
		os.Exit(0)
	}

	sess, err = session.NewSession(
		&aws.Config{
			Region: aws.String(awsRegion),
		},
	)
	if err != nil {
		fmt.Println("Session error\n", err)
		os.Exit(1)
	}

	cweClient := cloudwatchevents.New(sess)
	lambdaClient := lambda.New(sess)

	cweRulesBeforeApply, err = cweClient.ListRules(nil)
	if err != nil {
		fmt.Println("API error\n", err)
		os.Exit(1)
	}

	describedRules := Rules{}
	err = loadYaml(file, &describedRules)
	if err != nil {
		fmt.Println("File error\n", err)
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
			fmt.Println("API error\n", err)
			os.Exit(1)
		}

		err = removePermissonFromLambda(lambdaClient, describedRules.Rules)
		if err != nil {
			fmt.Println("API error\n", err)
			os.Exit(1)
		}

		// Grant permission to invoke lambda function from CloudWatch Events
		cweRulesAfterApply, err = cweClient.ListRules(nil)
		describedRules.Rules = AssociateRules(cweRulesAfterApply.Rules, describedRules.Rules)
		for i, rule := range describedRules.Rules {
			t, _ := fetchActualTargetsByRule(cweClient, rule)
			describedRules.Rules[i].Targets = AssociateTargets(t, describedRules.Rules[i].Targets)
		}

		err = addPermissionToLambdaFromCloudWatchEvents(lambdaClient, describedRules.Rules)
		if err != nil {
			fmt.Println("Grant permission error\n", err)
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
