package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func main() {
	var apply bool
	var dryrun bool
	var file string

	flag.BoolVar(&apply, "apply", false, "apply to CloudWatch Events")
	flag.BoolVar(&dryrun, "dry-run", false, "dry-run")
	flag.StringVar(&file, "file", "config.yml", "file path to setting yaml")
	flag.StringVar(&file, "f", "config.yml", "file path to setting yaml (shorthand)")
	flag.Parse()

	sess, err := session.NewSession(nil)
	if err != nil {
		fmt.Errorf("Error %v", err)
	}

	rules := Rules{}
	err := loadYaml(file, &rules)
	if err != nil {
		return err
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
