package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func main() {
	var apply bool
	var dryrun bool

	flag.BoolVar(&apply, "apply", false, "apply to CloudWatch Events")
	flag.BoolVar(&dryrun, "dry-run", false, "dry-run")
	flag.Parse()

	fmt.Println(apply)
	fmt.Println(dryrun)

	sess, err := session.NewSession(nil)
	if err != nil {
		fmt.Errorf("Error %v", err)
	}

	cwe := cloudwatchevents.New(sess)
	result, err := cwe.ListRules(nil)

	if err != nil {
		fmt.Println("Error", err)

	} else {
		fmt.Println("Success")
		fmt.Println(result)
	}
}
