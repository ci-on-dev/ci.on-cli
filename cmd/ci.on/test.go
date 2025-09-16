package main

import (
	"log"
	"os"

	"github.com/ci-on-dev/ci.on-cli/internal/assert"
	"github.com/ci-on-dev/ci.on-cli/internal/runner"
)

func RunTest(args []string) {
	if len(args) < 2 || args[0] != "-t" {
		log.Fatalf("usage: ci.on test -t <testfile>")
	}
	testPath := args[1]

	suite, err := assert.LoadSuite(testPath)
	if err != nil {
		log.Fatal(err)
	}

	var r runner.Runner
	if suite.Pipeline.Provider == "gitlab" {
		r = runner.NewLocalRunner()
	} else {
		r = runner.NewDocker(runner.Config{})
	}

	// valida asserts
	res := assert.ExecuteSuite(suite, r)
	exit := 0
	if !res.Passed {
		exit = 1
	}
	if err := res.WriteReports(suite.Reports); err != nil {
		log.Println("report:", err)
	}
	os.Exit(exit)
}
