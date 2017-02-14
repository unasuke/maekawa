package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	cwe := cloudwatchevents.New(sess)
	result, err := cwe.ListRules(nil)

	if err != nil {
		fmt.Println("Error", err)

	} else {
		fmt.Println("Success")
		fmt.Println(result)
	}
}
