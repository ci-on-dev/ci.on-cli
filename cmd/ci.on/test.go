package main

import (
	"fmt"
	"os"

	"github.com/ci-on-dev/ci.on-cli/internal/assert"
	"github.com/ci-on-dev/ci.on-cli/internal/report"
	"github.com/ci-on-dev/ci.on-cli/internal/runner"
)

func (service mainService) RunTest(testPath string) {

	suite, err := assert.LoadSuite(testPath)
	if err != nil {
		service.logService.Fatal(500, "assert_load_suite", err.Error())
	}

	var r runner.Runner
	if suite.Pipeline.Provider == "gitlab" {
		r = runner.NewLocalRunner(service.logService)
	} else {
		r = runner.NewDocker(runner.Config{}, service.logService)
	}

	// valida asserts
	res := service.assertService.ExecuteSuite(suite, r)
	exit := 0
	if !res.Passed {
		exit = 1
	}
	if err := report.WriteJUnitFile(res, "reports/junit.xml"); err != nil {
		fmt.Println("Erro ao gerar junit.xml:", err)
	}
	os.Exit(exit)
}
